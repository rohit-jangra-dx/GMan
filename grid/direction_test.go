package grid

import (
	"testing"
)

func TestCreateDirection(t *testing.T) {
	// normal case
	directionStrArrays := []string{"E", "W", "S", "N"}
	expectedResult := []Direction{East, West, South, North}

	for i, direction := range directionStrArrays {
		result, err := CreateDirection(direction)
		if err != nil {
			t.Error(err.Error())
		}
		expect(expectedResult[i], result, t)
	}

	// invalid date
	_, err := CreateDirection("b")
	if err == nil {
		t.Error("expected invalid direction error got nothing")
	}

}

func TestRotate(t *testing.T) {
	d := North

	//testing the left rotatiion
	d.Rotate(-1)
	expect(West, d, t)

	// testing the right rotation
	d.Rotate(1)
	expect(North, d, t)

	// // testing circular motion
	d.Rotate(-1)
	d.Rotate(1)
	expect(North, d, t)

	// // testing wrong values
	d.Rotate(2)
	expect(North, d, t)

	d.Rotate(-2)
	expect(North, d, t)
}

// to log error
func expect(expected any, got any, t *testing.T) {
	if expected != got {
		t.Errorf("expected {%v} got {%v}", expected, got)
	}
}
