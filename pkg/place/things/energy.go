package things

import (
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type Energy struct {
	id    uuid.UUID
	Type  place.ThingType
	place *place.Place
}

func NewEnergy() *Energy {
	return &Energy{
		id:    uuid.New(),
		Type:  place.Energy,
		place: nil,
	}
}

func (energy Energy) GetID() uuid.UUID {
	return energy.id
}
func (energy Energy) GetType() place.ThingType {
	return energy.Type
}

func (energy Energy) GetPlace() *place.Place {
	return energy.place
}

func (energy *Energy) SetPlace(newPlace *place.Place) {
	energy.place = newPlace
}
