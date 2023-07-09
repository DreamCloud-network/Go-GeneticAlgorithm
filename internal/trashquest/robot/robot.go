package robot

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/environment"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/genome"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/landscape"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/landscape/landscape2d"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/thingstype"
	"github.com/google/uuid"
)

var (
	reward  = 1
	penalty = 1
)

var (
	ErrActionNotAllowed    = errors.New("action not allowed")
	ErrNoTrash             = errors.New("no trash in position")
	ErrorInvalidAction     = errors.New("invalid action")
	ErrorActionsEmpty      = errors.New("actions list empty")
	ErrorBoardUndefined    = errors.New("board undefined")
	ErrorPositionUndefined = errors.New("position undefined")
	ErrorGenesUndefined    = errors.New("genes undefined")
)

type Robot struct {
	id       uuid.UUID
	position *landscape2d.Position2D

	TrashCollected []*environment.Trash
	Points         int
	Board          *environment.TrashQuestBoard
	Genes          genome.Genes
}

func (robot *Robot) GetID() uuid.UUID {
	return robot.id
}
func (robot *Robot) GetType() thingstype.ThingType {
	return thingstype.Robot
}
func (robot *Robot) GetPosition() landscape.Position {
	return robot.position
}

func (robot *Robot) SetPosition(position landscape.Position) {
	if position == nil {
		robot.position = nil
	} else {
		robot.position = position.(*landscape2d.Position2D)
	}
}

// Returns a new robot
func NewRobot() *Robot {
	newRobot := &Robot{
		id:             uuid.New(),
		position:       nil,
		TrashCollected: nil,
		Points:         0,
		Board:          nil,
		Genes:          nil,
	}

	return newRobot
}

func (robot *Robot) moveNorth() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.North)
	if err != nil {
		//log.Println("mitchelrobot.Robot.moveNorth - Error moving north.")
		robot.Points -= penalty
		return ErrActionNotAllowed
	}
	return nil
}

// MoveEast moves the player to the south
func (robot *Robot) moveSouth() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.South)
	if err != nil {
		//log.Println("mitchelrobot.Robot.moveSouth - Error moving south.")
		robot.Points -= penalty
		return ErrActionNotAllowed
	}
	return nil
}

// MoveEast moves the player to the east
func (robot *Robot) moveEast() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.East)
	if err != nil {
		//log.Println("mitchelrobot.Robot.moveEast - Error moving east.")
		robot.Points -= penalty
		return ErrActionNotAllowed
	}
	return nil
}

// MoveEast moves the player to the west
func (robot *Robot) moveWest() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.West)
	if err != nil {
		//log.Println("mitchelrobot.Robot.moveWest - Error moving west.")
		robot.Points -= penalty
		return ErrActionNotAllowed
	}
	return nil
}

// Pickup picks up the trash in the cell
func (robot *Robot) pickupTrash() error {
	trashesInPosition := robot.position.GetThingsByType(thingstype.Trash)
	if len(trashesInPosition) == 0 {
		return ErrNoTrash
	}

	trashPickedUp := robot.position.RemoveThing(trashesInPosition[0]).(*environment.Trash)

	robot.TrashCollected = append(robot.TrashCollected, trashPickedUp)

	robot.Points += reward

	return nil
}

// NumTrashCollected returns the number of trashes collected by the robot
func (robot *Robot) NumTrashCollected() int {
	if robot.TrashCollected == nil {
		return 0
	}
	return len(robot.TrashCollected)
}

func (robot *Robot) ExecuteAction(action genome.Action) error {
	if robot.Board == nil {
		return ErrorBoardUndefined
	}

	if robot.position == nil {
		return ErrorPositionUndefined
	}

	switch action {
	case genome.StepNorth:
		return robot.moveNorth()
	case genome.StepSouth:
		return robot.moveSouth()
	case genome.StepEast:
		return robot.moveEast()
	case genome.StepWest:
		return robot.moveWest()
	case genome.RandomMove:
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		ramdonMove := genome.Action(r1.Intn(int(genome.StepWest + 1)))
		return robot.ExecuteAction(ramdonMove)
	case genome.Pickup:
		return robot.pickupTrash()
	case genome.DoNothing:
		// do nothing
	default:
		return ErrorInvalidAction
	}

	return nil
}

func (robot *Robot) ExecuteActions(actions []genome.Action) error {
	if actions == nil {
		return ErrorActionsEmpty
	}

	for _, action := range actions {
		err := robot.ExecuteAction(action)
		if err != nil && err == ErrorInvalidAction {
			log.Println("robot.ExecuteActions: error executing actions.")
			return err
		}
	}

	return nil
}

func (robot *Robot) Evaluate() error {
	//start := time.Now()

	robot.Board = environment.NewTrashQuestBoard(10)

	//actions := robot.Genes.GetActions()

	for cont := 0; cont < 100; cont++ {
		robot.Board.CleanAllThings()
		robot.Board.PopulateBoardWithTrash()
		robot.Board.Positions[0][0].AddThing(robot)

		err := robot.ExecuteActions(robot.Genes.GetActions())
		if err != nil {
			log.Println("robot.Evaluate: error executing actions.")
			return err
		}
	}

	//evaluateTime := time.Since(start)
	//log.Printf("Evaluate took %s", evaluateTime)

	return nil
}

func (robot *Robot) Mate(partner *Robot) []*Robot {
	if robot.Genes == nil || partner.Genes == nil {
		return nil
	}
	newGenes := robot.Genes.Mate(partner.Genes)

	children := make([]*Robot, 0)

	for _, gene := range newGenes {
		newChild := NewRobot()
		newChild.Genes = gene

		children = append(children, newChild)
	}

	return children
}

func (robot *Robot) ReplayASCII() {
	log.Println("Robot: ", robot.id)
	log.Println("Reaplying...")

	robot.Points = 0
	robot.TrashCollected = nil

	robot.Board = environment.NewTrashQuestBoard(10)
	robot.Board.Positions[0][0].AddThing(robot)

	for cont, action := range robot.Genes.GetActions() {
		time.Sleep(1 * time.Second)

		robot.ExecuteAction(action)

		message := robot.Board.DrawsASCII()
		message += "Step: " + strconv.Itoa(cont)
		message += "\n\rPosition: " + robot.position.String()
		message += "\n\rLast Move: " + action.String()
		if cont < len(robot.Genes.GetActions())-1 {
			message += "\n\rNext Move: " + robot.Genes.GetActions()[cont+1].String()
		}
		message += "\n\rTrash collected: " + strconv.Itoa(robot.NumTrashCollected())
		message += "\n\rTotalPoints: " + strconv.Itoa(robot.Points)
		log.Println(message)
	}

	log.Println("Finished.")
}

func (robot *Robot) Clone() *Robot {
	newRobot := NewRobot()
	newRobot.Genes = robot.Genes.Duplicate()

	return newRobot
}
