from .builder import MazeBuilder
from .doors import Door, DoorNeedingSpell
from .factories import MazeFactory
from .mazes import Maze
from .common import Direction
from .rooms import Room, RoomWithABomb, EnchantedRoom
from .spells import Spell
from .walls import Wall, BombedWall


def build(builder: MazeBuilder) -> Maze:
    builder.build_maze()
    builder.build_room(1)
    builder.build_room(2)
    builder.build_door(1, 2)

    return builder.get_maze()


def create(factory: MazeFactory) -> Maze:
    a_maze = Maze()

    r1 = factory.make_room(1)
    r2 = factory.make_room(2)

    d = factory.make_door(r1, r2)

    a_maze.add_room(r1)
    a_maze.add_room(r2)

    r1.set_side(Direction.North, factory.make_wall())
    r1.set_side(Direction.East, d)
    r1.set_side(Direction.South, factory.make_wall())
    r1.set_side(Direction.West, factory.make_wall())

    r2.set_side(Direction.North, factory.make_wall())
    r2.set_side(Direction.East, factory.make_wall())
    r2.set_side(Direction.South, factory.make_wall())
    r2.set_side(Direction.West, d)

    return a_maze


class MazeGame:
    def create(self) -> Maze:
        m = self.make_maze()

        r1 = self.make_room(1)
        r2 = self.make_room(2)

        d = self.make_door(r1, r2)

        m.add_room(r1)
        m.add_room(r2)

        r1.set_side(Direction.North, self.make_wall())
        r1.set_side(Direction.East, d)
        r1.set_side(Direction.South, self.make_wall())
        r1.set_side(Direction.West, self.make_wall())

        r2.set_side(Direction.North, self.make_wall())
        r2.set_side(Direction.East, self.make_wall())
        r2.set_side(Direction.South, self.make_wall())
        r2.set_side(Direction.West, d)

        return m

    def make_maze(self) -> Maze:
        return Maze()

    def make_room(self, n: int) -> Room:
        return Room(n)

    def make_wall(self) -> Wall:
        return Wall()

    def make_door(self, r1: Room, r2: Room) -> Door:
        return Door(r1, r2)


class BombedMazeGame(MazeGame):
    def make_wall(self) -> Wall:
        return BombedWall()

    def make_room(self, n: int) -> Room:
        return RoomWithABomb(n)


class EnchantedMazeGame(MazeGame):
    def make_room(self, n: int) -> Room:
        return EnchantedRoom(n, Spell("game spell"))

    def make_door(self, r1: Room, r2: Room) -> Door:
        return DoorNeedingSpell(r1, r2, Spell('spell game door'))


