package travelingsalesman

import (
	"strconv"
	"strings"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/circuit"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/population"
)

type TravelingSalesman struct {
	place.Thing
	alaive  bool
	fitness float64
	genome  Genome
	circuit circuit.Circuit
}

// Create a new traveling salesman with random values based in the number of the nodes.
func NewTravelingSalesman(c circuit.Circuit) *TravelingSalesman {

	if len(c.Nodes) < 2 {
		return nil
	}

	newGenome := NewRandomGenome(&c)

	newTravelingSalesman := TravelingSalesman{
		Thing:   place.NewThing(place.Life),
		alaive:  true,
		fitness: 0,
		genome:  *newGenome,
		circuit: c,
	}

	//newTravelingSalesman.Position = &c.Nodes[0]
	newTravelingSalesman.MoveToPlace(&c.Nodes[0])

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

// Run the genetic algorithm solution for the traveling salesman problem.
func (ts *TravelingSalesman) Run() error {

	positionsAlreadyVisited := make(map[*place.Place]bool)

	for ts.alaive {
		// Get the next connection to use.
		connectionNum, err := ts.genome.GetNextPosition(ts.getPositionNumber())
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.Run")
			ts.alaive = false
			return err
		}

		// Get the next connection.
		connection, err := ts.Position.GetConnectionNumber(connectionNum)
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.Run")
			ts.alaive = false
			return err
		}

		// Get next position and its distance.
		nextPosition := connection.GetDestination()
		distance := connection.GetWeight()

		val, ok := positionsAlreadyVisited[nextPosition]
		// If the key exists and val is true
		if ok && val {
			//log.Println("travelingsalesman.TravelingSalesman.Run")
			//ts.alaive = false
			return ErrInvalidMove
		}

		// Update the fitness.
		ts.fitness += distance

		// Update the actual position.
		ts.MoveToPlace(nextPosition)
		positionsAlreadyVisited[nextPosition] = true

		if ts.Position == &ts.circuit.Nodes[0] {
			// Verify if all the positions have been visited.
			if len(positionsAlreadyVisited) == len(ts.circuit.Nodes) {
				// This means that the objective have been achieved
				ts.fitness *= ts.fitness
				return nil
			} else {
				// In this case the traveling salesman is stuck in a loop.
				//log.Println("travelingsalesman.TravelingSalesman.Run")
				//ts.alaive = false
				return ErrInvalidMove
			}
		}
	}

	return nil
}

// Print the solution of the traveling salesman problem.
func (ts *TravelingSalesman) PrintSolution() string {
	var strBuilder strings.Builder

	// Initialize traveler
	ts.MoveToPlace(&ts.circuit.Nodes[0])
	ts.alaive = true
	ts.fitness = 0

	positionsAlreadyVisited := make(map[*place.Place]bool)

	strBuilder.WriteString("\n\r|" + strconv.Itoa(ts.getPositionNumber()) + "|")

	for {
		// Get the next connection to use.
		connectionNum, err := ts.genome.GetNextPosition(ts.getPositionNumber())
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.PrintSolution")
			strBuilder.Reset()
			strBuilder.WriteString("Error reading connection number: " + err.Error())
			return strBuilder.String()
		}

		// Get the next connection.
		connection, err := ts.Position.GetConnectionNumber(connectionNum)
		if err != nil {
			//log.Println("travelingsalesman.TravelingSalesman.PrintSolution")
			strBuilder.Reset()
			strBuilder.WriteString("Error getting the next connection: " + err.Error())
			return strBuilder.String()
		}

		// Get next position and its distance.
		nextPosition := connection.GetDestination()
		distance := connection.GetWeight()

		val, ok := positionsAlreadyVisited[nextPosition]
		// If the key exists and val is true
		if ok && val {
			//log.Println("travelingsalesman.TravelingSalesman.PrintSolution")
			strBuilder.WriteString(" -- " + strconv.FormatFloat(distance, 'f', 2, 64) + " --> |" + strconv.Itoa(int(ts.circuit.GetNodeNumber(nextPosition))) + "|")
			strBuilder.WriteString(" -- Invalid move -- Position already visited before.")
			return strBuilder.String()
		}

		strBuilder.WriteString(" -- " + strconv.FormatFloat(distance, 'f', 2, 64) + " --> |" + strconv.Itoa(int(ts.circuit.GetNodeNumber(nextPosition))) + "|")

		// Update the actual position.
		ts.MoveToPlace(nextPosition)
		positionsAlreadyVisited[nextPosition] = true

		if ts.Position == &ts.circuit.Nodes[0] {
			// Verify if all the positions have been visited.
			if len(positionsAlreadyVisited) == len(ts.circuit.Nodes) {
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

	newChildren := NewTravelingSalesman(ts.circuit)
	newChildren.genome = *newGenes

	return newChildren, nil
}

func (ts *TravelingSalesman) getPositionNumber() int {
	return ts.circuit.GetNodeNumber(ts.Position)
}
