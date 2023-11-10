package place

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Place struct {
	sync.Mutex

	// Other places connected with this place
	connections []PlaceConnection

	// Contains the things that are in this place
	//thingsInPlace map[uuid.UUID]Thing
	thingsInPlace map[ThingType]map[uuid.UUID]Thing
}

func NewPlace() Place {
	return Place{
		connections:   make([]PlaceConnection, 0),
		thingsInPlace: make(map[ThingType]map[uuid.UUID]Thing),
	}
}

// AddThing adds a thing to the place
func (p *Place) AddConnection(place *Place, weight float64) {
	place.Lock()
	defer place.Unlock()

	newCOnnection := NewConnection(p, place, weight)

	p.connections = append(p.connections, newCOnnection)
}

// AddThing adds a thing to the place
func (p *Place) AddThing(thing Thing) {
	p.Lock()
	defer p.Unlock()

	thingType := thing.GetType()

	_, ok := p.thingsInPlace[thingType]
	if !ok {
		p.thingsInPlace[thingType] = make(map[uuid.UUID]Thing)
	}
	p.thingsInPlace[thingType][thing.GetID()] = thing
	thing.SetPlace(p)
}

// Removes a specific thing from the place
// Return true if the thing was removed
func (p *Place) GetThing(thing Thing) bool {
	p.Lock()
	defer p.Unlock()

	// See if the thing is in place
	_, ok := p.thingsInPlace[thing.GetType()][thing.GetID()]

	if ok {
		delete(p.thingsInPlace[thing.GetType()], thing.GetID())
		thing.SetPlace(nil)
		return true
	}
	return false
}

// Returns all the things that are in the place without removing them
func (p *Place) LookAllThings() []Thing {
	p.Lock()
	defer p.Unlock()

	things := make([]Thing, 0)
	for thingsType := range p.thingsInPlace {
		things = append(things, p.LookThingsType(thingsType)...)
	}

	return things
}

// Return the number of things of one type that are in the place
func (p *Place) CountThingsType(thingType ThingType) int {
	p.Lock()
	defer p.Unlock()

	things, ok := p.thingsInPlace[thingType]
	if ok {
		return len(things)
	}
	return 0
}

// Returns all the things types that are in the place without removing them.
func (p *Place) LookThingsType(thingType ThingType) []Thing {
	p.Lock()
	defer p.Unlock()

	things := make([]Thing, 0, len(p.thingsInPlace[thingType]))
	for _, thing := range p.thingsInPlace[thingType] {
		things = append(things, thing)
	}

	return things
}

// Remove the first thing of the type from the place
func (p *Place) GetOneThingType(thingType ThingType) Thing {
	p.Lock()
	defer p.Unlock()

	things, ok := p.thingsInPlace[thingType]
	if ok {
		for thingID, thing := range things {
			delete(p.thingsInPlace[thing.GetType()], thingID)
			thing.SetPlace(nil)
			return thing
		}
	}

	return nil
}

// Remove the things of a specific type from the place
func (p *Place) GetAllThingsType(thingType ThingType) []Thing {
	p.Lock()
	defer p.Unlock()

	things, ok := p.thingsInPlace[thingType]
	if ok {
		// Copy all the things
		thingsCopy := make([]Thing, 0, len(things))
		for _, thing := range things {
			thingsCopy = append(thingsCopy, thing)
		}

		// Remove all the things
		for thingID, thing := range things {
			delete(p.thingsInPlace[thing.GetType()], thingID)
			thing.SetPlace(nil)
		}

		return thingsCopy
	}

	return nil
}

// HasThing returns true if the thing is in the place
func (p *Place) HasThing(thing Thing) bool {
	p.Lock()
	defer p.Unlock()

	_, ok := p.thingsInPlace[thing.GetType()][thing.GetID()]

	return ok
}

// GetRandomConnection returns a random connection
func (p *Place) GetRandomConnection() (PlaceConnection, error) {
	p.Lock()
	defer p.Unlock()

	if len(p.connections) == 0 {
		return NewEmptyConnection(), ErrorNoConnection
	}

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomIndex := r1.Intn(len(p.connections))

	return p.connections[randomIndex], nil
}

// GetConnection returns the connection at the index
func (p *Place) GetConnectionNumber(index int) (PlaceConnection, error) {
	p.Lock()
	defer p.Unlock()

	if index >= len(p.connections) {
		return NewEmptyConnection(), ErrorInvalidConnectionIndex
	}

	return p.connections[index], nil
}

// GetConnection returns the connection with a specific place.
func (p *Place) GetConnection(dest *Place) (PlaceConnection, error) {
	p.Lock()
	defer p.Unlock()

	for _, conn := range p.connections {
		if conn.GetDestination() == dest {
			return conn, nil
		}
	}

	return NewEmptyConnection(), ErrorNoConnection
}

// GetConnections returns the connections of the place
func (p *Place) GetConnections() []PlaceConnection {
	p.Lock()
	defer p.Unlock()

	return p.connections
}

// Print the things in the place
func (p *Place) PrintThingsInPlace() string {
	var strBuider strings.Builder

	for thingsType, things := range p.thingsInPlace {
		thingStr, err := thingsType.String()
		if err != nil {
			return ""
		}
		thingNumber := len(things)
		strBuider.WriteString("\n\r")
		strBuider.WriteString(thingStr)
		strBuider.WriteString(": ")
		strBuider.WriteString(strconv.Itoa(thingNumber))
	}

	return strBuider.String()
}
