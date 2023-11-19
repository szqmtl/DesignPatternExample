from abc import ABC, abstractmethod
from enum import Enum


class MapSite(ABC):
    @abstractmethod
    def enter(self):
        ...


class Direction(Enum):
    North = 0
    South = 1
    East = 2
    West = 3
