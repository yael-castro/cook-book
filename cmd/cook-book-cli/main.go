//go:build cli

package main

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/yael-castro/cook-book/internal/container"
	"os"
	"os/signal"
)

func main() {
	// Building main context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// DI container input action!
	var cmd cobra.Command

	// DI container in action!
	// TODO: close connections with external repositories
	err := container.Inject(ctx, &cmd)
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	err = cmd.ExecuteContext(ctx) // TODO: evaluate error for exit code
	if err != nil {
		cmd.PrintErr(err)
		return
	}
}
