package components

import (
	"fmt"
)

type Factory interface {
	MakeWall() Wall
	MakeDoor(Room, Room) Door
	MakeRoom(int) Room
	SetRoomSide(Room, Direction, MapSite)
}

type MazeFactory struct{}

func (m MazeFactory) MakeWall() Wall {
	return createSimpleWall()
}

func (m MazeFactory) MakeRoom(n int) Room {
	return createSimpleRoom(n)
}

func (m MazeFactory) MakeDoor(r1, r2 Room) Door {
	return createSimpleDoor(r1, r2)
}

func (m MazeFactory) SetRoomSide(r Room, direction Direction, site MapSite) {
	switch r := r.(type) {
	case *SimpleRoom:
		r.SetSide(direction, site)
	default:
		fmt.Printf("unsupport room type(%T)", r)
	}
}

type EnchantedMazeFactory struct{}

func (m EnchantedMazeFactory) MakeWall() Wall {
	return createSimpleWall()
}

func (m EnchantedMazeFactory) MakeRoom(n int) Room {
	return createEnchantedRoom(n, Spell{word: "magic~"})
}

func (m EnchantedMazeFactory) MakeDoor(r1, r2 Room) Door {
	return createSpellDoor(r1, r2)
}

func (m EnchantedMazeFactory) SetRoomSide(r Room, direction Direction, site MapSite) {
	switch r := r.(type) {
	case *EnchantedRoom:
		r.SetSide(direction, site)
	default:
		fmt.Printf("unsupport room type(%T)", r)
	}
}

type BombedMazeFactory struct{}

func (m BombedMazeFactory) SetRoomSide(room Room, direction Direction, site MapSite) {
	switch room := room.(type) {
	case *RoomWithBomb:
		room.SetSide(direction, site)
	default:
		fmt.Printf("unsupport room type(%T)", room)
	}
}

func (m BombedMazeFactory) MakeWall() Wall {
	return createBombedWall()
}

func (m BombedMazeFactory) MakeRoom(n int) Room {
	return createRoomWithBomb(n)
}

func (m BombedMazeFactory) MakeDoor(r1, r2 Room) Door {
	return createSimpleDoor(r1, r2)
}
