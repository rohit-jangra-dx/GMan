package parsers

import (
	"Gman/configs"
	"Gman/gman"
	"Gman/gmancontroller"
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
	coordinatesStr, direction := args[:2], args[2]
	// converting them to int first
	coordinatesInt, err := convertStringToInt(coordinatesStr)
	if err != nil {
		return err
	}

	// checking whether coordinates are valid or not
	flag := configs.IsCoordinatesValid(coordinatesInt[0], coordinatesInt[1])
	if !flag {
		return fmt.Errorf("wrong coordinates are given according to grid")
	}

	g, err := gman.CreateGman(coordinatesInt[0], coordinatesInt[1], direction, 200, configs.GameConfigInstance)
	if err != nil {
		return err
	}
	c.gman = &g
	return nil
}

func (c *CommandContext) handleDestinationCommand(args []string) error{
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments for DESTINATION command")
	}

	coordinatesStr := args

	// converting them to int
	coordinatesInt, err := convertStringToInt(coordinatesStr)
	if err != nil {
		return err
	}
	
	// checking for coordinates validity
	flag := configs.IsCoordinatesValid(coordinatesInt[0],coordinatesInt[1])
	if !flag {
		return fmt.Errorf("wrong coordinated are given according to grid")
	}
	// now creating controller to move gman to target location
	ctrl := gmancontroller.CreateController(c.gman)
	c.controller = &ctrl

	// moving it
	c.controller.MoveGmanToDestination(coordinatesInt[0],coordinatesInt[1])
	return nil
}

func (c *CommandContext) handlePrintPowerCommand() error {
	fmt.Printf("POWER  %d",c.gman.Power)
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
