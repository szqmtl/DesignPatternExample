package components

type MapSite interface {
	Enter()
}

type MapSitePointer[MP MapSite] interface {
	*MP
}

type Direction int

const (
	DirectionNorth Direction = iota + 0
	DirectionSouth
	DirectionEast
	DirectionWest
)

var directions = [...]string{
	"North",
	"South",
	"East",
	"West",
}

func (d Direction) String() string {
	return directions[d]
}
