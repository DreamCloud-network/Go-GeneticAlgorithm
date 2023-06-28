package trashquest

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/board"
)

var (
	ErrActionNotAllowed = errors.New("action not allowed")
	ErrNoTrash          = errors.New("no trash in cell")
)

var (
	moveReward   = 0
	pickupReward = 1
	faultPenalty = -1
)

type TrashPlayer struct {
	TrashQuestBoard *TrashBoard
	Position        *board.Cell
	CollectedItems  []board.Item
	Points          int
}

func (player *TrashPlayer) AddItem(item board.Item) {
	player.CollectedItems = append(player.CollectedItems, item)
}

// Returns a new player
func NewTrashPlayer(tBoard *TrashBoard) *TrashPlayer {
	var newTrashQuestPlayer TrashPlayer

	newTrashQuestPlayer.TrashQuestBoard = tBoard
	newTrashQuestPlayer.Position = tBoard.Board.Cells[0][0]
	newTrashQuestPlayer.Points = 0

	return &newTrashQuestPlayer
}

// MoveSequence is a function that executes a sequence of actions
func (player *TrashPlayer) MoveSequence(action []Action, printMovingSequence bool) error {
	for _, action := range action {

		movingString := ""

		if printMovingSequence {
			position, err := player.GetPosition()
			if err != nil {
				log.Println("trashquest.TrashQuestPlayer.MoveSequence: error getting player position")
				return err
			}

			actionsSrting, err := action.String()
			if err != nil {
				log.Println("trashquest.TrashQuestPlayer.MoveSequence: error getting action string")
				return err
			}
			movingString = position + " -> " + actionsSrting + " -> "
		}

		err := player.Execute(action)
		if printMovingSequence {
			position, err2 := player.GetPosition()
			if err2 != nil {
				log.Println("trashquest.TrashQuestPlayer.MoveSequence: error getting player position")
				return err2
			}
			movingString += position

			if err != nil {
				movingString += " - fault"
			} else if action == Pickup {
				movingString += " - trash picked up"
			}
		}

		if printMovingSequence {
			fmt.Println(movingString, " - reward: ", player.Points)
		}
	}

	return nil
}

func (player *TrashPlayer) Execute(action Action) error {
	switch action {
	case StepNorth:
		err := player.MoveNorth()
		if err != nil {
			player.Points += faultPenalty
			return err
		}
		player.Points += moveReward
		return nil
	case StepSouth:
		err := player.MoveSouth()
		if err != nil {
			player.Points += faultPenalty
			return err
		}
		player.Points += moveReward
		return nil
	case StepEast:
		err := player.MoveEast()
		if err != nil {
			player.Points += faultPenalty
			return err
		}
		player.Points += moveReward
		return nil
	case StepWest:
		err := player.MoveWest()
		if err != nil {
			player.Points += faultPenalty
			return err
		}
		player.Points += moveReward
		return nil
	case RandomMove:
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		ramdonMove := Action(r1.Intn(int(StepWest + 1)))
		return player.Execute(ramdonMove)
	case Pickup:
		err := player.PickupTrash()
		if err != nil {
			return err
		}
		player.Points += pickupReward
		return nil
	case DoNothing:
		// do nothing
	default:
		player.Points += faultPenalty
		return ErrorInvalidAction
	}

	return nil
}

func (player *TrashPlayer) GetPosition() (string, error) {
	line, column, err := player.TrashQuestBoard.Board.GetCellPosition(player.Position)
	if err != nil {
		log.Println("trashquest.TrashQuestPlayer.GetPosition: error getting cell position")
		return "", err
	}

	return "(" + strconv.Itoa(line) + "," + strconv.Itoa(column) + ")", nil
}

// NumberOfTrash returns the number of trash collected by the player
func (player *TrashPlayer) NumberOfTrash() int {

	numberOfTrash := len(player.CollectedItems)

	//log.Println("trashquest.TrashQuestBoard.NumberOfTrahs - Number of trash: " + strconv.Itoa(numberOfTrash))

	return numberOfTrash
}

// MoveEast moves the player to the north
func (player *TrashPlayer) MoveNorth() error {
	if player.Position.North == nil {
		return ErrActionNotAllowed
	}
	player.Position = player.Position.North
	return nil
}

// MoveEast moves the player to the south
func (player *TrashPlayer) MoveSouth() error {
	if player.Position.South == nil {
		return ErrActionNotAllowed
	}
	player.Position = player.Position.South
	return nil
}

// MoveEast moves the player to the east
func (player *TrashPlayer) MoveEast() error {
	if player.Position.East == nil {
		return ErrActionNotAllowed
	}
	player.Position = player.Position.East
	return nil
}

// MoveEast moves the player to the west
func (player *TrashPlayer) MoveWest() error {
	if player.Position.West == nil {
		return ErrActionNotAllowed
	}
	player.Position = player.Position.West
	return nil
}

// Pickup picks up the trash in the cell
func (player *TrashPlayer) PickupTrash() error {
	if !player.Position.HasItem(board.Trash) {
		return ErrNoTrash
	}
	player.Position.RemoveItem(board.Trash)
	player.AddItem(board.Trash)
	return nil
}
