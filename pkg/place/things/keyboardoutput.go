package things

import (
	"sync"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type KeyboardOutput struct {
	sync.Mutex

	id    uuid.UUID
	Type  place.ThingType
	place *place.Place

	timereceived time.Time
	val          string
}

func NewKeyboardOutput(input string) *KeyboardOutput {
	return &KeyboardOutput{
		id:    uuid.New(),
		Type:  place.KeyboardOutput,
		place: nil,

		timereceived: time.Now(),
		val:          input,
	}
}

func (kbo *KeyboardOutput) GetID() uuid.UUID {
	return kbo.id
}
func (kbo *KeyboardOutput) GetType() place.ThingType {
	kbo.Lock()
	defer kbo.Unlock()

	return kbo.Type
}

func (kbo *KeyboardOutput) GetPlace() *place.Place {
	kbo.Lock()
	defer kbo.Unlock()

	return kbo.place
}

// This function is not to be used to add or remove a thing from a place
func (kbo *KeyboardOutput) SetPlace(newPlace *place.Place) {
	//TO DO: Confirm if it is not needed to lock before setting the place
	//kbo.Lock()
	//defer kbo.Unlock()

	kbo.place = newPlace
}

func (kbo *KeyboardOutput) GetValue() string {
	kbo.Lock()
	defer kbo.Unlock()

	return kbo.val
}

func (kbo *KeyboardOutput) GetTime() time.Time {
	kbo.Lock()
	defer kbo.Unlock()

	return kbo.timereceived
}
