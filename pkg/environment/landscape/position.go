package landscape

type Direction interface {
	String() string
}

type Position interface {
	String() string
	GetThings() []Thing
	GetLandscape() Landscape
	HasThing(thing Thing) bool
}
