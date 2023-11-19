from .factories import MazeFactory
from .mazes import Maze
from .common import Direction


class MazeGame:
    def create(self, factory: MazeFactory) -> Maze:
        aMaze = Maze()

        r1 = factory.make_room(1)
        r2 = factory.make_room(2)

        d = factory.make_door(r1, r2)

        aMaze.add_room(r1)
        aMaze.add_room(r2)

        r1.set_side(Direction.North, factory.make_wall())
        r1.set_side(Direction.East, d)
        r1.set_side(Direction.South, factory.make_wall())
        r1.set_side(Direction.West, factory.make_wall())

        r2.set_side(Direction.North, factory.make_wall())
        r2.set_side(Direction.East, factory.make_wall())
        r2.set_side(Direction.South, factory.make_wall())
        r2.set_side(Direction.West, d)

        return aMaze
