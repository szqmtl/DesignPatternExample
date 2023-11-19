from .common import MapSite


class Wall(MapSite):
    def enter(self):
        pass

    def __str__(self) -> str:
        return "simple wall{}"


class BombedWall(Wall):
    def enter(self):
        pass

    def __str__(self):
        return "bombed wall{}"
