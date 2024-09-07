package gmancontroller

import (
	"Gman/configs"
	"Gman/gman"
	"fmt"
	"sort"
)

// for the turn map key
type Delta struct {
	X int
	Y int
}

// generate the normailzed delta struct key for turnMap
func createDelta(deltaX int, deltaY int) Delta {
	if deltaX != 0 {
		deltaX = deltaX / abs(deltaX)
	} else {
		deltaX = 0
	}
	if deltaY != 0 {
		deltaY = deltaY / abs(deltaY)
	} else {
		deltaY = 0
	}
	return Delta{deltaX, deltaY}
}

// todo: move this function away from this code
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// type alias
type DirectionList []configs.Direction

// generates a map that gives u the directions of destination point relative to origin point
func getTurnMap() map[Delta]DirectionList {
	turnMap := make(map[Delta]DirectionList)

	turnMap[Delta{X: 0, Y: 1}] = DirectionList{configs.North}
	turnMap[Delta{X: 0, Y: -1}] = DirectionList{configs.South}

	turnMap[Delta{X: 1, Y: 0}] = DirectionList{configs.East}
	turnMap[Delta{X: -1, Y: 0}] = DirectionList{configs.West}

	turnMap[Delta{X: 1, Y: 1}] = DirectionList{configs.East, configs.North}
	turnMap[Delta{X: -1, Y: -1}] = DirectionList{configs.West, configs.South}

	turnMap[Delta{X: -1, Y: 1}] = DirectionList{configs.West, configs.North}
	turnMap[Delta{X: 1, Y: -1}] = DirectionList{configs.East, configs.South}

	return turnMap
}

// returns the list of direction you need to turn in order to reach to destination (optimized)
func findDirectionsOfDestination(origin gman.Point, x int, y int) (DirectionList, error) {
	deltaX := x - origin.X
	deltaY := y - origin.Y

	move_diff := createDelta(deltaX, deltaY)

	//get the map and return the array matching the keys
	turnMap := getTurnMap()

	directionList, exists := turnMap[move_diff]
	if !exists {
		return nil, fmt.Errorf("map is broken, check the cases again")
	}
	//this will make sure that least turn possible are taken to move to distance
	ozList := optimizeDirectionList(directionList, origin.D)

	return ozList, nil
}

// sorts the directions list in order from closest to the current direction origin is facing to the furthest
func optimizeDirectionList(directions DirectionList, currentDirection configs.Direction) DirectionList {
	sort.Slice(directions, func(i, j int) bool {
		d1 := directions[i]
		d2 := directions[j]

		diff1 := minMod(d1, currentDirection)
		diff2 := minMod(d2, currentDirection)

		return diff1 < diff2
	})
	return directions
}

// calculates the minimum modular difference betwenn two directions
func minMod(d1, d2 configs.Direction) int {
	diff := int((d1 - d2 + configs.DirectionLen) % configs.DirectionLen)
	return min(diff, configs.DirectionLen-diff)
}

// returns minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
