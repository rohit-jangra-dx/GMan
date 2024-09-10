package gmancontroller

import (
	"Gman/grid"
	"fmt"
	"sort"
)

// type alias
type DirectionList []grid.Direction

// generates a map that gives u the directions of destination point relative to origin point
func getTurnMap() map[grid.Point]DirectionList {
	turnMap := make(map[grid.Point]DirectionList)

	turnMap[grid.Point{X: 0, Y: 1}] = DirectionList{grid.North}
	turnMap[grid.Point{X: 0, Y: -1}] = DirectionList{grid.South}

	turnMap[grid.Point{X: 1, Y: 0}] = DirectionList{grid.East}
	turnMap[grid.Point{X: -1, Y: 0}] = DirectionList{grid.West}

	turnMap[grid.Point{X: 1, Y: 1}] = DirectionList{grid.East, grid.North}
	turnMap[grid.Point{X: -1, Y: -1}] = DirectionList{grid.West, grid.South}

	turnMap[grid.Point{X: -1, Y: 1}] = DirectionList{grid.West, grid.North}
	turnMap[grid.Point{X: 1, Y: -1}] = DirectionList{grid.East, grid.South}

	return turnMap
}

// returns the list of direction you need to turn in order to reach to destination (optimized)
func findDirectionsOfDestination(origin grid.Point, originDirection grid.Direction, destination grid.Point) (DirectionList, error) {

	// used for the mapping
	directionalDiff := destination.GetNormalizedDifference(origin)

	//get the map and return the array matching the keys
	turnMap := getTurnMap()
	directionList, exists := turnMap[directionalDiff]
	if !exists {
		return nil, fmt.Errorf("map is broken, check the cases again")
	}

	//this will make sure that least turn possible are taken to move to distance
	ozList := optimizeDirectionList(directionList, originDirection)

	return ozList, nil
}

// sorts the directions list in order from closest to the current direction origin is facing to the furthest
func optimizeDirectionList(directions DirectionList, currentDirection grid.Direction) DirectionList {
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
func minMod(d1, d2 grid.Direction) int {
	diff := int((d1 - d2 + grid.DirectionLen) % grid.DirectionLen)
	return min(diff, grid.DirectionLen-diff)
}

// returns minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
