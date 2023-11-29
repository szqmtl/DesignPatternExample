from .components.builder import MazeBuilder, StandardMazeBuilder, CountingMazeBuilder
from .components.games import MazeGame, create, build, EnchantedMazeGame, BombedMazeGame
from .components.factories import MazeFactory, EnchantedMazeFactory, BombedMazeFactory


def run_abstract_factory():
    print("starting the game")
    game = MazeGame()

    m1 = create(MazeFactory())
    print(m1)

    m2 = create(EnchantedMazeFactory())
    print(m2)

    m3 = create(BombedMazeFactory())
    print(m3)


def run_builder():
    game = MazeGame()
    builder1 = MazeBuilder()

    build(builder1)
    m1 = builder1.get_maze()
    print(m1)

    builder2 = StandardMazeBuilder()

    build(builder2)
    m2 = builder2.get_maze()
    print(m2)

    builder3 = CountingMazeBuilder()

    build(builder3)
    r, d = builder3.get_counts()
    print("maze 3: room count: {}, door count: {}".format(r, d))

def run_factory_method():
    m1 = MazeGame().create()
    print(m1)

    m2 = EnchantedMazeGame().create()
    print(m2)

    m3 = BombedMazeGame().create()
    print(m3)
