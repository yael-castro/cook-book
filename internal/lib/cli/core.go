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

type Commander interface {
	Command(context.Context, ...string) error
}

type Helper interface {
	Help()
}

// CommanderFunc defines a function to handle command line inputs for sub-commands
type CommanderFunc func(context.Context, ...string) error

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
}
