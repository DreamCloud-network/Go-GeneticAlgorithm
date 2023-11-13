package things

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

type Earth struct {
	id    uuid.UUID
	Type  place.ThingType
	place *place.Place

	Nutrients int
}

func NewEarth(p int) *Earth {
	return &Earth{
		id:        uuid.New(),
		Type:      place.Earth,
		place:     nil,
		Nutrients: p,
	}
}

func (earth Earth) GetID() uuid.UUID {
	return earth.id
}
func (earth Earth) GetType() place.ThingType {
	return earth.Type
}

func (earth Earth) GetPlace() *place.Place {
	return earth.place
}

func (earth *Earth) SetPlace(newPlace *place.Place) {
	earth.place = newPlace
}
