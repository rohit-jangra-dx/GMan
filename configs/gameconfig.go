package configs

type GameConfig struct {
	CostPerMove int
	CostPerTurn int
	GridSize    int
	InitialPower int
}

var GameConfigInstance = GameConfig{
	CostPerMove: 10,
	CostPerTurn: 5,
	GridSize:    6,
	InitialPower: 200,
}
