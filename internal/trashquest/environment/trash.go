package environment

import (
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/landscape"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/landscape/landscape2d"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/thingstype"
	"github.com/google/uuid"
)

type Trash struct {
	id       uuid.UUID
	position *landscape2d.Position2D
}

func NewTrash() *Trash {
	return &Trash{
		id:       uuid.New(),
		position: nil,
	}
}

func (trash *Trash) GetID() uuid.UUID {
	return trash.id
}

func (trash *Trash) GetType() thingstype.ThingType {
	return thingstype.Trash
}

func (trash *Trash) GetPosition() landscape.Position {
	return trash.position
}

func (trash *Trash) SetPosition(position landscape.Position) {
	if trash.position != nil {
		trash.position = position.(*landscape2d.Position2D)
	} else {
		trash.position = nil
	}
}

func (trash *Trash) String() string {
	return "Trash"
}

func (trash *Trash) Ascii() string {
	return "+"
}
