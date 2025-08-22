package cmd

import (
	"github.com/spf13/cobra"
	"money/internal/app"
	internalHttp "money/internal/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the service.",
	Run:   startFunc,
}

func startFunc(_ *cobra.Command, _ []string) {
	app.WithGracefulShutdown()
	app.WithDatabase()
	app.WithRepositories()
	app.WithRedis()
	app.WithServices()

	internalHttp.NewServer().Serve()
}
