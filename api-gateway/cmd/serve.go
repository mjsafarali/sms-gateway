package cmd

import (
	"api-gateway/internal/app"
	internalHttp "api-gateway/internal/http"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the service.",
	Run:   startFunc,
}

func startFunc(_ *cobra.Command, _ []string) {
	app.WithGracefulShutdown()
	app.WithDatabase()
	app.WithRedis()

	internalHttp.NewServer().Serve()
}
