package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	port    int
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "xi-server",
	Short: "XI HTTP server",
	Long:  "XI HTTP server — the main backend for the XI platform",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server on port:", port)
		if verbose {
			fmt.Println("Verbose logging enabled")
		}
		// TODO: call internal/server.Start(port)
	},
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default $HOME/.xi-server.yaml)")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port to run server")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")

	viper.SetEnvPrefix("XI_SERVER")
	viper.AutomaticEnv() // read env variables like XI_SERVER_PORT
}

// initConfig reads in config file
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, _ := os.UserHomeDir()
		viper.AddConfigPath(home)
		viper.SetConfigName(".xi-server")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
