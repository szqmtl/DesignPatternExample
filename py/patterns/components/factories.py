from .spells import Spell
from .walls import Wall, BombedWall
from .rooms import Room, EnchantedRoom, RoomWithBomb
from .doors import Door, DoorNeedingSpell


class MazeFactory:
    def make_wall(self) -> Wall:
        return Wall()

    def make_room(self, n: int) -> Room:
        return Room(n)

    def make_door(self, r1: Room, r2: Room) -> Door:
        return Door(r1, r2)


class EnchantedMazeFactory(MazeFactory):
    def make_wall(self) -> Wall:
        return Wall()

    def make_room(self, n: int) -> Room:
        return EnchantedRoom(n)

    def make_door(self, r1: Room, r2: Room) -> Door:
        return DoorNeedingSpell(r1, r2, Spell("magic word"))


class BombedMazeFactory(MazeFactory):
    def make_wall(self) -> Wall:
        return BombedWall()

    def make_room(self, n: int) -> Room:
        return RoomWithBomb(n)

    def make_door(self, r1: Room, r2: Room) -> Door:
        return Door(r1, r2)

