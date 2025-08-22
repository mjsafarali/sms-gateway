package cmd

import (
	"github.com/spf13/cobra"
	"hermes/internal/app"
	"hermes/internal/cmq"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the service.",
	Run:   startFunc,
}

func startFunc(_ *cobra.Command, _ []string) {
	app.WithNats()
	app.WithServices()
	app.WithGracefulShutdown()

	cmqShutdownRequest := make(chan struct{})
	cmqShutdownReady := cmq.
		NewConsumer().
		Start().
		WaitForSignals(cmqShutdownRequest)

	<-app.A.Ctx.Done()
	cmqShutdownRequest <- struct{}{}
	<-cmqShutdownReady
}
