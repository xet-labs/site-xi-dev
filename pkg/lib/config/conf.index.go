package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	model_config "xi/internal/app/model/config"
	"xi/pkg/lib/cfg"
	"xi/pkg/lib/env"
	"xi/pkg/lib/hook"
	"xi/pkg/lib/util"

	"github.com/fsnotify/fsnotify"
	koanfJson "github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

type ConfigLib struct {
	Dir          []string
	DirDefault   []string
	Files        []string
	FilesDefault []string
	FilesLoaded  []string

	LoadDefaults *bool
	Initialized  bool

	Hooks *hook.Hook
	Raw   *koanf.Koanf
	watch *fsnotify.Watcher
	mu    sync.RWMutex
	once  sync.Once
}

var (
	Config = &ConfigLib{
		DirDefault: []string{"pkg/config", "internal/app/config", "config"},
		Hooks:      &hook.Hook{},
	}

	reJsonEnv         = regexp.MustCompile(`\$\{([A-Z0-9_]+)(:-([^}]*))?\}`)
	reJsonEnvPost     = regexp.MustCompile(`(?m)(,\s*)?__REMOVE__(,\s*)?|^__REMOVE__(,\s*)?`)
	reJsonDoubleQuote = regexp.MustCompile(`""([^"\n\r]+?)""`)
	reJsonIntCast     = regexp.MustCompile(`:\s*"(-?\d+)\.int"`)
	reJsonBoolStr     = regexp.MustCompile(`:\s*"(true|false|1|0)"`)
	reJsonVar         = regexp.MustCompile(`\$\{([^}:]*)(:-([^}]*))?\}|\$\{\}`)
)

func (c *ConfigLib) Init(filePath ...string) {
	c.once.Do(func() {
		c.Hooks.AddPost(PostHooks...)
		c.LoadConfigs()
		c.InitCore(filePath...)

		if err := Config.Daemon(); err != nil {
			log.Warn().Caller().Err(err).Msgf("config Daemon: setup failed")
		}
	})
}

func (c *ConfigLib) InitCore(files ...string) error {
	env.Env.Init()
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Hooks.RunPre()

	if len(files) > 0 {
		c.Files = files
	} else if !c.Initialized {
		c.Files = c.FilesDefault
	}

	// log.Debug().
	// 	Str("dir", strings.Join(c.Dir, ", ")).
	// 	Str("dirDefault", strings.Join(c.DirDefault, ", ")).
	// 	Str("files", strings.Join(c.Files, ", ")).
	// 	Str("filesDefault", strings.Join(c.FilesDefault, ", ")).Msg("config params")

	if len(c.Files) == 0 {
		if !c.Initialized {
			err := errors.New("no configuration files to load")
			log.Fatal().Caller().Err(err).Msg("config")
			return err
		}
		log.Warn().Caller().Err(errors.New("no configuration files to reload")).Msg("config")
	}

	// Create a new Koanf instance for atomic reload
	newKoanf := koanf.New(".")
	var newFilesLoaded []string
	// Load all config files into newKoanf
	for _, path := range c.Files {

		raw, err := os.ReadFile(path)
		if err != nil {
			log.Warn().Caller().Err(err).Str("file", path).
				Msg("config couldnt get file")
			continue
		}

		// Preproces and load env from the config
		// the env needs to be loaded and reprocess the config as it might be using env
		if _, tmp, err := c.preProcess(raw); err == nil {
			for _, file := range tmp.App.EnvFiles {
				if err := env.Env.Load(file); err != nil {
					if !c.Initialized {
						log.Fatal().Caller().Err(err).Str("env", file).Str("file", path).
							Msg("config Preprocess failed to load Env")
					} else {
						log.Warn().Err(err).Str("env", file).Str("file", path).
							Msg("config Preprocess failed to load Env")
					}
					continue
				}
			}
		} else {
			if !c.Initialized {
				log.Error().Caller().Err(err).Str("file", path).Msg("config Preprocess failed for env")
			} else {
				log.Warn().Err(err).Str("file", path).Msg("config Preprocess failed for env")

			}
			continue
		}

		// Fully preprocess data
		var resolved []byte
		if resolved, _, err = c.preProcess(raw); err != nil {
			if !c.Initialized {
				log.Error().Caller().Err(err).Str("file", path).Msg("config Preprocess failed")
			} else {
				log.Warn().Err(err).Str("file", path).Msg("config Preprocess failed")
			}
		}
		// Sync/merge Json data
		if err := newKoanf.Load(rawbytes.Provider(resolved), koanfJson.Parser()); err != nil {
			if !c.Initialized {
				log.Error().Caller().Err(err).Str("file", path).
					Msg("config load failed: JSON is valid but merging into runtime configuration failed")
			} else {
				log.Warn().Err(err).Str("file", path).
					Msg("config load failed: JSON is valid but merging into runtime configuration failed")
			}
			continue
		}
		newFilesLoaded = append(newFilesLoaded, path)
	}

	// Fail/Warn if nothing loaded/reloaded
	if len(newFilesLoaded) == 0 {
		if !c.Initialized {
			err := errors.New("no configuration files could be loaded")
			log.Fatal().Caller().Err(err).
				Str("files", strings.Join(c.Files, ", ")).Msg("config")
			return err
		}

		err := errors.New("no configuration files could be reloaded")
		log.Warn().Err(err).
			Str("files", strings.Join(c.Files, ", ")).Msg("config")
		return err
	}

	// Globally store newKoanf
	c.Raw = newKoanf
	cfg.RUpdate(newKoanf)
	c.FilesLoaded = newFilesLoaded

	c.PostProcess()

	// Success msg on Config load/reload
	if !c.Initialized {
		log.Info().Str("files", strings.Join(newFilesLoaded, ", ")).Msg("config loaded")
	} else {
		log.Debug().Str("files", strings.Join(newFilesLoaded, ", ")).Msg("config reloaded")
	}
	c.Initialized = true
	return nil
}

