package parsers

import (
	"fmt"
	"Gman/grid"
	"Gman/configs"
	"Gman/gman"
	"Gman/gmancontroller"
)

const (
	ArgsLengthForSourceCommand = 3
	ArgsLengthForDestinationCommand = 2
)

// command handlers
func (c *CommandContext) handleSourceCommand(args []string) error {
	if len(args) != ArgsLengthForSourceCommand {
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
	if len(args) != ArgsLengthForDestinationCommand {
		return fmt.Errorf("invalid number of arguments for DESTINATION command")
	}

	x, y := args[0], args[1]

	destination, err := grid.CreatePoint(x, y)
	if err != nil {
		return err
	}

	ctrl := gmancontroller.CreateController(c.gman)
	c.controller = &ctrl

	// moving gman
	c.controller.MoveGmanToDestination(destination)
	return nil
}

func (c *CommandContext) handlePrintPowerCommand() error {
	fmt.Printf("POWER  %d", c.gman.Power)
	return nil
}