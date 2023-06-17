// Package health used to
package health

import (
	"context"
	"github.com/yael-castro/cb-search-engine-api/pkg/server"
	"log"
	"net/http"
)

// Ping defines a function to monitor the health to external repositories
type Ping func(context.Context) error

// Config defines a
type Config struct {
	// Version indicates the server version
	Version string
	// PingMap functions to monitor the health to external repository connections
	PingMap map[string]Ping
	// Logger logger to sets the
	Logger *log.Logger
}

// Checker defines a health checker for the server.Server
type Checker interface {
	// Check checks the server health
	Check(w http.ResponseWriter, r *http.Request)
}

// NewChecker builds a health checker for the server based on the Config
func NewChecker(config Config) Checker {
	if config.Logger == nil {
		config.Logger = log.Default()
	}

	return &checker{
		Version: config.Version,
		PingMap: config.PingMap,
		logger:  config.Logger,
	}
}

type checker struct {
	Version string
	PingMap map[string]Ping
	logger  *log.Logger
}

func (c checker) Check(w http.ResponseWriter, r *http.Request) {
	unavailable := false
	services := make(map[string]string)

	for service, ping := range c.PingMap {
		err := ping(r.Context())
		if err != nil {
			unavailable = true
			c.logger.Println(err)
			services[service] = "PING"
			continue
		}

		services[service] = "PONG"
	}

	status := http.StatusOK
	if unavailable {
		status = http.StatusServiceUnavailable
	}

	_ = server.JSON(w, status, map[string]any{
		"version":  c.Version,
		"services": services,
	})
}
