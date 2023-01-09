package main

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/internal/container"
	"github.com/yael-castro/cb-search-engine-api/internal/domain/generic/server"
	"log"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.TODO()
	log.SetFlags(log.Flags() | log.Llongfile)

	var s server.Server

	err := container.New().Inject(&s)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(s.Serve(ctx, ":"+port))
}
