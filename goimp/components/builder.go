package components

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(n int)
	BuildDoor(r1, r2 int)
	GetMaze() *Maze
}

type SimpleMazeBuilder struct {
	maze *Maze
}

func (m *SimpleMazeBuilder) BuildMaze() {
	m.maze = createMaze()
}

func (m *SimpleMazeBuilder) BuildRoom(n int) {
	m.maze.AddRoom(createSimpleRoom(n))
}

func (m *SimpleMazeBuilder) BuildDoor(r1, r2 int) {
	room1 := m.maze.RoomNo(r1)
	room2 := m.maze.RoomNo(r2)
	d := createSimpleDoor(room1, room2)

	room1.SetSide(DirectionSouth, d)
	room2.SetSide(DirectionNorth, d)
}

func (m *SimpleMazeBuilder) GetMaze() *Maze {
	return m.maze
}

type StandardMazeBuilder struct {
	maze *Maze
}

func (s *StandardMazeBuilder) GetMaze() *Maze {
	return s.maze
}

func (s *StandardMazeBuilder) BuildMaze() {
	s.maze = createMaze()
}

func (s *StandardMazeBuilder) BuildRoom(n int) {
	if s.maze.RoomNo(n) == nil {
		r := createSimpleRoom(n)

		r.SetSide(DirectionNorth, createSimpleWall())
		r.SetSide(DirectionSouth, createSimpleWall())
		r.SetSide(DirectionEast, createSimpleWall())
		r.SetSide(DirectionWest, createSimpleWall())

		s.maze.AddRoom(r)
	}
}

func (s *StandardMazeBuilder) BuildDoor(fr, tr int) {
	r1 := s.maze.RoomNo(fr)
	r2 := s.maze.RoomNo(tr)

	d := createSimpleDoor(r1, r2)
	r1.SetSide(s.commonWall(r1, r2), d)
	r2.SetSide(s.commonWall(r2, r1), d)
}

func (s *StandardMazeBuilder) commonWall(r1, r2 Room) Direction {
	if (r1.GetNumber() - r2.GetNumber()) > 0 {
		return DirectionWest
	} else {
		return DirectionEast
	}
}

type CountingMazeBuilder struct {
	maze  *Maze
	rooms int
	doors int
}

func (c *CountingMazeBuilder) BuildMaze() {
	c.maze = createMaze()
}

func (c *CountingMazeBuilder) GetMaze() *Maze {
	return c.maze
}

func (c *CountingMazeBuilder) BuildDoor(_, _ int) {
	c.doors++
}

func (c *CountingMazeBuilder) BuildRoom(_ int) {
	c.rooms++
}

func (c *CountingMazeBuilder) GetCounts() (int, int) {
	return c.rooms, c.doors
}
