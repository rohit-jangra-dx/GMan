package main

import (
	"Gman/configs"
	"Gman/gman"
	"Gman/gmancontroller"
	"fmt"
)

func main() {
	fmt.Println("Hello, Modules!")
	point := gman.Point{
		X: 3,
		Y: 6,
		D: configs.North,
	}
	gmanInstance := gman.CreateGman(point, 200, configs.GameConfigInstance)
	controller := gmancontroller.CreateController(&gmanInstance)

	power := controller.MoveGmanToDestination(1, 0)
	fmt.Println(power)
}
