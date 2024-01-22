package main

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/yael-castro/cb-search-engine-api/cmd/cli/container"
	"os"
	"os/signal"
)

func main() {
	// Creates main context
	ctx, cancelFunc := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		cancelFunc()
	}()

	// DI container input action!
	var cmd cobra.Command

	err := container.Inject(ctx, &cmd)
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	_ = cmd.ExecuteContext(ctx) // TODO: evaluate error for exit code
}
