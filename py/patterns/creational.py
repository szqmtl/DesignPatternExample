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
