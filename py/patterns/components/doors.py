from .common import MapSite
from .rooms import Room
from .spells import Spell


class Door(MapSite):
    def __init__(self, r1: Room, r2: Room):
        self.__room1 = r1
        self.__room2 = r2

    def enter(self):
        pass

    def get_room1(self):
        return self.__room1

    def get_room2(self):
        return self.__room2

    def __str__(self) -> str:
        return 'simple door between room {} and room {}'.format(self.__room1.get_number(), self.__room2.get_number())


class DoorNeedingSpell(Door):
    def __init__(self, r1: Room, r2: Room, spell: Spell):
        super().__init__(r1, r2)
        self.__spell = spell

    def __str__(self) -> str:
        return ('spell door between room {} and room {} with spell {}'
                .format(self.get_room1().get_number(), self.get_room2().get_number(), self.__spell.word()))
