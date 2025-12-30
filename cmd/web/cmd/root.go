package cmd

import (
	"fmt"
	"os"

	webApp "xi/internal/app/web"
	"xi/internal/config/cfg"
	"xi/internal/infra"
	"xi/internal/services"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xlvini [global options] [command] [command options]",
	Short: "xlvini",
	Long:  "xlvini — " + cfg.Org.Abbr + "'s command-line tool and agent",
	Run: func(cmd *cobra.Command, args []string) {
		cfg.App.Initialized = true

		infra.Logger.InitCore()
		services.InitCore() // Init services
		services.InitPost() // Init services

		webApp.App.Run(cfg.App.Port) // Init server
	},
}

// Execute executes the root command
func Init() {
	Execute()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	// Persistent flags
	a := rootCmd.PersistentFlags()

	a.StringArrayVarP(&cfg.App.ConfigFiles, "config", "", []string{}, "config files (default 'pkg/config/*', 'config/*')")
	a.StringArrayVarP(&cfg.App.ConfigDirs, "config-dir", "", []string{}, "config directories (default 'pkg/config/', 'config/')")
	a.BoolVarP(&cfg.App.ConfigHotReloadSig, "config-hot-reload", "", false, "enable config Hot-Reload")
	a.StringVarP(&cfg.App.Port, "port", "p", cfg.App.Port, "port to run server")
	a.BoolVarP(&cfg.App.Verbose, "verbose", "v", false, "enable verbose logging")
}
