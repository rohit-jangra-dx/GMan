package gmancontroller

import (
	"Gman/gman"
	"Gman/grid"
	"math"
)

// method to turn the gman
func turnGman(g *gman.Gman, direction grid.Direction) {
	currentDirection := g.Direction

	diff := int(math.Abs(float64((int(direction) - int(currentDirection)))))

	switch {
	case diff == 2:
		g.Turn(gman.Left)
		g.Turn(gman.Left)

	case direction > currentDirection:

		if currentDirection == grid.North {
			g.Turn(gman.Left)
		} else {
			g.Turn(gman.Right)
		}

	case direction < currentDirection:

		if direction == grid.North {
			g.Turn(gman.Right)
		} else {
			g.Turn(gman.Left)
		}
	}

}

// it would calculate the steps need to take in particular direction
func deduceSteps(origin grid.Point, originDirection grid.Direction, destination grid.Point) int {
	if originDirection == grid.East || originDirection == grid.West {
		return int(math.Abs(float64(origin.X - destination.X)))
	}
	return int(math.Abs(float64(origin.Y - destination.Y)))
}

func moveGman(g *gman.Gman, destination grid.Point) {
	steps := deduceSteps(g.Origin, g.Direction, destination)
	g.Move(steps)
}
