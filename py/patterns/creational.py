from .components.builder import MazeBuilder, StandardMazeBuilder, CountingMazeBuilder
from .components.games import MazeGame
from .components.factories import MazeFactory, EnchantedMazeFactory, BombedMazeFactory


def run_abstract_factory():
    print("starting the game")
    game = MazeGame()

    m1 = game.create(MazeFactory())
    print(m1)

    m2 = game.create(EnchantedMazeFactory())
    print(m2)

    m3 = game.create(BombedMazeFactory())
    print(m3)


def run_builder():
    game = MazeGame()
    builder1 = MazeBuilder()

    game.build(builder1)
    m1 = builder1.get_maze()
    print(m1)

    builder2 = StandardMazeBuilder()

    game.build(builder2)
    m2 = builder2.get_maze()
    print(m2)

    builder3 = CountingMazeBuilder()

    game.build(builder3)
    r, d = builder3.get_counts()
    print("maze 3: room count: {}, door count: {}".format(r, d))
