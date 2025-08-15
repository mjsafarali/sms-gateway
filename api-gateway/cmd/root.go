package cmd

import (
	"api-gateway/internal/config"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "api-gateway",
	Short: "The api-gateway Service!",
}

func initConfig() {
	configPath, err := rootCmd.Flags().GetString("config")
	if err != nil {
		log.Fatalf("cant get config path: %s", err)
	}

	// Load config file
	config.Init(configPath)

	logLevel := config.GetInt("LOGGER.LEVEL")
	log.SetLevel(log.Level(logLevel))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