// Load Config Files
func (c *ConfigLib) LoadConfigs() {
	if c.LoadDefaults != nil && !*c.LoadDefaults {
		return
	}
	// Generate config files list
	for _, dir := range c.DirDefault {
		files, err := util.File.GetWithExt(".json", dir)
		if err != nil {
			log.Warn().Err(err).Str("path", dir).Str("dir", dir).Msg("config error accessing")
		}
		c.FilesDefault = append(c.FilesDefault, util.Str.UniqueSort(files)...)
	}
}

// Process and store Config globally
func (c *ConfigLib) PostProcess() {
	// Process PostHooks and their data
	rawDat, errs := c.Hooks.RunPost()
	for _, e := range errs {
		c.Error(e)
	}
	for _, r := range rawDat {
		if rawConf, ok := r.(map[string]any); ok {
			c.MergeConf(&rawConf)
		}
	}

	// Generate Global Config
	rawCfg, err := c.All()
	if err != nil {
		if !c.Initialized {
			log.Fatal().Caller().Err(err).Msg("global config generation failed")
		}
		log.Warn().Caller().Err(err).Msg("global config generation failed")
		return
	}

	// Store Global Config to global 'cfg'
	cfg.Update(rawCfg)
}

func (c *ConfigLib) Error(err error) {
	panic("unimplemented")
}

func (c *ConfigLib) preProcess(rawJson []byte) ([]byte, model_config.Config, error) {
	// resolve Json varsand cleanups
	resolved, err := c.resolveJsonVars(c.cleanJson(c.resolveJsonEnv(string(rawJson))))
	if err != nil {
		log.Error().Err(fmt.Errorf("unable to resolve environment variables or sanitize JSON")).Msg("config preprocess failed")
		return nil, model_config.Config{}, err
	}

	// Validate against config model
	structured := model_config.Config{}
	if err := json.Unmarshal([]byte(resolved), &structured); err != nil {
		log.Error().Caller().Err(fmt.Errorf("json does not match the expected Config structure")).
			Msg("config preprocess failed, invalid format")
		return nil, model_config.Config{}, err
	}

	return []byte(resolved), structured, nil
}

// resolveJsonEnv replaces ${ENV} or ${ENV:-fallback} with actual values
func (c *ConfigLib) resolveJsonEnv(input string) string {
	out := reJsonEnv.ReplaceAllStringFunc(input, func(match string) string {
		sub := reJsonEnv.FindStringSubmatch(match)
		key, def := sub[1], sub[3] // ENV, fallback

		if val, ok := os.LookupEnv(key); ok {
			return val
		}

		// Fallback
		if def != "" {
			return def
		}

		// No value, no fallback
		return "__REMOVE__"
	})
	// fmt.Printf("\n%s\n", out)
	return out
}

