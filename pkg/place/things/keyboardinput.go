package things

import (
	"sync"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type KeyboardInput struct {
	sync.Mutex

	id    uuid.UUID
	Type  place.ThingType
	place *place.Place

	timereceived time.Time
	val          string
}

func NewKeyboardInput(input string) *KeyboardInput {
	return &KeyboardInput{
		id:    uuid.New(),
		Type:  place.KeyboardInput,
		place: nil,

		timereceived: time.Now(),
		val:          input,
	}
}

func (kbi *KeyboardInput) GetID() uuid.UUID {
	return kbi.id
}
func (kbi *KeyboardInput) GetType() place.ThingType {
	kbi.Lock()
	defer kbi.Unlock()

	return kbi.Type
}

func (kbi *KeyboardInput) GetPlace() *place.Place {
	kbi.Lock()
	defer kbi.Unlock()

	return kbi.place
}

// This function is not to be used to add or remove a thing from a place
func (kbi *KeyboardInput) SetPlace(newPlace *place.Place) {
	//TO DO: Confirm if it is not needed to lock before setting the place
	//kbi.Lock()
	//defer kbi.Unlock()

	kbi.place = newPlace
}

func (kbi *KeyboardInput) GetValue() string {
	kbi.Lock()
	defer kbi.Unlock()

	return kbi.val
}

func (kbi *KeyboardInput) GetTime() time.Time {
	kbi.Lock()
	defer kbi.Unlock()

	return kbi.timereceived
}
