package things

import (
	"sync"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type Nutrient struct {
	sync.Mutex

	id    uuid.UUID
	Type  place.ThingType
	place *place.Place
}

func NewNutrient() *Nutrient {
	return &Nutrient{
		id:    uuid.New(),
		Type:  place.Nutrient,
		place: nil,
	}
}

func (nutrient *Nutrient) GetID() uuid.UUID {
	nutrient.Lock()
	defer nutrient.Unlock()

	return nutrient.id
}
func (nutrient *Nutrient) GetType() place.ThingType {
	nutrient.Lock()
	defer nutrient.Unlock()

	return nutrient.Type
}

func (nutrient *Nutrient) GetPlace() *place.Place {
	nutrient.Lock()
	defer nutrient.Unlock()

	return nutrient.place
}

// This function is not to be used to add or remove a thing from a place
func (nutrient *Nutrient) SetPlace(newPlace *place.Place) {
	nutrient.place = newPlace
}
