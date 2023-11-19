package components

type MazeGame struct{}

func (MazeGame) CreateGame(factory Factory) *Maze {
	var aMaze = createMaze()

	var r1 = factory.MakeRoom(1)
	var r2 = factory.MakeRoom(2)

	var d = factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	factory.SetRoomSide(r1, DirectionNorth, factory.MakeWall())
	factory.SetRoomSide(r1, DirectionEast, d)
	factory.SetRoomSide(r1, DirectionSouth, factory.MakeWall())
	factory.SetRoomSide(r1, DirectionWest, factory.MakeWall())

	factory.SetRoomSide(r2, DirectionNorth, factory.MakeWall())
	factory.SetRoomSide(r2, DirectionEast, factory.MakeWall())
	factory.SetRoomSide(r2, DirectionSouth, factory.MakeWall())
	factory.SetRoomSide(r2, DirectionWest, d)

	return aMaze
}
