package main

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/cmd/cli/container"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/cli"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	checkpoint := time.Now()

	// Sets logger flags
	log.SetFlags(0)

	// Creates main context
	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		cancel()
	}()

	// DI container in action!
	var c cli.CLI

	err := container.New().Inject(&c)
	if err != nil {
		log.Fatal(err)
	}

	// Runs the CLI
	err = c.Execute(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("\nexecution time:", time.Now().Sub(checkpoint))
}
