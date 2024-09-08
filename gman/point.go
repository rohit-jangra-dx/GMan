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

func createPoint(x int, y int, d string) (Point, error){
	switch d {
	case "E":
		return Point{x, y, configs.East}, nil
	case "W":
		return Point{x, y, configs.West}, nil
	case "N":
		return Point{x, y, configs.North}, nil
	case "S":
		return Point{x, y, configs.South}, nil
	default:
		return Point{}, fmt.Errorf("wrong direction argument given")
	}

}
