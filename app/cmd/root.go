// Package cmd contains commands that can be run from the command line
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultCfgFile = "config.yaml"

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "ingrid-coding-assignment",
		Short: "An implementation of the small web service to assess candidate's development skills",
		Long: `A small web service that takes the source and a list of destinations
and returns a list of routes between source and each destination.`,
	}
)

// Execute executes the root command.
func Execute() error {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultCfgFile, "config file")
	rootCmd.AddCommand(serveCmd)

	return rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
