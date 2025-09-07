package conf

import (
	"encoding/json"

	model_config "xi/internal/app/model/config"
)

// Alias to *koanf.Koanf.Get()
func (c *ConfLib) RGet(path string) any {
	return c.Raw.Get(path)
}

// Return all config as a map
func (c *ConfLib) RAll() (map[string]any, error) {
	var cfgRaw map[string]any
	if err := c.Raw.Unmarshal("", &cfgRaw); err != nil {
		return nil, err
	}
	return cfgRaw, nil
}

// Return all config as JSON bytes
func (c *ConfLib) RAllJson() ([]byte, error) {
	cfg, err := c.RAll()
	if err != nil {
		return nil, err
	}
	return json.Marshal(cfg)
}

// Return all config as pretty-printed JSON bytes
func (c *ConfLib) RAllJsonPretty() ([]byte, error) {
	cfg, err := c.RAll()
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(cfg, "", "  ")
}

// Return all config as strongly-typed model
func (c *ConfLib) All() (model_config.Config, error) {
	cfgBytes, err := c.RAllJson()
	if err != nil {
		return model_config.Config{}, err
	}

	var cfg model_config.Config
	if err := json.Unmarshal(cfgBytes, &cfg); err != nil {
		return model_config.Config{}, err
	}

	return cfg, nil
}

// Return strongly-typed config as JSON bytes
func (c *ConfLib) AllJson() ([]byte, error) {
	cfg, err := c.All()
	if err != nil {
		return nil, err
	}
	return json.Marshal(cfg)
}

// Return strongly-typed config as pretty-printed JSON bytes
func (c *ConfLib) AllJsonPretty() ([]byte, error) {
	cfg, err := c.All()
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(cfg, "", "  ")
}
