package components

import "fmt"

type Room interface {
	MapSite
	GetNumber() int
	SetSide(Direction, MapSite)
}

func CreateSimpleRoom(n int) Room {
	return &SimpleRoom{roomNumber: n, sides: make(map[Direction]MapSite)}
}

func createSimpleRoom(n int) *SimpleRoom {
	return &SimpleRoom{roomNumber: n, sides: make(map[Direction]MapSite)}
}

type SimpleRoom struct {
	roomNumber int
	sides      map[Direction]MapSite
}

func (r SimpleRoom) GetNumber() int {
	return r.roomNumber
}

func (r SimpleRoom) GetSide(d Direction) MapSite {
	return r.sides[d]
}

func (r SimpleRoom) SetSide(d Direction, s MapSite) {
	r.sides[d] = s
}

func (r SimpleRoom) Enter() {
	fmt.Printf("entering a simple room")
}

func (r SimpleRoom) String() string {
	s := fmt.Sprintf("simple room: %v, %v\nsides: ", r.roomNumber, len(r.sides))
	for k, v := range r.sides {
		s += fmt.Sprintf("%s: %s, ", k, v)
	}
	return s + "\n"
}

func createRoomWithBomb(n int) Room {
	return &RoomWithBomb{roomNumber: n, sides: make(map[Direction]MapSite)}
}

type RoomWithBomb struct {
	roomNumber int
	sides      map[Direction]MapSite
}

func (r RoomWithBomb) GetNumber() int {
	return r.roomNumber
}

func (r *RoomWithBomb) GetSide(d Direction) MapSite {
	return r.sides[d]
}

func (r *RoomWithBomb) SetSide(d Direction, s MapSite) {
	r.sides[d] = s
}

func (r RoomWithBomb) Enter() {
	fmt.Printf("entering a room")
}

func (r RoomWithBomb) String() string {
	s := fmt.Sprintf("bomb room: %v, %v\nsides: ", r.roomNumber, len(r.sides))
	for k, v := range r.sides {
		s += fmt.Sprintf("%s: %s, ", k, v)
	}
	return s + "\n"
}

func CreateEnchantedRoom(n int) Room {
	return &EnchantedRoom{roomNumber: n, spell: Spell{word: "magic, magic"}, sides: make(map[Direction]MapSite)}
}

func createEnchantedRoom(n int, s Spell) Room {
	return &EnchantedRoom{roomNumber: n, spell: s, sides: make(map[Direction]MapSite)}
}

type EnchantedRoom struct {
	roomNumber int
	sides      map[Direction]MapSite
	spell      Spell
}

func (r EnchantedRoom) GetNumber() int {
	return r.roomNumber
}

func (r *EnchantedRoom) GetSide(d Direction) MapSite {
	return r.sides[d]
}

func (r EnchantedRoom) SetSide(d Direction, s MapSite) {
	r.sides[d] = s
}

func (r EnchantedRoom) Enter() {
	fmt.Printf("entering a room")
}

func (r EnchantedRoom) String() string {
	s := fmt.Sprintf("enchanted room: %v, %v\nsides: ", r.roomNumber, len(r.sides))
	for k, v := range r.sides {
		s += fmt.Sprintf("%s: %s, ", k, v)
	}
	return s + "\n"
}
