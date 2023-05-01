package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

// usageTemplate is the template made to indicate how to the CLI must be use
const usageTemplate = `%s %s

Usage:

  %s <command> [arguments]

The commands are:
  
  version
`

// BinaryName is the name of the executable file
var BinaryName string

func init() {
	_, BinaryName = path.Split(os.Args[0])
}

// Configuration contains the settings to builds an instance of CLI interface
type Configuration struct {
	// Version indicates the compilation version for the CLI
	Version string
	// Description describes the use that have the CLI
	Description string
	// FlagSet initial flag set
	FlagSet *flag.FlagSet
	// Commands contains a hash map to identify and process the subcommands
	Commands map[string]Command
}

// Command defines a function to handle command line inputs for sub-commands
type Command func(context.Context, ...string) error

// New builds a CLI based on the Configuration
func New(config Configuration) CLI {
	return &cli{
		version:     config.Version,
		description: config.Description,
		flagSet:     config.FlagSet,
		commands:    config.Commands,
	}
}

// CLI command line interface
type CLI interface {
	// Execute runs the command line interface to interpret the
	Execute(context.Context) error
}

type cli struct {
	version     string
	description string
	flagSet     *flag.FlagSet
	commands    map[string]Command
}

// usage prints how to use the CLI
func (c *cli) usage() {
	usage := usageTemplate

	for cmd := range c.commands {
		usage += "  " + cmd + "\n"
	}

	log.Printf(usage, BinaryName, c.description, BinaryName)
}

func (c *cli) Execute(ctx context.Context) error {
	if !c.flagSet.Parsed() {
		c.flagSet.Parse(os.Args[1:])
	}

	args := c.flagSet.Args()

	if len(args) < 1 {
		c.usage()
		return nil
	}

	subcommand := args[0]

	if subcommand == "version" {
		log.Printf("%s version %s\n", BinaryName, c.version)
		return nil
	}

	command, ok := c.commands[subcommand]
	if !ok {
		return fmt.Errorf("invalid sub-command: \"%s\" does not exists", subcommand)
	}

	return command(ctx, append([]string{subcommand}, args[1:]...)...)
}
