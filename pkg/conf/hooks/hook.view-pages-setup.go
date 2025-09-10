package confHook

import (
	"errors"
	"maps"

	"xi/pkg/cfg"
	"github.com/rs/zerolog/log"
)

// In your ConfLib initialization:
func ViewPagesSetup(args ...any) (any, error){

	// Fetch defaults and pages
	pageDefault := cfg.RGet("web.default").(map[string]any)
	pages := cfg.RGet("web.pages").(map[string]any)

	rawJson, err := cfg.RAll()
	if err != nil {
		log.Error().Caller().Err(err).Msg("config post hook")
		return nil, err
	}
	viewData, ok := rawJson["web"].(map[string]any)
	if !ok {
		err := error(errors.New("'web' is missing or not a map"))
		log.Error().Caller().Err(err).Msg("config post hook")
		return nil, err
	}

	// Ensure "pages" exists inside viewData
	viewPages, ok := viewData["pages"].(map[string]any)
	if !ok {
		// Create it if missing
		viewPages = make(map[string]any)
		viewData["pages"] = viewPages
	}

	// Merge defaults into each page
	for page, val := range pages {
		pageConf, ok := val.(map[string]any)
		if !ok {
			log.Warn().Caller().Str("page", page).Msg("config Postview: Page data setup failed")
			continue
		}

		// Copy defaults first, then page-specific config
		rawConf := make(map[string]any)
		if pageDefault != nil {
			maps.Copy(rawConf, pageDefault)
		}
		maps.Copy(rawConf, pageConf)

		viewPages[page] = rawConf
	}

    return rawJson, nil
}
