package components

import "fmt"

type Door interface {
	MapSite
}

type SimpleDoor struct {
	room1  Room
	room2  Room
	isOpen bool
}

func CreateSimpleDoor(r1, r2 Room) Door {
	return &SimpleDoor{room1: r1, room2: r2}
}

func createSimpleDoor(r1, r2 Room) *SimpleDoor {
	return &SimpleDoor{room1: r1, room2: r2}
}

func (d SimpleDoor) Enter() {
	fmt.Printf("entering room 1: %T from room 2: %T, isOpen: %v", d.room1, d.room2, d.isOpen)
}

func (d SimpleDoor) String() string {
	return fmt.Sprintf("Simple Door{room 1: %T, room 2: %T, isOpen: %v}", d.room1, d.room2, d.isOpen)
}

type DoorNeedingSpell struct {
	room1  Room
	room2  Room
	isOpen bool
	spell  Spell
}

func CreateSpellDoor(r1, r2 Room) Door {
	return &DoorNeedingSpell{room1: r1, room2: r2, spell: Spell{"big magic~"}}
}

func createSpellDoor(r1, r2 Room) *DoorNeedingSpell {
	return &DoorNeedingSpell{room1: r1, room2: r2, spell: Spell{"big magic~"}}
}

func (d DoorNeedingSpell) Enter() {
	fmt.Printf("entering room 1: %T from room 2: %T, isOpen: %v", d.room1, d.room2, d.isOpen)
}

func (d DoorNeedingSpell) String() string {
	return fmt.Sprintf("Spell Door{room 1: %T, room 2: %T, isOpen: %v}", d.room1, d.room2, d.isOpen)
}