func (c *ConfigLib) cleanJson(input string) string {
	// Remove "__REMOVE__"
	out := reJsonEnvPost.ReplaceAllString(input, "")

	// Optionally: fix trailing commas, multiple newlines, etc.
	out = strings.ReplaceAll(out, ",\n}", "\n}")
	out = strings.ReplaceAll(out, ",\n]", "\n]")

	// Fix ""value"" to "value", but skip empty ""
	out = reJsonDoubleQuote.ReplaceAllString(out, `"$1"`)

	out = reJsonIntCast.ReplaceAllString(out, ": $1")

	// Replaces string "true"/"false" -> true/false
	out = reJsonBoolStr.ReplaceAllStringFunc(out, func(match string) string {
		// Extract actual boolean value from the match using the submatch
		submatches := reJsonBoolStr.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}

		switch strings.ToLower(submatches[1]) {
		case "true":
			return ": true"
		case "false":
			return ": false"
		}
		return match // fallback
	})

	// fmt.Printf("\n%s\n", out)
	return out
}

// resolveJsonVars walks entire koanf data and resolves {{key.path}} expressions
func (c *ConfigLib) resolveJsonVars(input string) (string, error) {
	var data any
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return "", err
	}

	var resolveValue func(any) any

	resolveValue = func(val any) any {
		switch v := val.(type) {
		case string:
			return reJsonVar.ReplaceAllStringFunc(v, func(match string) string {
				sub := reJsonVar.FindStringSubmatch(match)
				key := sub[1]
				def := sub[3]
				if val := c.Raw.String(key); val != "" {
					return val
				}
				return def
			})
		case map[string]any:
			for k, vv := range v {
				if str, ok := vv.(string); ok && reJsonVar.MatchString(str) {
					sub := reJsonVar.FindStringSubmatch(str)
					key := sub[1]
					def := sub[3]

					val := c.Raw.Get(key)
					if val != nil {
						switch typed := val.(type) {
						case map[string]any, []any:
							// If exactly like "${key}", replace whole field with object/array
							if str == "${"+key+"}" {
								v[k] = typed
							} else {
								v[k] = str
							}
						case string:
							if typed != "" {
								v[k] = typed
							} else {
								v[k] = def
							}
						default:
							v[k] = typed
						}
					} else {
						v[k] = def
					}
				} else {
					v[k] = resolveValue(vv)
				}
			}
			return v

		case []any:
			for i, vv := range v {
				v[i] = resolveValue(vv)
			}
			return v

		default:
			return v
		}
	}

	resolved := resolveValue(data)

	outBytes, err := json.Marshal(resolved)
	if err != nil {
		return "", err
	}
	return string(outBytes), nil
}

// sync json connfig with existing config
func (c *ConfigLib) MergeConf(rawConf *map[string]any) error {
	// Convert map[string]any to proper []byte(json) for further processing
	jsonConf, err := json.Marshal(rawConf)
	if err != nil {
		log.Error().Caller().Err(err)
		return err
	}

	// Merge config
	if err := c.Raw.Load(rawbytes.Provider(jsonConf), koanfJson.Parser()); err != nil {
		log.Error().Caller().Err(err)
		return err
	}

	return nil
}

// Config Daemon to reload config file changes
func (c *ConfigLib) Daemon() error {
	if c.watch != nil {
		return nil // already watching
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Warn().Err(err).Msg("config Daemon failed to launch")
		return err
	}
	c.watch = watcher

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove|fsnotify.Rename) != 0 {
					log.Debug().Str("event", event.Op.String()).Str("file", event.Name).Msg("config changed")

					if err := c.InitCore(); err != nil {
						log.Warn().Caller().Err(err).Msg("config reload failed")
					}
					// Sleep briefly to avoid partial writes
					time.Sleep(100 * time.Millisecond)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Warn().Caller().Err(err).Msgf("config.daemon file watch err")
			}
		}
	}()

	for _, path := range c.Files {
		// Ensure file exists before watching (else no event will be triggered)
		if _, err := os.Stat(path); err == nil {
			if err := watcher.Add(path); err != nil {
				log.Warn().Msgf("config daemon failed to watch %s: %v", path, err)
			}
		} else {
			log.Warn().Msgf("config daemon missing file: %s", path)
		}
	}

	return nil
}
