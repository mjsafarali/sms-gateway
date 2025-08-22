package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"money/internal/config"
	"os"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:              "money",
	Short:            "The money Service!",
	PersistentPreRun: preRun,
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file path",
	)

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(workerCmd)
}

func preRun(_ *cobra.Command, _ []string) {
	config.Init(cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
