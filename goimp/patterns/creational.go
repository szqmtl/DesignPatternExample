package patterns

import (
	"fmt"

	"goimp/components"
)

func RunAbstractFactory() {
	game := components.MazeGame{}
	m1 := game.CreateGame(components.MazeFactory{})
	fmt.Printf("maze 1: %v\n", m1)

	m2 := game.CreateGame(components.EnchantedMazeFactory{})
	fmt.Printf("maze 2: %v\n", m2)

	m3 := game.CreateGame(components.BombedMazeFactory{})
	fmt.Printf("maze 3: %v\n", m3)
}
