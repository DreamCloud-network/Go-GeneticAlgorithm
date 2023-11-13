package environment

import (
	"log"
	"testing"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/environment/thingstype"
)

func TestBoard(t *testing.T) {
	t.Log("Testing Board")

	board := NewTrashQuestBoard(10)

	log.Println("Trash in the board: ", len(board.GetAllThingByType(thingstype.Trash)))
}

func TestBoardPrint(t *testing.T) {
	t.Log("Testing Board Print")

	board := NewTrashQuestBoard(10)

	log.Println(board.DrawsASCII())
}

/*
func TestNillvsInt(t *testing.T) {
	t.Log("Testing Nill vs Int")

	board := NewTrashQuestBoard(10)

	distance := 1
	direction := landscape2d.South

	start := time.Now()
	relativePosition := board.Positions[0][0].GetRelativePosition(direction, distance)
	if relativePosition == nil {
		t.Log("relativePosition is nil")
	}
	elapsed := time.Since(start)
	log.Println("Pointer test took:", elapsed)

	start = time.Now()
	relativePosition = board.GetRelativeTestPosition(direction, distance)
	if relativePosition == nil {
		t.Log("relativePosition is nil")
	}
	elapsed = time.Since(start)
	log.Println("Integer test took:", elapsed)

}

// GetNextPosition return the position in the direction and distance specified
func (board *TrashQuestBoard) GetRelativeTestPosition(direction landscape2d.Directions, distance int) *landscape2d.Position2D {
	if distance == 0 {
		return board.Positions[0][0]
	}

	relativePosition := board.Positions[0][0]

	for cont := 0; cont < distance; cont++ {
		actualPosition := relativePosition

		switch direction {
		case landscape2d.North:
			if actualPosition.Y > 0 {
				relativePosition = board.Positions[actualPosition.X][actualPosition.Y-1]
			}
		case landscape2d.South:
			if actualPosition.Y < len(board.Positions[0])-1 {
				relativePosition = board.Positions[actualPosition.X][actualPosition.Y+1]
			}
		case landscape2d.East:
			if actualPosition.X < len(board.Positions)-1 {
				relativePosition = board.Positions[actualPosition.X+1][actualPosition.Y]
			}
		case landscape2d.West:
			relativePosition = board.Positions[actualPosition.X-1][actualPosition.Y]
		case landscape2d.NorthEast:
			relativePosition = board.Positions[actualPosition.X+1][actualPosition.Y-1]
		case landscape2d.NorthWest:
			relativePosition = board.Positions[actualPosition.X-1][actualPosition.Y-1]
		case landscape2d.SouthEast:
			relativePosition = board.Positions[actualPosition.X+1][actualPosition.Y+1]
		case landscape2d.SouthWest:
			relativePosition = board.Positions[actualPosition.X-1][actualPosition.Y+1]

		}
	}

	return relativePosition
}
*/
