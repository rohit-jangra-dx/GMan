package parsers

import (
	"Gman/configs"
	"Gman/gman"
	"Gman/gmancontroller"
	"Gman/grid"
	"fmt"
	"strconv"
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

// command handlers
func (c *CommandContext) handleSourceCommand(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("invalid number of arguments for SOURCE command")
	}

	x, y, d := args[0], args[1], args[2]

	//converting args into valid point
	origin, err := grid.CreatePoint(x, y)
	if err != nil {
		return err
	}
	direction, err := grid.CreateDirection(d)
	if err != nil {
		return err
	}

	g := gman.CreateGman(origin, direction, configs.GameConfigInstance)
	c.gman = &g
	return nil
}

func (c *CommandContext) handleDestinationCommand(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments for DESTINATION command")
	}

	x, y := args[0], args[1]

	destination, err := grid.CreatePoint(x, y)
	if err != nil {
		return err
	}

	ctrl := gmancontroller.CreateController(c.gman)
	c.controller = &ctrl

	// moving it
	c.controller.MoveGmanToDestination(destination)
	return nil
}

func (c *CommandContext) handlePrintPowerCommand() error {
	fmt.Printf("POWER  %d", c.gman.Power)
	return nil
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

// some helper functions
func convertStringToInt(args []string) ([]int, error) {
	intArray := make([]int, len(args))

	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid argument, could not convert str to int")
		}
		intArray[i] = num
	}
	return intArray, nil
}
