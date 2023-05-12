package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"
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
	// Commanders contains a hash map to identify and process the subcommands
	Commanders map[string]Commander
}

// New builds a CLI based on the Configuration
func New(config Configuration) CLI {
	return &cli{
		version:     config.Version,
		description: config.Description,
		flagSet:     config.FlagSet,
		commanders:  config.Commanders,
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
	commanders  map[string]Commander
}

// usage prints how to use the CLI
func (c *cli) usage() {
	usage := usageTemplate

	for cmd := range c.commanders {
		usage += "  " + cmd + "\n"
	}

	log.Printf(usage, BinaryName, c.description, BinaryName)
}

func (c *cli) Execute(ctx context.Context) error {
	checkpoint := time.Now()

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

	commander, ok := c.commanders[subcommand]
	if !ok {
		return fmt.Errorf("invalid sub-commander: \"%s\" does not exists", subcommand)
	}

	err := commander.Command(ctx, append([]string{subcommand}, args[1:]...)...)
	if err != nil {
		if helper, ok := commander.(Helper); ok {
			helper.Help()
		}
		return err
	}

	log.Println("\nExecution time:", time.Now().Sub(checkpoint))
	return nil
}
