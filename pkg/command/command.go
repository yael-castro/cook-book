package command

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
)

func init() {
	_, BinaryName = path.Split(os.Args[0])
}

var BinaryName string

const (
	// usageTemplate is the template made to indicate how to the CLI must be use
	usageTemplate = `%s - %s

Usage:

  %s <command> [options]

Help:

  %s <command> -help

The commands are:

`

	helpTemplate = `%s %s - %s

Usage:

  %s %s [options]

Options:

`
)

// Command defines a function to handle command line inputs for sub-commands
type Command interface {
	Name() string
	Description() string
	Execute(context.Context, ...string) error
}

type Configuration struct {
	Version     string
	Description string
	Commands    []Command
}

func New(config Configuration) Command {
	// Setting sub-command "version"
	config.Commands = append([]Command{Version(config.Version)}, config.Commands...)

	commanders := make(map[string]Command)

	for _, commander := range config.Commands {
		commanders[commander.Name()] = commander
	}

	return &command{
		flags:       flag.NewFlagSet(BinaryName, flag.ContinueOnError),
		description: config.Description,
		commands:    commanders,
	}
}

type command struct {
	description string
	commands    map[string]Command
	flags       *flag.FlagSet
}

func (c command) Name() string {
	return c.flags.Name()
}

func (c command) Description() string {
	return c.description
}

func (c command) Execute(ctx context.Context, args ...string) error {
	if len(args) < 2 {
		c.Usage()
		return nil
	}

	subcommand := args[1]

	command, ok := c.commands[subcommand]
	if !ok {
		c.Usage()
		return fmt.Errorf("invalid sub-command: \"%s\" does not exists", subcommand)
	}

	err := command.Execute(ctx, append([]string{subcommand}, args[2:]...)...)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil // Ignore error
		}

		return err
	}

	return nil
}

// Usage prints how to use the command
func (c command) Usage() {
	usage := usageTemplate

	for _, command := range c.commands {
		usage += fmt.Sprintf("  %s - %s\n", command.Name(), command.Description())
	}

	usage += "\n"
	fmt.Printf(usage, BinaryName, c.description, BinaryName, BinaryName)
}

func Version(version string) Command {
	return versionCommand{
		version: version,
	}
}

type versionCommand struct {
	version string
}

func (v versionCommand) Name() string {
	return "version"
}

func (v versionCommand) Description() string {
	return "prints the current version"
}

func (v versionCommand) Execute(context.Context, ...string) error {
	fmt.Println(v.version)
	return nil
}
