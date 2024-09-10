package grid

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

const DirectionLen = 4

// takes args as string and returns a valid direction
func CreateDirection(arg string) (Direction, error) {

	switch arg {
	case "E":
		return East, nil
	case "W":
		return West, nil
	case "N":
		return North, nil
	case "S":
		return South, nil
	default:
		return 0, fmt.Errorf("wrong direction argument given")
	}
}

// takes -1 or 1 to rotate direction clockwise or counterclockwise
func (d *Direction) Rotate(turn int) {
	if turn > 1 || turn < -1 {
		fmt.Println("wrong argument passed, either 1 or -1 would be applicable")
		return
	}

	newDirectionValue := (int(*d) + turn + DirectionLen) % DirectionLen

	*d = Direction(newDirectionValue)
}
