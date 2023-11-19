package components

type Spell struct {
	word string
}

func (s *Spell) Word(word string) {
	s.word = word
}

func (s Spell) String() string {
	return "spell: " + s.word
}
