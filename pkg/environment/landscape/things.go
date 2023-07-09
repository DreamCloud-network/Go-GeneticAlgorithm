package landscape

import (
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/thingstype"
	"github.com/google/uuid"
)

type Thing interface {
	GetID() uuid.UUID

	GetType() thingstype.ThingType

	GetPosition() Position
	SetPosition(position Position)
}
