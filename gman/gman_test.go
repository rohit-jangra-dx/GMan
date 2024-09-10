package gman

import (
	"Gman/configs"
	"Gman/grid"
	"testing"
)

func TestCreateGman(t *testing.T) {
	point := grid.Point{X: 1, Y: 1}
	direction := grid.North
	gameConfig := configs.GameConfigInstance

	gman := CreateGman(point, direction, gameConfig)

	if gman.Origin.X != 1 || gman.Origin.Y != 1 || gman.Direction != grid.North {
		t.Errorf("expected origin {%v, %v} and  direction {%v} found {%v, %v}, {%v}",
			gman.Origin.X, gman.Origin.Y, gman.Direction, point.X, point.Y, direction)
	}
}

func TestTurnGman(t *testing.T) {

	gman := Gman{Origin: grid.Point{X: 1, Y: 1}, Direction: grid.East, Power: configs.GameConfigInstance.InitialPower, gameConfig: configs.GameConfigInstance}

	// normal test case
	gman.Turn(Left)
	if gman.Direction != grid.North || gman.Power == configs.GameConfigInstance.InitialPower {
		t.Errorf("expected {%v}, got {%v}", grid.North, gman.Direction)
	}

	// trying anything else expecept left or right
	gman.Turn("right")
	if gman.Direction != grid.North {
		t.Error("shouldn't not be able to turn coz of invalid argument passed, but it did")
	}

}

func TestMoveGman(t *testing.T) {
	gman := Gman{Origin: grid.Point{X: 1, Y: 1}, Direction: grid.East, Power: configs.GameConfigInstance.InitialPower, gameConfig: configs.GameConfigInstance}

	// moving and turning all direction
	initalPower := configs.GameConfigInstance.InitialPower
	costPerMove := configs.GameConfigInstance.CostPerMove

	gman.Move(2)
	if gman.Origin.X != 3 || gman.Power != (initalPower-(2*costPerMove)) {
		t.Errorf("{%d} ,{%d}", gman.Power, (initalPower - (2 * costPerMove)))
	}

}

func TestMoveGmanBasedOnDirection(t *testing.T) {
	gman := CreateGman(grid.Point{X: 1, Y: 1}, grid.East, configs.GameConfigInstance)
	gman.moveEast(1)
	if gman.Origin.X != 2 {
		t.Errorf("expected {%v} got {%v}", 2, gman.Origin.X)
	}

	gman.moveWest(1)
	if gman.Origin.X != 1 {
		t.Errorf("expected {%v} got {%v}", 1, gman.Origin.X)
	}

	gman.moveNorth(1)
	if gman.Origin.Y != 2 {
		t.Errorf("expected {%v} got {%v}", 2, gman.Origin.X)
	}

	gman.moveSouth(1)
	if gman.Origin.X != 1 {
		t.Errorf("expected {%v} got {%v}", 1, gman.Origin.X)
	}

}
