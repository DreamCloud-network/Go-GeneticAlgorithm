package trashquest

import (
	"testing"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/board"
)

func TestGame(t *testing.T) {
	t.Log("Testing Game")

	newTrashBoard := NewTrashBoard(10)

	t.Log("Board columns: ", newTrashBoard.Board.GetNumberOfColumns())
	t.Log("Board rows: ", newTrashBoard.Board.GetNumberOfRows())
	t.Log("Board cells: ", newTrashBoard.Board.GetNumberOfCells())

	player := NewTrashPlayer(newTrashBoard)

	playerPosition, err := player.GetPosition()
	if err != nil {
		t.Log("Error getting player position")
		t.Error(err)
	}

	t.Log("Player position (line | column): ", playerPosition)

	err = player.Execute(StepNorth)
	if err != nil {
		t.Log("Expected error moving player to North. Got: ", err)
	}

	err = player.Execute(StepWest)
	if err != nil {
		t.Log("Expected error moving player to West. Got: ", err)
	}

	err = player.Execute(StepEast)
	if err != nil {
		t.Log("Error moving player to East.")
		t.Error(err)
	}

	playerPosition, err = player.GetPosition()
	if err != nil {
		t.Log("Error getting player position")
		t.Error(err)
	}

	t.Log("Player position (line | column): ", playerPosition)

	err = player.Execute(StepSouth)
	if err != nil {
		t.Log("Error moving player to South.")
		t.Error(err)
	}

	playerPosition, err = player.GetPosition()
	if err != nil {
		t.Log("Error getting player position")
		t.Error(err)
	}

	t.Log("Player position (line | column): ", playerPosition)
}

func TestTrashCollection(t *testing.T) {
	t.Log("Testing Trash Collection")

	newTrashBoard := NewTrashBoard(10)

	t.Log("Board columns: ", newTrashBoard.Board.GetNumberOfColumns())
	t.Log("Board rows: ", newTrashBoard.Board.GetNumberOfRows())
	t.Log("Board cells: ", newTrashBoard.Board.GetNumberOfCells())
	t.Log("Board trash: ", newTrashBoard.NumberOfTrash())

	player := NewTrashPlayer(newTrashBoard)

	// Collect all trash in the board
	movingEast := true
	for newTrashBoard.NumberOfTrash() > 0 {
		err := player.Execute(Pickup)
		if err != nil {
			if movingEast {
				err = player.Execute(StepEast)
				if err != nil {
					movingEast = false
					err = player.Execute(StepSouth)
					if err != nil {
						t.Log("Error moving player to South.")
						t.Error(err)
						return
					}
				}
			} else {
				err = player.Execute(StepWest)
				if err != nil {
					movingEast = true
					err = player.Execute(StepSouth)
					if err != nil {
						t.Log("Error moving player to South.")
						t.Error(err)
						return
					}
				}
			}
		}
	}

	playerPosition, err := player.GetPosition()
	if err != nil {
		t.Log("Error getting player position")
		t.Error(err)
	}

	t.Log("Player position (line | column): ", playerPosition)
	t.Log("Board trash: ", newTrashBoard.NumberOfTrash())
	t.Log("Trash Collected: ", player.NumberOfTrash())
}

func TestActions(t *testing.T) {
	t.Log("Testing Actions")

	player := NewTrashPlayer(NewTrashBoard(10))

	player.TrashQuestBoard.Board.GetCell(0, 0).AddItem(board.Trash)

	start := time.Now()
	err := player.Execute(StepNorth)
	if err != nil {
		t.Log("Expected error moving player to North. Got: ", err)
		t.Error(err)
	}
	elapsed := time.Since(start)
	t.Logf("Move north took %s", elapsed)

	start = time.Now()
	err = player.Execute(StepSouth)
	if err != nil {
		t.Log("Error moving player to South.")
		t.Error(err)
	}
	elapsed = time.Since(start)
	t.Logf("Move south took %s", elapsed)

	start = time.Now()
	err = player.Execute(StepNorth)
	if err != nil {
		t.Log("Expected error moving player to North. Got: ", err)
		t.Error(err)
	}
	elapsed = time.Since(start)
	t.Logf("Move north took %s", elapsed)

	start = time.Now()
	err = player.Execute(StepEast)
	if err != nil {
		t.Log("Error moving player to East.")
		t.Error(err)
	}
	elapsed = time.Since(start)
	t.Logf("Move east took %s", elapsed)

	start = time.Now()
	err = player.Execute(StepWest)
	if err != nil {
		t.Log("Expected error moving player to West. Got: ", err)
		t.Error(err)
	}
	elapsed = time.Since(start)
	t.Logf("Move west took %s", elapsed)

	start = time.Now()
	err = player.Execute(Pickup)
	if err != nil {
		t.Log("Error picking up trash.")
		t.Error(err)
	}
	elapsed = time.Since(start)
	t.Logf("Pickup took %s", elapsed)

	err = player.Execute(StepSouth)
	if err != nil {
		t.Log("Error moving player to South.")
		t.Error(err)
	}

	err = player.Execute(StepEast)
	if err != nil {
		t.Log("Error moving player to East.")
		t.Error(err)
	}

	start = time.Now()
	err = player.Execute(RandomMove)
	if err != nil {
		t.Log("Error doing random move.")
		t.Error(err)
	}

	elapsed = time.Since(start)
	t.Logf("Random move took %s", elapsed)

}
