package cmd

import (
	"api-gateway/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:              "api-gateway",
	Short:            "The api-gateway Service!",
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
