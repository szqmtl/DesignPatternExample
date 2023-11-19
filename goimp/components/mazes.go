package components

import "fmt"

type Maze struct {
	rooms map[int]Room
}

func createMaze() *Maze {
	return &Maze{
		rooms: make(map[int]Room, 0),
	}
}

func (m *Maze) AddRoom(r Room) {
	m.rooms[r.GetNumber()] = r
}

func (m *Maze) RoomNo(n int) Room {
	return m.rooms[n]
}

func (m *Maze) String() string {
	s := "\nmaze: \n"
	for k, v := range m.rooms {
		s += fmt.Sprintf("%d: %s\n", k, v)
	}
	return s
}
