package main

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/cmd/server/container"
	"log"
	"net/http"
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
	var h http.Handler

	err := container.New().Inject(&h)
	if err != nil {
		log.Fatal(err)
	}

	// Builds the server
	server := http.Server{
		Addr:    ":" + port,
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		server.Close()
	}()

	// Runs the server
	log.Printf("Server version '%s' is running on port '%s'\n", container.GitCommit, port)
	log.Println(server.ListenAndServe())
}
