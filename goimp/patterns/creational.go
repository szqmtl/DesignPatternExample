package patterns

import (
	"fmt"

	"goimp/components"
)

func RunAbstractFactory() {
	game := components.MazeGame{}
	m1 := game.Create(components.MazeFactory{})
	fmt.Printf("maze 1: %v\n", m1)

	m2 := game.Create(components.EnchantedMazeFactory{})
	fmt.Printf("maze 2: %v\n", m2)

	m3 := game.Create(components.BombedMazeFactory{})
	fmt.Printf("maze 3: %v\n", m3)
}

func RunBuilder() {
	game := components.MazeGame{}

	m1 := game.Build(&components.SimpleMazeBuilder{})
	fmt.Printf("maze 1: %v\n", m1)

	m2 := game.Build(&components.StandardMazeBuilder{})
	fmt.Printf("maze 2: %v\n", m2)

	b := &components.CountingMazeBuilder{}
	game.Build(b)
	rooms, doors := b.GetCounts()
	fmt.Printf("maze rooms:%v, doors: %v\n", rooms, doors)
}
