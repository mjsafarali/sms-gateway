package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hermes/internal/config"
	"os"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:              "hermes",
	Short:            "The hermes Service!",
	PersistentPreRun: preRun,
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file path",
	)

	rootCmd.AddCommand(startCmd)
}

func preRun(_ *cobra.Command, _ []string) {
	// Load config file
	config.Init(cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
