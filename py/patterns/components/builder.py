from patterns.components.common import Direction
from patterns.components.doors import Door
from patterns.components.mazes import Maze
from patterns.components.rooms import Room
from patterns.components.walls import Wall


class MazeBuilder:
    def __init__(self):
        self.__maze = None

    def build_maze(self):
        self.__maze = Maze()

    def build_room(self, n: int):
        self.__maze.add_room(Room(n))

    def build_door(self, from_room: int, to_room: int):
        r1 = self.__maze.get_room(from_room)
        r2 = self.__maze.get_room(to_room)
        d = Door(r1, r2)

        r1.set_side(Direction.South, d)
        r2.set_side(Direction.North, d)

    def get_maze(self):
        return self.__maze


class StandardMazeBuilder(MazeBuilder):
    def build_maze(self):
        super().build_maze()

    def build_room(self, n: int):
        if super().get_maze().get_room(n) is None:
            room = Room(n)
            super().get_maze().add_room(room)

            room.set_side(Direction.North, Wall())
            room.set_side(Direction.South, Wall())
            room.set_side(Direction.East, Wall())
            room.set_side(Direction.West, Wall())

    def build_door(self, from_room: int, to_room: int):
        r1 = super().get_maze().get_room(from_room)
        r2 = super().get_maze().get_room(to_room)

        d = Door(r1, r2)

        r1.set_side(self.__common_wall(r1, r2), d)
        r2.set_side(self.__common_wall(r2, r1), d)

    def __common_wall(self, r1: Room, r2: Room) -> Direction:
        if (r1.get_number() - r2.get_number()) > 0:
            return Direction.West
        else:
            return Direction.East


class CountingMazeBuilder(MazeBuilder):
    def __init__(self):
        super().__init__()
        self.__rooms = 0
        self.__doors = 0

    def build_door(self, from_room: int, to_room: int):
        self.__doors += 1

    def  build_room(self, n: int):
        self.__rooms += 1

    def get_counts(self) -> (int, int):
        return self.__rooms, self.__doors

