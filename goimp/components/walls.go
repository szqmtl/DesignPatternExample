package components

import "fmt"

type Wall interface {
	MapSite
}

type SimpleWall struct{}

func CreateSimpleWall() Wall {
	return new(SimpleWall)
}

func createSimpleWall() *SimpleWall {
	return new(SimpleWall)
}

func (w SimpleWall) Enter() {
	fmt.Println("crossing a simple walls")
}

func (w SimpleWall) String() string {
	return "SimpleWall{}"
}

type BombedWall struct {
}

func CreateBombedWall() Wall {
	return new(BombedWall)
}

func createBombedWall() *BombedWall {
	return new(BombedWall)
}

func (w BombedWall) Enter() {
	fmt.Println("crossing a bombed walls")
}

func (w BombedWall) String() string {
	return "BombedWall{}"
}
