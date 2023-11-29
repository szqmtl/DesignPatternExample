from .common import MapSite, Direction
from .spells import Spell


class Room(MapSite):
    def __init__(self, n: int) -> None:
        self.__roomNumber = n
        self.__sides: dict[int, MapSite] = {}

    def enter(self):
        pass

    def get_side(self, d: Direction) -> MapSite:
        return self.__sides[d.value]

    def get_sides(self) -> dict[int, MapSite]:
        return self.__sides

    def set_side(self, d: Direction, s: MapSite):
        self.__sides[d.value] = s

    def get_number(self) -> int:
        return self.__roomNumber

    def __str__(self):
        s = 'simple room({}) -'.format(self.__roomNumber)
        for x in self.__sides:
            s = '{} {}: {}'.format(s, Direction(x).name, self.__sides[x])

        return s


class RoomWithABomb(Room):
    def __int__(self, n: int):
        super().__init__(n)

    def __str__(self):
        s = 'boom room({}) -'.format(self.get_number())
        sides = self.get_sides()
        for x in sides:
            s = '{} {}: {}'.format(s, Direction(x).name, sides[x])

        return s


class EnchantedRoom(Room):

    def __init__(self, n: int, spell: Spell):
        super().__init__(n)
        self.__spell = spell

    def __str__(self):
        s = 'enchanted room({}, {}) -'.format(self.get_number(), self.__spell)
        sides = self.get_sides()
        for x in sides:
            s = '{} {}: {}'.format(s, Direction(x).name, sides[x])

        return s
