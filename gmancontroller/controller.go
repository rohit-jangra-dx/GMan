package gmancontroller

import (
	"Gman/gman"
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

func (c *Controller) MoveGmanToDestination(x int, y int) int {
	directionList, err := findDirectionsOfDestination(c.gman.Origin, x, y)

	if err != nil {
		fmt.Printf("%s", err)
		return 0
	}

	for _, turn := range directionList {
		turnGman(c.gman, turn)
		moveGman(c.gman, x, y)
	}
	return c.gman.Power
}
