//go:build http

package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cook-book/internal/container"
	"log"
	"os"
	"os/signal"
	"time"
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
	ctx, cancelFunc := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		cancelFunc()
	}()

	e := echo.New()

	// DI container input action!
	err := container.Inject(ctx, e)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		<-ctx.Done()

		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()

		_ = e.Shutdown(ctx)
	}()

	// Runs the http
	log.Printf("%v\n", e)
	log.Printf("Server version '%s' is running on port '%s'\n", container.GitCommit, port)
	log.Fatal(e.Start(":" + port))
}
