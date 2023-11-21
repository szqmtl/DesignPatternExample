from .builder import MazeBuilder
from .factories import MazeFactory
from .mazes import Maze
from .common import Direction


class MazeGame:
    def create(self, factory: MazeFactory) -> Maze:
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

    def build(self, builder: MazeBuilder) -> Maze:
        builder.build_maze()
        builder.build_room(1)
        builder.build_room(2)
        builder.build_door(1, 2)

        return builder.get_maze()

