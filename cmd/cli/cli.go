package main

import (
	"context"
	"fmt"
	"github.com/yael-castro/cb-search-engine-api/cmd/cli/container"
	"github.com/yael-castro/cb-search-engine-api/pkg/command"
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
	var c command.Command

	err := container.Inject(ctx, &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Runs the CLI
	err = c.Execute(ctx, os.Args...)
	if err != nil {
		fmt.Println(err)
	}
}
