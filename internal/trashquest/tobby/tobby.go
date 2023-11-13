package tobby

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/internal/trashquest/environment"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/environment/landscape"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/environment/landscape/landscape2d"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/environment/thingstype"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/population"
	"github.com/google/uuid"
)

const (
	trashReward  = 10
	wallPenalty  = 5
	trashPenalty = 1
	numMoves     = 200
	numRolls     = 100
)

type Tobby struct {
	id       uuid.UUID
	position *landscape2d.Position2D

	TrashCollected []*environment.Trash
	Fitness        float64
	Board          *environment.TrashQuestBoard
	Genes          TobbyGene
	LifeCicle      int
}

func (robot *Tobby) GetID() uuid.UUID {
	return robot.id
}
func (robot *Tobby) GetType() thingstype.ThingType {
	return thingstype.Robot
}
func (robot *Tobby) GetPosition() landscape.Position {
	return robot.position
}

func (robot *Tobby) SetPosition(position landscape.Position) {
	if position == nil {
		robot.position = nil
	} else {
		robot.position = position.(*landscape2d.Position2D)
	}
}

// Returns a new robot
func NewTobby() *Tobby {
	return &Tobby{
		id:             uuid.New(),
		position:       nil,
		TrashCollected: nil,
		Fitness:        0,
		Board:          nil,
		Genes:          *NewRandomGenes(),
		LifeCicle:      0,
	}
}

func NewRandomTobby() *Tobby {
	newTobby := NewTobby()

	return newTobby
}

func (robot *Tobby) moveNorth() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.North)
	if err != nil {
		//log.Println("Tobby.moveNorth - Error moving north.")
		robot.Fitness -= wallPenalty
		return ErrActionNotAllowed
	}
	return nil
}

// MoveEast moves the player to the south
func (robot *Tobby) moveSouth() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.South)
	if err != nil {
		//log.Println("Tobby.moveSouth - Error moving south.")
		robot.Fitness -= wallPenalty
		return ErrActionNotAllowed
	}
	return nil
}

// MoveEast moves the player to the east
func (robot *Tobby) moveEast() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.East)
	if err != nil {
		//log.Println("Tobby.moveEast - Error moving east.")
		robot.Fitness -= wallPenalty
		return ErrActionNotAllowed
	}
	return nil
}

// MoveEast moves the player to the west
func (robot *Tobby) moveWest() error {
	err := robot.position.MoveThingToDirection(robot, landscape2d.West)
	if err != nil {
		//log.Println("Tobby.moveWest - Error moving west.")
		robot.Fitness -= wallPenalty
		return ErrActionNotAllowed
	}
	return nil
}

// Pickup picks up the trash in the cell
func (robot *Tobby) pickupTrash() error {
	trashesInPosition := robot.position.GetThingsByType(thingstype.Trash)
	if len(trashesInPosition) == 0 {
		robot.Fitness -= trashPenalty
		return ErrNoTrash
	}

	trashPickedUp := robot.position.RemoveThing(trashesInPosition[0]).(*environment.Trash)

	robot.TrashCollected = append(robot.TrashCollected, trashPickedUp)

	robot.Fitness += trashReward

	return nil
}

// NumTrashCollected returns the number of trashes collected by the robot
func (robot *Tobby) NumTrashCollected() int {
	if robot.TrashCollected == nil {
		return 0
	}
	return len(robot.TrashCollected)
}

func (robot *Tobby) ExecuteAction(action Action) error {
	if robot.Board == nil {
		return ErrorBoardUndefined
	}

	if robot.position == nil {
		return ErrorPositionUndefined
	}

	switch action {
	case StepNorth:
		robot.LifeCicle++
		return robot.moveNorth()
	case StepSouth:
		robot.LifeCicle++
		return robot.moveSouth()
	case StepEast:
		robot.LifeCicle++
		return robot.moveEast()
	case StepWest:
		robot.LifeCicle++
		return robot.moveWest()
	case RandomMove:
		return robot.ExecuteAction(GetRandomMove())
	case Pickup:
		robot.LifeCicle++
		return robot.pickupTrash()
	default:
		return ErrorInvalidAction
	}
}

