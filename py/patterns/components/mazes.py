from .rooms import Room


class Maze:
    def __init__(self) -> None:
        self.__rooms: dict[int, Room] = {}

    def add_room(self, r: Room):
        self.__rooms[r.get_number()] = r

    def get_room(self, n: int) -> any:
        if n in self.__rooms:
            return self.__rooms[n]

    def __str__(self):
        s = 'maze:\n'
        for k in self.__rooms:
            s = "{} {}: {}\n".format(s, k, self.__rooms[k])

        return s
