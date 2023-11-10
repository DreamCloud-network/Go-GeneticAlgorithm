package environment

import (
	"sync"

	"github.com/google/uuid"
)

type Thing struct {
	sync.Mutex

	id        uuid.UUID
	thingType ThingType

	externalEnvironment *Environment
	internalEnvironment Environment

	object interface{}
}

func NewThing(t ThingType, obj interface{}) *Thing {
	return &Thing{
		Mutex: sync.Mutex{},

		id:        uuid.New(),
		thingType: t,

		externalEnvironment: nil,
		internalEnvironment: *NewEnvironment(),

		object: obj,
	}
}

func (thing *Thing) GetID() uuid.UUID {
	thing.Lock()
	defer thing.Unlock()

	return thing.id
}

func (thing *Thing) SetType(t ThingType) {
	thing.Lock()
	defer thing.Unlock()

	thing.thingType = t
}

func (thing *Thing) GetType() ThingType {
	thing.Lock()
	defer thing.Unlock()

	return thing.thingType
}

func (thing *Thing) setExternalEnvironment(newEnvironment *Environment) {
	thing.Lock()
	defer thing.Unlock()

	thing.externalEnvironment = newEnvironment
}

func (thing *Thing) GetExternalEnvironment() *Environment {
	thing.Lock()
	defer thing.Unlock()

	return thing.externalEnvironment
}

func (thing *Thing) GetInternalEnvironment() *Environment {
	thing.Lock()
	defer thing.Unlock()

	return &thing.internalEnvironment
}

func (thing *Thing) SetObject(obj interface{}) {
	thing.Lock()
	defer thing.Unlock()

	thing.object = obj
}

func (thing *Thing) GetObject() interface{} {
	thing.Lock()
	defer thing.Unlock()

	return thing.object
}
