//go:build http

package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/cook-book/internal/container"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

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
	err := container.Inject(ctx, e)
	if err != nil {
		log.Println(err)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()

		const gracePeriod = 10 * time.Second
		ctx, cancelFunc := context.WithTimeout(context.Background(), gracePeriod)
		defer cancelFunc()

		_ = e.Shutdown(ctx)
		log.Println("Shutting down")

		db, err := container.MongoDatabase(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		err = db.Client().Disconnect(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Disconnected from database")
	}()

	log.Printf("Server version '%s' is running on port '%s'\n", container.GitCommit, port)
	_ = e.Start(":" + port)
}
