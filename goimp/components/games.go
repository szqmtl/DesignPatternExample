package components

func CreateGame(factory Factory) *Maze {
	var aMaze = createMaze()

	var r1 = factory.MakeRoom(1)
	var r2 = factory.MakeRoom(2)

	var d = factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(DirectionNorth, factory.MakeWall())
	r1.SetSide(DirectionEast, d)
	r1.SetSide(DirectionSouth, factory.MakeWall())
	r1.SetSide(DirectionWest, factory.MakeWall())

	r2.SetSide(DirectionNorth, factory.MakeWall())
	r2.SetSide(DirectionEast, factory.MakeWall())
	r2.SetSide(DirectionSouth, factory.MakeWall())
	r2.SetSide(DirectionWest, d)

	return aMaze
}

func BuildGame(builder MazeBuilder) *Maze {
	builder.BuildMaze()
	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}

type WallFunc func() Wall
type RoomFunc func(n int) Room
type DoorFunc func(r1 Room, r2 Room) Door

type MazeGame struct {
	makeWall WallFunc
	makeRoom RoomFunc
	makeDoor DoorFunc
}

func (m MazeGame) Create() *Maze {
	mz := createMaze()

	r1 := m.makeRoom(1)
	r2 := m.makeRoom(2)
	mz.AddRoom(r1)
	mz.AddRoom(r2)

	d := m.makeDoor(r1, r2)

	r1.SetSide(DirectionNorth, m.makeWall())
	r1.SetSide(DirectionEast, d)
	r1.SetSide(DirectionSouth, m.makeWall())
	r1.SetSide(DirectionWest, m.makeWall())

	r2.SetSide(DirectionNorth, m.makeWall())
	r2.SetSide(DirectionEast, m.makeWall())
	r2.SetSide(DirectionSouth, m.makeWall())
	r2.SetSide(DirectionWest, d)

	return mz
}

func MakeSimpleGame() *MazeGame {
	return &MazeGame{makeDoor: CreateSimpleDoor, makeRoom: CreateSimpleRoom, makeWall: CreateSimpleWall}
}

func MakeBombedGame() *MazeGame {
	return &MazeGame{makeDoor: CreateSimpleDoor, makeRoom: createRoomWithBomb, makeWall: CreateBombedWall}
}

func MakeEnchantedGame() *MazeGame {
	return &MazeGame{makeDoor: CreateSpellDoor, makeRoom: CreateEnchantedRoom, makeWall: CreateSimpleWall}
}
