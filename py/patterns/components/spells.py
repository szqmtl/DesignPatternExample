class Spell:
    def __init__(self, word: str):
        self.__word = word

    def word(self):
        return self.__word

    def __str__(self):
        return 'spell: ' + self.__word
