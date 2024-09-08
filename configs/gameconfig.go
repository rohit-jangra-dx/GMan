package configs

type GameConfig struct {
	CostPerMove int
	CostPerTurn int
	GridSize    int
}

var GameConfigInstance = GameConfig{
	CostPerMove: 10,
	CostPerTurn: 5,
	GridSize:    6,
}

func IsCoordinatesValid(x int, y int)bool {
	if x > GameConfigInstance.GridSize || x < 0{
		return false
	} else if  y > GameConfigInstance.GridSize || y < 0{
		return false
	}
	return true
}