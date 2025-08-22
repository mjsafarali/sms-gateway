package cmd

import (
	"github.com/spf13/cobra"
	"money/internal/app"
	"money/internal/cmq"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start the worker.",
	Run:   workerFunc,
}

func workerFunc(_ *cobra.Command, _ []string) {
	app.WithGracefulShutdown()
	app.WithDatabase()
	app.WithNats()
	app.WithServices()
	app.WithRepositories()

	cmqShutdownRequest := make(chan struct{})
	cmqShutdownReady := cmq.
		NewConsumer().
		Start().
		WaitForSignals(cmqShutdownRequest)

	<-app.A.Ctx.Done()
	cmqShutdownRequest <- struct{}{}
	<-cmqShutdownReady
}
