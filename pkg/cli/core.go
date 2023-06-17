package cli

import (
	"context"
	"flag"
	"log"
)

const helpTemplate = `%s %s is a tool to generate random recipes

Usage:

  %s %s [options]

Options:

`

type Helper interface {
	Help()
}

// Commander defines a function to handle command line inputs for sub-commands
type Commander interface {
	// Command executes the sub-command with a raw string arguments
	Command(context.Context, ...string) error
}

// CommanderFunc functional interface for the Commander port
type CommanderFunc func(context.Context, ...string) error

// Command executes f
func (f CommanderFunc) Command(ctx context.Context, args ...string) error {
	return f(ctx, args...)
}

func Usage(set *flag.FlagSet) {
	log.Printf(
		helpTemplate,
		BinaryName,
		set.Name(),
		BinaryName,
		set.Name(),
	)
	set.PrintDefaults()
	log.Println()
}