func (robot *Tobby) getPositionSignature() int {
	// Sequence center -> East -> North -> West -> South
	// 0 -> no trash in position
	// 1 -> trash in position
	// 2 -> wall in position

	posSignature := 0.0

	if robot.position.HasThingType(thingstype.Trash) {
		posSignature += 1 * math.Pow(3, 0)
	} else {
		posSignature += 0 * math.Pow(3, 0)
	}

	eastPosition := robot.position.GetRelativePosition(landscape2d.East, 1)
	if eastPosition != nil {
		if eastPosition.HasThingType(thingstype.Trash) {
			posSignature += 1 * math.Pow(3, 1)
		} else {
			posSignature += 0 * math.Pow(3, 1)
		}
	} else {
		posSignature += 2 * math.Pow(3, 1)
	}

	northPosition := robot.position.GetRelativePosition(landscape2d.North, 1)
	if northPosition != nil {
		if northPosition.HasThingType(thingstype.Trash) {
			posSignature += 1 * math.Pow(3, 2)
		} else {
			posSignature += 0 * math.Pow(3, 2)
		}
	} else {
		posSignature += 2 * math.Pow(3, 2)
	}

	westPosition := robot.position.GetRelativePosition(landscape2d.West, 1)
	if westPosition != nil {
		if westPosition.HasThingType(thingstype.Trash) {
			posSignature += 1 * math.Pow(3, 3)
		} else {
			posSignature += 0 * math.Pow(3, 3)
		}
	} else {
		posSignature += 2 * math.Pow(3, 3)
	}

	southPosition := robot.position.GetRelativePosition(landscape2d.South, 1)
	if southPosition != nil {
		if southPosition.HasThingType(thingstype.Trash) {
			posSignature += 1 * math.Pow(3, 4)
		} else {
			posSignature += 0 * math.Pow(3, 4)
		}
	} else {
		posSignature += 2 * math.Pow(3, 4)
	}

	return int(posSignature)
}

func (robot *Tobby) ReplayASCII() error {
	log.Println("Robot: ", robot.id)
	log.Println("Reaplying...")

	robot.Fitness = 0
	robot.TrashCollected = nil

	robot.Board = environment.NewTrashQuestBoard(10)
	robot.Board.Positions[0][0].AddThing(robot)

	positionSignature := robot.getPositionSignature()
	action := robot.Genes.GetAction(positionSignature)

	message := robot.Board.DrawsASCII()
	message += "Step: 0"
	message += "\n\rPosition: " + robot.position.String() + " - Signature: " + strconv.Itoa(robot.getPositionSignature())
	message += "\n\rLast Move: ---"
	message += "\n\rNext Move: " + action.String()
	message += "\n\rTrash collected: " + strconv.Itoa(robot.NumTrashCollected())
	message += "\n\rTotalPoints: " + strconv.FormatFloat(robot.Fitness, 'f', 2, 64)
	log.Println(message)

	for cont := 0; cont < numMoves; cont++ {
		//time.Sleep(1 * time.Second)
		fmt.Scanln()

		positionSignature = robot.getPositionSignature()
		action := robot.Genes.GetAction(positionSignature)

		err := robot.ExecuteAction(action)
		if err != nil && err == ErrorInvalidAction {
			log.Println("robot.Run: error executing action.")
			return err
		}

		message := robot.Board.DrawsASCII()
		message += "Step: " + strconv.Itoa(cont)
		message += "\n\rPosition: " + robot.position.String() + " - Signature: " + strconv.Itoa(robot.getPositionSignature())
		message += "\n\rLast Move: " + action.String()
		message += "\n\rNext Move: " + robot.Genes.GetAction(robot.getPositionSignature()).String()
		message += "\n\rTrash collected: " + strconv.Itoa(robot.NumTrashCollected())
		message += "\n\rTotalPoints: " + strconv.FormatFloat(robot.Fitness, 'f', 2, 64)
		log.Println(message)
	}

	log.Println("Finished.")

	return nil
}

func (tobby *Tobby) Run() error {

	for cont := 0; cont < numRolls; cont++ {
		tobby.Board = environment.NewTrashQuestBoard(10)
		tobby.Board.Positions[0][0].AddThing(tobby)
		tobby.LifeCicle = 0

		for tobby.LifeCicle < numMoves {

			err := tobby.ExecuteAction(tobby.Genes.GetAction(tobby.getPositionSignature()))
			if err != nil && err == ErrorInvalidAction {
				log.Println("robot.Run: error executing action.")
				return err
			}
		}
	}

	tobby.Fitness = tobby.Fitness / numRolls

	return nil
}

// Individual interface implementation

func (robot *Tobby) GetFitness() float64 {
	return robot.Fitness
}

func (robot *Tobby) Mate(individual population.Individual) (population.Individual, error) {
	partner := individual.(*Tobby)
	newGenes := robot.Genes.Mate(&partner.Genes)
	if newGenes == nil {
		return nil, ErrorMatingWithPartner
	}

	newChildren := NewTobby()
	newChildren.Genes = *newGenes

	return newChildren, nil
}
