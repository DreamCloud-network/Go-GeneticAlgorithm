package environment

import (
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/landscape/landscape2d"
)

const trashProbability = 50 // %

type TrashQuestBoard struct {
	landscape2d.Landscape2D
}

// Creates a ne trash board populated with trashes.
func NewTrashQuestBoard(sideSize int) *TrashQuestBoard {

	var newBoard TrashQuestBoard

	newBoard.Landscape2D = landscape2d.New(sideSize, sideSize)

	newBoard.PopulateBoardWithTrash()

	return &newBoard
}

// Populate the board with trash in random positions with prob% chance of having a trash in each position.
func (board *TrashQuestBoard) PopulateBoardWithTrash() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for _, positions := range board.Landscape2D.Positions {
		for _, position := range positions {
			if r1.Intn(100) < trashProbability {
				newTrash := NewTrash()
				position.AddThing(newTrash)
			}
		}
	}
}
