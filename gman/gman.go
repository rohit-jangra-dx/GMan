package gman

import (
	"Gman/configs"
	"fmt"
)

type Point struct {
	X int
	Y int
	D configs.Direction
}

type Turn string

const (
	Left  Turn = "Left"
	Right Turn = "Right"
)

type Gman struct {
	Origin     Point
	Power      int
	gameConfig configs.GameConfig
}

func CreateGman(Origin Point, Power int, gameConfig configs.GameConfig) Gman {
	return Gman{
		Origin:     Origin,
		Power:      Power,
		gameConfig: gameConfig,
	}
}

func (g *Gman) Turn(direction Turn) {

	turn_increment := map[Turn]int{
		"Left":  -1,
		"Right": 1,
	}

	increment, exists := turn_increment[direction]
	if !exists {
		fmt.Println("Invalid turn direction argument passed!")
	}

	g.Origin.D.Rotate(increment)
	g.Power -= g.gameConfig.CostPerTurn
}

// logic for moving in current direction
func (g *Gman) moveNorth(steps int) {
	g.Origin.Y += steps
	fmt.Printf("Moved North by %d steps. new Origin are %d-%d\n", steps, g.Origin.X, g.Origin.Y)
}

func (g *Gman) moveSouth(steps int) {
	g.Origin.Y -= steps
	fmt.Printf("Moved North by %d steps. new Origin are %d-%d\n", steps, g.Origin.X, g.Origin.Y)
}

func (g *Gman) moveEast(steps int) {
	g.Origin.X += steps
	fmt.Printf("Moved North by %d steps. new Origin are %d-%d\n", steps, g.Origin.X, g.Origin.Y)
}

func (g *Gman) moveWest(steps int) {
	g.Origin.X -= steps
	fmt.Printf("Moved North by %d steps. new Origin are %d-%d\n", steps, g.Origin.X, g.Origin.Y)
}

func (g *Gman) getMoveMap() map[configs.Direction]func(int) {
	return map[configs.Direction]func(int){
		configs.North: g.moveNorth,
		configs.East:  g.moveEast,
		configs.South: g.moveSouth,
		configs.West:  g.moveWest,
	}
}

func (g *Gman) Move(steps int) {
	moveMap := g.getMoveMap()

	moveFunc, exists := moveMap[g.Origin.D]
	if !exists {
		fmt.Println("Invalid direction!")
		return
	}

	moveFunc(steps)

	g.Power -= (steps * g.gameConfig.CostPerMove)
	fmt.Printf("Remaining Power: %d\n", g.Power)
}
