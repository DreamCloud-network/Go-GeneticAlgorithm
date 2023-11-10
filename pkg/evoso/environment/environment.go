package environment

import (
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Environment struct {
	sync.Mutex

	id uuid.UUID

	// Contains the things that are in this place
	thingsInPlace map[ThingType]map[*Thing]*Thing

	// Create a map of mutex to control access conforming the type of thing
	thingsMutex map[ThingType]*sync.Mutex
}

func NewEnvironment() *Environment {
	return &Environment{
		Mutex:         sync.Mutex{},
		id:            uuid.New(),
		thingsInPlace: make(map[ThingType]map[*Thing]*Thing),
		thingsMutex:   make(map[ThingType]*sync.Mutex),
	}
}

// Return if the hash of a certain thing type have already been created.
func (env *Environment) hasThingTypeHash(thingType ThingType) bool {
	env.Lock()
	_, ok := env.thingsInPlace[thingType]
	env.Unlock()

	return ok
}

// HasThing returns true if the thing is in the place
func (env *Environment) HasThing(thing *Thing) bool {
	if thing == nil {
		return false
	}

	thingType := thing.GetType()

	if env.hasThingTypeHash(thingType) {
		env.thingsMutex[thingType].Lock()
		_, ok := env.thingsInPlace[thingType][thing]
		env.thingsMutex[thingType].Unlock()
		return ok
	}

	return false
}

// AddThing adds a thing to the place.
// Returns true if the thing was added or if it already was in place.
func (env *Environment) AddThing(thing *Thing) bool {
	if thing == nil {
		return false
	}

	thingType := thing.GetType()

	if !env.hasThingTypeHash(thingType) {
		env.Lock()
		env.thingsInPlace[thingType] = make(map[*Thing]*Thing)
		env.thingsMutex[thingType] = &sync.Mutex{}
		env.Unlock()
	}

	env.thingsMutex[thingType].Lock()
	env.thingsInPlace[thingType][thing] = thing
	env.thingsMutex[thingType].Unlock()

	thing.setExternalEnvironment(env)

	return true
}

// Removes a specific thing from the place
// Return true if the thing was removed
func (env *Environment) GetThing(thing *Thing) bool {
	if thing == nil {
		return false
	}

	thingType := thing.GetType()

	if env.hasThingTypeHash(thingType) {
		env.thingsMutex[thingType].Lock()
		defer env.thingsMutex[thingType].Unlock()

		_, ok := env.thingsInPlace[thingType][thing]
		if ok {
			delete(env.thingsInPlace[thingType], thing)
			thing.setExternalEnvironment(nil)
			return true
		}
	}

	return false
}

// Remove the first thing of the type from the place
func (env *Environment) GetOneThingType(thingType ThingType) *Thing {
	env.Lock()
	things, ok := env.thingsInPlace[thingType]
	env.Unlock()

	if ok {
		env.thingsMutex[thingType].Lock()
		defer env.thingsMutex[thingType].Unlock()

		for thing := range things {
			delete(env.thingsInPlace[thingType], thing)
			thing.setExternalEnvironment(nil)
			return thing
		}
	}

	return nil
}

// Remove all the things of a certain type from the place
func (env *Environment) GetAllThingsType(thingType ThingType) []*Thing {
	env.Lock()
	things, ok := env.thingsInPlace[thingType]
	env.Unlock()

	if ok && (len(things) > 0) {
		env.thingsMutex[thingType].Lock()
		defer env.thingsMutex[thingType].Unlock()

		thingsVet := make([]*Thing, 0, len(things))
		for thing := range things {
			// Add thing to the vector
			thingsVet = append(thingsVet, thing)

			// Remove thing from place
			delete(env.thingsInPlace[thing.GetType()], thing)
			thing.setExternalEnvironment(nil)
		}

		return thingsVet
	}

	return nil
}

// Returns all the things types that are in the place without removing them.
func (env *Environment) LookThingsType(thingType ThingType) []*Thing {
	env.Lock()
	things, ok := env.thingsInPlace[thingType]
	env.Unlock()

	if ok && (len(things) > 0) {
		env.thingsMutex[thingType].Lock()
		defer env.thingsMutex[thingType].Unlock()

		thingsVet := make([]*Thing, len(things))
		pos := 0
		for _, thing := range things {
			thingsVet[pos] = thing
			pos++
		}
		return thingsVet
	}

	return nil
}

// Return the number of things of one type that are in the place
func (env *Environment) NumThingsType(thingType ThingType) int {
	env.Lock()
	defer env.Unlock()

	things, ok := env.thingsInPlace[thingType]
	if ok {
		return len(things)
	}
	return 0
}

// Print the things in the place
func (env *Environment) PrintThingsInPlace() string {
	env.Lock()
	defer env.Unlock()

	var strBuider strings.Builder

	for thingsType, things := range env.thingsInPlace {
		thingNumber := len(things)
		strBuider.WriteString("\n\r")
		strBuider.WriteString(thingsType.String())
		strBuider.WriteString(": ")
		strBuider.WriteString(strconv.Itoa(thingNumber))
	}

	return strBuider.String()
}
