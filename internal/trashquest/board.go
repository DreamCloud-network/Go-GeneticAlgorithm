package trashquest

import (
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/board"
)

type TrashBoard struct {
	Board *board.Board
}

func NewTrashBoard(size int) *TrashBoard {
	var newTrashQuestBoard TrashBoard

	newTrashQuestBoard.Board = board.NewBoard(size, size)

	newTrashQuestBoard.PopulateBoardWithTrash()

	return &newTrashQuestBoard
}

func (tBoard *TrashBoard) PopulateBoardWithTrash() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < len(tBoard.Board.Cells); i++ {
		for j := 0; j < len(tBoard.Board.Cells[i]); j++ {
			if r1.Intn(100) > 50 {
				tBoard.Board.Cells[i][j].Items = append(tBoard.Board.Cells[i][j].Items, board.Trash)
			}
		}
	}
}

// NumberOfTrash returns the number of trash in the board
func (tBoard *TrashBoard) NumberOfTrash() int {
	var numberOfTrash int

	for i := 0; i < len(tBoard.Board.Cells); i++ {
		for j := 0; j < len(tBoard.Board.Cells[i]); j++ {
			if tBoard.Board.Cells[i][j].HasItem(board.Trash) {
				numberOfTrash++
			}
		}
	}

	//log.Println("trashquest.TrashQuestBoard.NumberOfTrahs - Number of trash: " + strconv.Itoa(numberOfTrash))

	return numberOfTrash
}
