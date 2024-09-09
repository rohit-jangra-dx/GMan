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
