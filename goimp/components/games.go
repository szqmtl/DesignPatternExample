package components

type MazeGame struct{}

func (MazeGame) Create(factory Factory) *Maze {
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

func (MazeGame) Build(builder MazeBuilder) *Maze {
	builder.BuildMaze()
	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}
