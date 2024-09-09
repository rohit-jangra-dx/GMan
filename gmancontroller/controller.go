package gmancontroller

import (
	"Gman/gman"
	"Gman/grid"
	"fmt"
)


type Controller struct {
	gman *gman.Gman
}

// constructor for the Controller
func CreateController(g *gman.Gman) Controller {
	return Controller{
		gman: g,
	}
}

func (c *Controller) MoveGmanToDestination(destination grid.Point) int {

	directionList, err := findDirectionsOfDestination(c.gman.Origin,c.gman.Direction, destination)

	if err != nil {
		fmt.Printf("%s", err)
		return 0
	}

	for _, turn := range directionList {
		turnGman(c.gman, turn)
		moveGman(c.gman, destination)
	}
	return c.gman.Power
}
