package board

import (
	"errors"
)

var (
	ErrCellNotFound = errors.New("cell not found")
)

type Board struct {
	Cells [][]*Cell
}

func NewBoard(width, height int) *Board {

	// Initializing cell matrix
	matrix := make([][]*Cell, height)
	for i := range matrix {
		matrix[i] = make([]*Cell, width)
	}

	newBoard := Board{
		Cells: matrix,
	}

	// Alocate cells
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newBoard.Cells[i][j] = NewCell()
		}
	}

	newBoard.connectCells()

	return &newBoard
}

func (board *Board) connectCells() {
	for i := 0; i < len(board.Cells); i++ {
		for j := 0; j < len(board.Cells[i]); j++ {
			if i != 0 {
				board.Cells[i][j].North = board.Cells[i-1][j]
			}
			if i != len(board.Cells)-1 {
				board.Cells[i][j].South = board.Cells[i+1][j]
			}
			if j != 0 {
				board.Cells[i][j].West = board.Cells[i][j-1]
			}
			if j != len(board.Cells[i])-1 {
				board.Cells[i][j].East = board.Cells[i][j+1]
			}
		}
	}
}

func (board *Board) GetCell(x, y int) *Cell {
	return board.Cells[x][y]
}

func (board *Board) GetCellPosition(cell *Cell) (int, int, error) {
	for i := 0; i < len(board.Cells); i++ {
		for j := 0; j < len(board.Cells[i]); j++ {
			if board.Cells[i][j] == cell {
				return i, j, nil
			}
		}
	}

	return -1, -1, ErrCellNotFound
}

func (board *Board) GetNumberOfCells() int {
	return len(board.Cells) * len(board.Cells[0])
}

func (board *Board) GetNumberOfRows() int {
	return len(board.Cells)
}

func (board *Board) GetNumberOfColumns() int {
	return len(board.Cells[0])
}

// CleamItems removes all the items from the board
func (board *Board) CleamItems() {
	for i := 0; i < len(board.Cells); i++ {
		for j := 0; j < len(board.Cells[i]); j++ {
			board.Cells[i][j].CleamItems()
		}
	}
}
