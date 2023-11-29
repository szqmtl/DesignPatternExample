package patterns

import (
	"fmt"

	"goimp/components"
)

func RunAbstractFactory() {
	m1 := components.CreateGame(components.MazeFactory{})
	fmt.Printf("maze 1: %v\n", m1)

	m2 := components.CreateGame(components.EnchantedMazeFactory{})
	fmt.Printf("maze 2: %v\n", m2)

	m3 := components.CreateGame(components.BombedMazeFactory{})
	fmt.Printf("maze 3: %v\n", m3)
}

func RunBuilder() {
	m1 := components.BuildGame(&components.SimpleMazeBuilder{})
	fmt.Printf("maze 1: %v\n", m1)

	m2 := components.BuildGame(&components.StandardMazeBuilder{})
	fmt.Printf("maze 2: %v\n", m2)

	b := &components.CountingMazeBuilder{}
	components.BuildGame(b)
	rooms, doors := b.GetCounts()
	fmt.Printf("maze rooms:%v, doors: %v\n", rooms, doors)
}

func RunFactoryMethods() {
	m1 := components.MakeSimpleGame().Create()
	fmt.Printf("maze 1: %v\n", m1)

	m2 := components.MakeBombedGame().Create()
	fmt.Printf("maze 2: %v\n", m2)

	m3 := components.MakeEnchantedGame().Create()
	fmt.Printf("maze 3: %v\n", m3)
}
