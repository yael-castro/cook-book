package main

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/cmd/server/container"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server"
	"log"
	"os"
	"os/signal"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Sets logger flags
	log.SetFlags(log.Flags() | log.Llongfile)

	// Creates main context
	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		cancel()
	}()

	// DI container in action!
	var s server.Server

	err := container.New().Inject(&s)
	if err != nil {
		log.Fatal(err)
	}

	// Runs the server
	log.Printf("Server version '%s' is running on port '%s'\n", container.GitCommit, port)
	log.Println(s.Serve(ctx, ":"+port))
}
