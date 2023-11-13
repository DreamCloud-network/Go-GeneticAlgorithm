package things

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type Light struct {
	id    uuid.UUID
	Type  place.ThingType
	place *place.Place

	Power int
}

func NewLight(p int) *Light {
	return &Light{
		id:    uuid.New(),
		Type:  place.Light,
		place: nil,
		Power: p,
	}
}

func (light Light) GetID() uuid.UUID {
	return light.id
}
func (light Light) GetType() place.ThingType {
	return light.Type
}

func (light Light) GetPlace() *place.Place {
	return light.place
}

func (light *Light) SetPlace(newPlace *place.Place) {
	light.place = newPlace
}
