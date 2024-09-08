package gman

import (
	"Gman/configs"
	"fmt"
)


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

func CreateGman(x int, y int, d string, Power int, gameConfig configs.GameConfig) (Gman, error) {
	origin, err := createPoint(x, y, d)
	if err != nil {
		fmt.Println(err.Error())
		return Gman{}, fmt.Errorf("gman could not be created successfully")
	}
	return Gman{
		Origin:     origin,
		Power:      Power,
		gameConfig: gameConfig,
	}, nil
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
}
