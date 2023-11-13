package things

import (
	"sync"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type TimeQuantum struct {
	sync.Mutex

	id    uuid.UUID
	Type  place.ThingType
	place *place.Place
}

func NewTimeQuantum() *TimeQuantum {
	return &TimeQuantum{
		Mutex: sync.Mutex{},
		id:    uuid.New(),
		Type:  place.TimeQuantum,
		place: nil,
	}
}

func (timeQ *TimeQuantum) GetID() uuid.UUID {
	return timeQ.id
}
func (timeQ *TimeQuantum) GetType() place.ThingType {
	timeQ.Lock()
	defer timeQ.Unlock()

	return timeQ.Type
}

func (timeQ *TimeQuantum) GetPlace() *place.Place {
	timeQ.Lock()
	defer timeQ.Unlock()

	return timeQ.place
}

// This function is not to be used to add or remove a thing from a place
func (timeQ *TimeQuantum) SetPlace(newPlace *place.Place) {
	//TO DO: Confirm if it is not needed to lock before setting the place
	//timeQ.Lock()
	//defer timeQ.Unlock()

	timeQ.place = newPlace
}
