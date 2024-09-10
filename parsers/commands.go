package parsers

import (
	"Gman/gman"
	"Gman/gmancontroller"
	"fmt"
	"strings"
)

type CommandContext struct {
	gman       *gman.Gman
	controller *gmancontroller.Controller
}

func CreateCommandContext() CommandContext {
	return CommandContext{
		gman:       nil,
		controller: nil,
	}
}

func (c *CommandContext) ExecuteCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return fmt.Errorf("no command found")
	}

	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "SOURCE":
		c.handleSourceCommand(args)
	case "DESTINATION":
		c.handleDestinationCommand(args)
	case "PRINT_POWER":
		c.handlePrintPowerCommand()
	default:
		return fmt.Errorf("invalid command found")
	}
	return nil
}
