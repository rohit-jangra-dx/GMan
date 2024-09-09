package gman

import (
	"Gman/configs"
	"Gman/grid"
	"fmt"
)


type Turn string

const (
	Left  Turn = "Left"
	Right Turn = "Right"
)

type Gman struct {
	Origin     	grid.Point
	Direction 	grid.Direction
	Power      	int
	gameConfig 	configs.GameConfig
}

func CreateGman(p grid.Point, d grid.Direction, gameConfig configs.GameConfig) Gman {

	return Gman{
		Origin:     p,
		Direction:  d,
		Power:      gameConfig.InitialPower,
		gameConfig: gameConfig,
	}
}

func (g *Gman) Turn(direction Turn) {

	turn_increment := map[Turn]int{
		Left:  -1,
		Right: 1,
	}

	increment, exists := turn_increment[direction]
	if !exists {
		fmt.Println("Invalid turn direction argument passed!")
	}

	g.Direction.Rotate(increment)
	g.Power -= g.gameConfig.CostPerTurn
}

// logic for moving in current direction
func (g *Gman) moveNorth(steps int) {
	g.Origin.Y += steps
}

func (g *Gman) moveSouth(steps int) {
	g.Origin.Y -= steps
}

func (g *Gman) moveEast(steps int) {
	g.Origin.X += steps
}

func (g *Gman) moveWest(steps int) {
	g.Origin.X -= steps
}

func (g *Gman) getMoveMap() map[grid.Direction]func(int) {
	return map[grid.Direction]func(int){
		grid.North: g.moveNorth,
		grid.East:  g.moveEast,
		grid.South: g.moveSouth,
		grid.West:  g.moveWest,
	}
}

func (g *Gman) Move(steps int) {
	moveMap := g.getMoveMap()

	moveFunc, exists := moveMap[g.Direction]
	if !exists {
		fmt.Println("Invalid direction!")
		return
	}

	moveFunc(steps)

	g.Power -= (steps * g.gameConfig.CostPerMove)
}
