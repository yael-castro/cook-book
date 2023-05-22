// Package health used to
package health

import (
	"context"
	server2 "github.com/yael-castro/cb-search-engine-api/pkg/server"
	"net/http"
)

// Ping defines a function to monitor the health to external repositories
type Ping func(context.Context) error

// Config defines a
type Config struct {
	// Version indicates the server version
	Version string
	// PingMap functions to monitor the health to external repositories
	PingMap map[string]Ping
	Logger  server2.Logger
}

// Checker defines a health checker for the server.Server
type Checker interface {
	// Check checks the server health
	Check(w http.ResponseWriter, r *http.Request)
}

// NewChecker builds a health checker for the server based on the Config
func NewChecker(config Config) Checker {
	return &checker{
		Version: config.Version,
		PingMap: config.PingMap,
	}
}

type checker struct {
	Version string
	PingMap map[string]Ping
	Logger  server2.Logger
}

func (c checker) Check(w http.ResponseWriter, r *http.Request) {
	unavailable := false
	services := make(map[string]string)

	for service, ping := range c.PingMap {
		err := ping(r.Context())
		if err != nil {
			if c.Logger != nil {
				c.Logger.Log()
			}

			unavailable = true
			services[service] = "PING"
			continue
		}

		services[service] = "PONG"
	}

	status := http.StatusOK
	if unavailable {
		status = http.StatusServiceUnavailable
	}

	_ = server2.JSON(w, status, map[string]any{
		"version":  c.Version,
		"services": services,
	})
}
