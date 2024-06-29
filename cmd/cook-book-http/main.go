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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		const defaultPort = "8080"
		port = defaultPort
	}

	// Sets logger flags
	log.SetFlags(log.Flags() | log.Llongfile)

	// Creates main context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	e := echo.New()

	// DI container in action!
	// TODO: close connections with external repositories
	err := container.Inject(ctx, e)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		<-ctx.Done()

		const gracePeriod = 10 * time.Second
		ctx, cancelFunc := context.WithTimeout(context.Background(), gracePeriod)
		defer cancelFunc()

		_ = e.Shutdown(ctx)
	}()

	log.Printf("Server version '%s' is running on port '%s'\n", container.GitCommit, port)
	_ = e.Start(":" + port)
}
