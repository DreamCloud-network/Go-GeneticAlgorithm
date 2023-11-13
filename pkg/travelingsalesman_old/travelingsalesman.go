package travelingsalesman

import (
	"log"
	"strconv"
	"strings"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/population"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/travelingsalesman/circuit"
)

type TravelingSalesman struct {
	alaive          bool
	fitness         float64
	genome          Genome
	circuit         circuit.Circuit
	initialPosition int
	finalPosition   int
}

// Create a new traveling salesman with random values based in the number of the nodes.
func NewTravelingSalesman(c circuit.Circuit, initialPos, finalPos int) *TravelingSalesman {

	if initialPos < 0 || initialPos >= len(c) {
		log.Println("travelingsalesman.NewTravelingSalesman")
		return nil
	}

	if finalPos < 0 || finalPos >= len(c) {
		log.Println("travelingsalesman.NewTravelingSalesman")
		return nil
	}

	newGenome := NewRandomGenome(len(c))

	newTravelingSalesman := TravelingSalesman{
		alaive:          true,
		fitness:         0,
		genome:          *newGenome,
		circuit:         c,
		initialPosition: initialPos,
		finalPosition:   finalPos,
	}

	return &newTravelingSalesman
}

// Return the fitness of the traveling salesman.
func (ts *TravelingSalesman) GetFitness() float64 {
	return ts.fitness
}

// Return true if the traveling salesman is alive.
func (ts *TravelingSalesman) IsAlive() bool {
	return ts.alaive
}

func (ts *TravelingSalesman) Benefit() {
	ts.fitness += ts.fitness
}

// Run the genetic algorithm solution for the traveling salesman problem.
func (ts *TravelingSalesman) Run() error {
	actualPos := ts.initialPosition

	positionsAlreadyVisited := make(map[int]bool)

	for ts.alaive {
		// Get the next position.
		nextPos, err := ts.genome.GetNextPosition(actualPos)
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.Run")
			ts.alaive = false
			return err
		}

		// Get the distance between the actual position and the next position.
		distance, err := ts.circuit.GetConnectionWeight(actualPos, nextPos)
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.Run")
			ts.alaive = false
			return err
		}

		// Verify if there is aconnection between the actual position and the next position.
		if distance == 0 {
			//log.Println("travelingsalesman.TravelingSales1man.Run")
			ts.alaive = false
			return ErrInvalidMove
		}

		val, ok := positionsAlreadyVisited[nextPos]
		// If the key exists and val is true
		if ok && val {
			//log.Println("travelingsalesman.TravelingSalesman.Run")
			ts.alaive = false
			return ErrInvalidMove
		}

		// Update the fitness.
		ts.fitness += distance

		// Update the actual position.
		actualPos = nextPos
		positionsAlreadyVisited[nextPos] = true

		if actualPos == ts.finalPosition {
			// Verify if all the positions have been visited.
			if len(positionsAlreadyVisited) == len(ts.circuit) {
				// This means that the objective have been achieved
				return nil
			} else {
				// In this case the traveling salesman is stuck in a loop.
				//log.Println("travelingsalesman.TravelingSalesman.Run")
				ts.alaive = false
				return ErrInvalidMove
			}
		}
	}

	return nil
}

// Print the solution of the traveling salesman problem.
func (ts *TravelingSalesman) PrintSolution() string {
	var strBuilder strings.Builder

	positionsAlreadyVisited := make(map[int]bool)

	actualPos := ts.initialPosition

	strBuilder.WriteString("\n\r|" + strconv.Itoa(actualPos) + "|")

	for {
		// Get the next position.
		nextPos, err := ts.genome.GetNextPosition(actualPos)
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.PrintSolution")
			strBuilder.Reset()
			strBuilder.WriteString("Error reading nextPosition: " + err.Error())
			return strBuilder.String()
		}

		// Get the distance between the actual position and the next position.
		distance, err := ts.circuit.GetConnectionWeight(actualPos, nextPos)
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.PrintSolution")
			strBuilder.Reset()
			strBuilder.WriteString("Error reading distance: " + err.Error())
			return strBuilder.String()
		}

		// Verify if there is aconnection between the actual position and the next position.
		if distance == 0 {
			//log.Println("travelingsalesman.TravelingSales1man.PrintSolution")
			strBuilder.WriteString(" -- " + strconv.FormatFloat(distance, 'f', 2, 64) + " --> |" + strconv.Itoa(int(nextPos)) + "|")
			strBuilder.WriteString(" -- Invalid move -- No connection.")
			return strBuilder.String()
		}

		val, ok := positionsAlreadyVisited[nextPos]
		// If the key exists and val is true
		if ok && val {
			//log.Println("travelingsalesman.TravelingSalesman.PrintSolution")
			strBuilder.WriteString(" -- " + strconv.FormatFloat(distance, 'f', 2, 64) + " --> |" + strconv.Itoa(int(nextPos)) + "|")
			strBuilder.WriteString(" -- Invalid move -- Position already visited before.")
			return strBuilder.String()
		}

		strBuilder.WriteString(" -- " + strconv.FormatFloat(distance, 'f', 2, 64) + " --> |" + strconv.Itoa(int(nextPos)) + "|")

		// Update the actual position.
		actualPos = nextPos
		positionsAlreadyVisited[nextPos] = true

		if actualPos == ts.finalPosition {
			// Verify if all the positions have been visited.
			if len(positionsAlreadyVisited) == len(ts.circuit) {
				// This means that the objective have been achieved
				strBuilder.WriteString(" <> Objective reached!")
				return strBuilder.String()
			} else {
				// In this case the traveling salesman is stuck in a loop.
				//log.Println("travelingsalesman.TravelingSalesman.Run")
				strBuilder.WriteString(" -- Did not passed in all other positions.")
				return strBuilder.String()
			}
		}
	}
}

// Mate two traveling salesman.
func (ts *TravelingSalesman) Mate(individual population.Individual) (population.Individual, error) {
	tsPartner := individual.(*TravelingSalesman)

	newGenes, err := ts.genome.Mate(&tsPartner.genome)
	if newGenes == nil {
		return nil, err
	}

	newChildren := NewTravelingSalesman(ts.circuit, ts.initialPosition, ts.finalPosition)
	newChildren.genome = *newGenes

	return newChildren, nil
}
