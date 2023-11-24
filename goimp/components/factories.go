package components

type Factory interface {
	MakeWall() Wall
	MakeDoor(Room, Room) Door
	MakeRoom(int) Room
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

type BombedMazeFactory struct{}

func (m BombedMazeFactory) MakeWall() Wall {
	return createBombedWall()
}

func (m BombedMazeFactory) MakeRoom(n int) Room {
	return createRoomWithBomb(n)
}

func (m BombedMazeFactory) MakeDoor(r1, r2 Room) Door {
	return createSimpleDoor(r1, r2)
}
