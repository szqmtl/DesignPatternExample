package components

import "fmt"

type Wall interface {
	MapSite
}

type SimpleWall struct{}

func createSimpleWall() *SimpleWall {
	return new(SimpleWall)
}

func (w SimpleWall) Enter() {
	fmt.Println("crossing a simple wall")
}

func (w SimpleWall) String() string {
	return "SimpleWall{}"
}

type BombedWall struct {
}

func createBombedWall() *BombedWall {
	return new(BombedWall)
}

func (w BombedWall) Enter() {
	fmt.Println("crossing a bombed wall")
}

func (w BombedWall) String() string {
	return "BombedWall{}"
}
