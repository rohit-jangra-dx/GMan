package gmancontroller

import (
	"Gman/configs"
	"Gman/gman"
)

// method to turn the gman
func turnGman(g *gman.Gman, direction configs.Direction) {
	currentDirection := g.Origin.D
	diff := abs(int(direction)-int(currentDirection))

	switch {
	case diff == 2:
		g.Turn(gman.Left)
		g.Turn(gman.Left)
	
	case direction > currentDirection:

		if currentDirection == configs.North {
			g.Turn(gman.Left)
		} else {
			g.Turn(gman.Right)
		}
	
	case direction < currentDirection:

		if direction == configs.North {
			g.Turn(gman.Right)
		} else {
			g.Turn(gman.Left)
		}
	}

}

// it would calculate the steps need to take in particular direction
func deduceSteps(origin gman.Point, x int, y int) int {
	if origin.D == configs.East || origin.D == configs.West {
		return abs(origin.X - x)
	}
	return abs(origin.Y - y)
}

func moveGman(g *gman.Gman, x int, y int) {
	steps := deduceSteps(g.Origin, x, y)
	g.Move(steps)
}
