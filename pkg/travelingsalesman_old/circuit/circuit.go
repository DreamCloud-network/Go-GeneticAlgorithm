package circuit

import (
	"log"
	"math/rand"
)

// Circuit is a 2D slice of float64s
// The first dimensions represents the nodes,
// the second dimension represent the connections with the weights
// 0 means no connection
type Circuit [][]float64

// Create a new cricuit with the given number of nodes and no connections
func NewCircuit(nodeNumber int) Circuit {
	circuit := make(Circuit, nodeNumber)
	for i := range circuit {
		circuit[i] = make([]float64, nodeNumber)
	}
	return circuit
}

func NewRandomCircuit(nodeNumber int) Circuit {
	circuit := NewCircuit(nodeNumber)
	for i := range circuit {
		for j := range circuit[i] {
			if i != j {
				circuit[i][j] = float64(rand.Intn(100))
			}
		}
	}
	return circuit
}

// Verify if the parameters nodes are ok to be processed.
// They must be grater than zero, less than the number of nodes and different.
func (circuit Circuit) verifyNodes(nodeA, nodeB int) error {
	if nodeA < 0 || nodeA >= len(circuit) {
		log.Println("circuit.Circuit.verifyNodes")
		return ErrOutOfRange
	}

	if nodeB < 0 || nodeB >= len(circuit) {
		log.Println("circuit.Circuit.verifyNodes")
		return ErrOutOfRange
	}

	if nodeA == nodeB {
		log.Println("circuit.Circuit.verifyNodes")
		return ErrSameNode
	}
	return nil
}

// SymmetricalConnectNodes connects two nodes with the given weight in both directions.
func (circuit Circuit) SymmetricalConnectNodes(nodeA, nodeB int, weight float32) error {

	err := circuit.verifyNodes(nodeA, nodeB)
	if err != nil {
		log.Println("circuit.Circuit.SymmetricalConnectNodes")
		return err
	}

	circuit[nodeA][nodeB] = float64(weight)
	circuit[nodeB][nodeA] = float64(weight)
	return nil
}

// AsymmetricalConnectNodes connects two nodes with the given weight in only A to B direction.
func (circuit Circuit) AsymmetricalConnectNodes(nodeA, nodeB int, weight float32) error {

	err := circuit.verifyNodes(nodeA, nodeB)
	if err != nil {
		log.Println("circuit.Circuit.AsymmetricalConnectNodes")
		return err
	}

	circuit[nodeA][nodeB] = float64(weight)
	return nil
}

// GetConnectionWeight returns the weight of the connection between the two nodes in the A to B direction.
func (circuit Circuit) GetConnectionWeight(nodeA, nodeB int) (float64, error) {
	if nodeA < 0 || nodeA >= len(circuit) {
		log.Println("circuit.Circuit.verifyNodes")
		return -1, ErrOutOfRange
	}

	if nodeB < 0 || nodeB >= len(circuit) {
		log.Println("circuit.Circuit.verifyNodes")
		return -1, ErrOutOfRange
	}

	return circuit[nodeA][nodeB], nil
}

// This function generate a test solution.
// For any size number the function generates a circuit that
// the nodes are connect in sequence with the higher weight 100,
// and populates the rest with ramdom values lesser than 100.
func TestSolution(circuitSize int) Circuit {

	newCricuit := NewCircuit(circuitSize)

	for nodeA := 0; nodeA < len(newCricuit); nodeA++ {
		for nodeB := 0; nodeB < len(newCricuit[nodeA]); nodeB++ {
			if nodeA == nodeB {
				newCricuit[nodeA][nodeB] = 0
			} else if nodeA == nodeB-1 {
				newCricuit[nodeA][nodeB] = 100
			} else {
				newCricuit[nodeA][nodeB] = float64(rand.Intn(100))
			}
		}
	}

	nodeA := 0
	nodeB := len(newCricuit) - 1
	newCricuit[nodeB][nodeA] = 100

	return newCricuit
}
