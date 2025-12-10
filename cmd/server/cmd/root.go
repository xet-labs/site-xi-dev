package cmd

import (
	"fmt"
	"os"

	"xi/internal/service"
	appPkg "xi/pkg/app"
	"xi/pkg/lib"
	"xi/pkg/lib/cfg"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xi-server",
	Short: "XI HTTP server",
	Long:  "XI HTTP server — the main backend for the XI",
	Run: func(cmd *cobra.Command, args []string) {
	 	cfg.App.Initialized = true

		lib.Logger.InitCore()
		service.InitCore() // Init services
		service.InitPost() // Init services

		appPkg.Server.Init() // Init server
		appPkg.Server.Start(cfg.App.Port) // Init server
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
	a:=rootCmd.PersistentFlags()

	a.StringArrayVarP(&cfg.App.ConfigFiles, "config", "", []string{}, "config files (default 'pkg/config/*', 'config/*')")
	a.StringArrayVarP(&cfg.App.ConfigDirs, "config-dir", "", []string{}, "config directories (default 'pkg/config/', 'config/')")
	a.BoolVarP(&cfg.App.ConfigHotReloadSig, "config-hot-reload", "", false, "enable config Hot-Reload")
	a.StringVarP(&cfg.App.Port, "port", "p", cfg.App.Port, "port to run server")
	a.BoolVarP(&cfg.App.Verbose, "verbose", "v", false, "enable verbose logging")
}
