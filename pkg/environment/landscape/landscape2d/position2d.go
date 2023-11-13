package landscape2d

import (
	"errors"
	"strconv"
	"sync"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/environment/landscape"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/environment/thingstype"
)

var (
	ErrorInvalidDirection = errors.New("invalid direction")
)

type Directions byte

const (
	Center Directions = iota
	East
	NorthEast
	North
	NorthWest
	West
	SouthWest
	South
	SouthEast
)

func (d Directions) String() string {
	switch d {
	case Center:
		return "Center"
	case East:
		return "East"
	case NorthEast:
		return "NorthEast"
	case North:
		return "North"
	case NorthWest:
		return "NorthWest"
	case West:
		return "West"
	case SouthWest:
		return "SouthWest"
	case South:
		return "South"
	case SouthEast:
		return "SouthEast"

	default:
		return ErrorInvalidDirection.Error()
	}
}

type Position2D struct {
	sync.Mutex

	X int
	Y int

	nextPositions []*Position2D

	things    []landscape.Thing
	landscape landscape.Landscape
}

func NewPosition2D(x int, y int) *Position2D {
	return &Position2D{
		Mutex:         sync.Mutex{},
		X:             x,
		Y:             y,
		nextPositions: make([]*Position2D, 9),
		things:        nil,
		landscape:     nil,
	}
}

//------------------------------------------------------------
// Position interface implementation

func (position *Position2D) String() string {
	return "(" + strconv.Itoa(position.X) + "," + strconv.Itoa(position.Y) + ")"
}

// Return a copy of all tginhs in the position.
func (position *Position2D) GetThings() []landscape.Thing {
	var things []landscape.Thing

	position.Lock()
	defer position.Unlock()

	things = append(things, position.things...)

	return things
}

func (position *Position2D) GetLandscape() landscape.Landscape {
	return position.landscape
}

func (position *Position2D) HasThing(thing landscape.Thing) bool {
	position.Lock()
	defer position.Unlock()

	for _, thingInPosition := range position.things {
		if thingInPosition.GetID() == thing.GetID() {
			return true
		}
	}

	return false
}

//------------------------------------------------------------

func (position *Position2D) HasThingType(thingType thingstype.ThingType) bool {
	position.Lock()
	defer position.Unlock()

	for _, thingInPosition := range position.things {
		if thingInPosition.GetType() == thingType {
			return true
		}
	}

	return false
}

func (position *Position2D) AddThing(thing landscape.Thing) {
	position.Lock()
	defer position.Unlock()

	position.things = append(position.things, thing)
	thing.SetPosition(position)
}

func (position *Position2D) GetThingsByType(thingType thingstype.ThingType) []landscape.Thing {
	position.Lock()
	defer position.Unlock()

	var things []landscape.Thing

	for _, thing := range position.things {
		if thing.GetType() == thingType {
			things = append(things, thing)
		}
	}

	return things
}

// RemoveThingByID remove a thing from the position using its ID
func (position *Position2D) RemoveThing(thingToRemove landscape.Thing) landscape.Thing {
	position.Lock()
	defer position.Unlock()

	var removedThing landscape.Thing

	for cont, thingInPosition := range position.things {
		if thingInPosition.GetID() == thingToRemove.GetID() {
			removedThing = thingInPosition
			removedThing.SetPosition(nil)
			position.things = append(position.things[:cont], position.things[cont+1:]...)
			return removedThing
		}
	}

	return nil
}

// RemoveThingByType remove the first found thing from the position using its type
func (position *Position2D) RemoveThingByType(thingType thingstype.ThingType) landscape.Thing {
	position.Lock()
	defer position.Unlock()

	for pos, thingInPosition := range position.things {
		if thingInPosition.GetType() == thingType {
			removedThing := thingInPosition
			removedThing.SetPosition(nil)
			position.things = append(position.things[:pos], position.things[pos+1:]...)
			return removedThing
		}
	}

	return nil
}

func (position *Position2D) CleanAllThings() {
	position.Lock()
	defer position.Unlock()

	position.things = nil
}

// GetNextPosition return the position in the direction and distance specified
func (position *Position2D) GetRelativePosition(direction Directions, distance int) *Position2D {
	position.Lock()
	defer position.Unlock()

	if distance == 0 {
		return position
	}

	var relativePosition *Position2D

	for cont := 0; cont < distance; cont++ {
		relativePosition = position.nextPositions[direction]
		if relativePosition == nil {
			return relativePosition
		}
	}

	return relativePosition
}

// MoveThingToDirection move a thing to the direction specified.
// if the next direction is nil, the thing rains in the same position.
func (position *Position2D) MoveThingToDirection(thing landscape.Thing, direction Directions) error {
	//position.Lock()
	//defer position.Unlock()

	actualPosition := thing.GetPosition().(*Position2D)

	newPosition := actualPosition.GetRelativePosition(direction, 1)

	if newPosition == nil {
		//directionStr, _ := direction.String()
		//log.Println("landscape2d.Landscape2D.MoveTring - error movinght thing to ", directionStr)
		return ErrorInvalidDirection
	}

	actualPosition.RemoveThing(thing)
	newPosition.AddThing(thing)

	return nil
}
