package grid

import (
	"testing"
)

func TestConvertStringToInt(t *testing.T) {

	// normal case
	num, err := convertStringToInt("23")
	if err != nil {
		t.Error(err)
	}
	expect(23, num, t)

	// error case
	_, err = convertStringToInt("x")

	if err == nil {
		t.Error("expected inavlid argument error, but got nothing")
	}

}

func TestIsCoordinatesValid(t *testing.T) {

	//true case
	flag := isCoordinatesValid(10, 10, 10)
	expect(true, flag, t)

	// false cases
	flag = isCoordinatesValid(-1, 10, 10)
	expect(false, flag, t)

	flag = isCoordinatesValid(10, -10, 10)
	expect(false, flag, t)

	flag = isCoordinatesValid(12, 12, 10)
	expect(false, flag, t)
}

func TestCreatePoint(t *testing.T) {

	// normal test
	result, err := CreatePoint("2", "2")
	if err != nil {
		t.Error(err.Error())
	}
	expect(2, result.X, t)
	expect(2, result.Y, t)

	// error cases

	_, err = CreatePoint("b", "2")
	if err == nil {
		t.Error("expected invalid point values error, but got nothing")
	}

	_, err = CreatePoint("10", "10")
	if err == nil {
		t.Error("expected coordniates out of grid error but got nothing")
	}

}

func TestNormalizeDifference(t *testing.T) {

	testCases := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{23, 1},
		{-10, -1},
	}
	result := 0
	for _, testCase := range testCases {
		result = normalizeCoordinate(testCase.in)
		if result != testCase.out {
			t.Errorf("expected {%v} got {%v}", testCase.out, result)
		}
	}
}

func TestGetNormalizedDifference(t *testing.T) {
	point := Point{2, 2}

	testCases := []struct {
		in  Point
		out Point
	}{
		{Point{2, 2}, Point{0, 0}},
		{Point{1, 2}, Point{1, 0}},
		{Point{1, 1}, Point{1, 1}},
		{Point{2, 1}, Point{0, 1}},
		{Point{3, 3}, Point{-1, -1}},
		{Point{2, 3}, Point{0, -1}},
		{Point{3, 2}, Point{-1, 0}},
		{Point{1, 3}, Point{1, -1}},
		{Point{3, 1}, Point{-1, 1}},
	}
	result := Point{0, 0}
	for _, testCase := range testCases {
		result = point.GetNormalizedDifference(testCase.in)
		if result.X != testCase.out.X || result.Y != testCase.out.Y {
			t.Errorf("got {%v, %v} expected {%v, %v}", result.X, result.Y, testCase.out.X, testCase.out.Y)
		}
	}
}
