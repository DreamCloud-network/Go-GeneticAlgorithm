package circuit

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place"
)

// Circuit is a 2D slice of float64s
// The first dimensions represents the nodes,
// the second dimension represent the connections with the weights
// 0 means no connection
type Circuit struct {
	Nodes []place.Place
}

// Create a new cricuit with the given number of nodes and no connections
func NewCircuit(nodeNumber int) Circuit {
	newCircuit := Circuit{
		Nodes: make([]place.Place, nodeNumber),
	}

	for i := range newCircuit.Nodes {
		newCircuit.Nodes[i] = place.NewPlace()
	}

	return newCircuit
}

// This function generate a test solution.
// For any size number the function generates a circuit that
// the nodes are connect in sequence with the higher weight 100,
// and populates the rest with ramdom values lesser than 100.
/*func NewTestCircuit() Circuit {
	newRandomCircuit := NewCircuit(4)

	newRandomCircuit.Nodes[0].AddConnection(&newRandomCircuit.Nodes[1], 10)

	newRandomCircuit.Nodes[1].AddConnection(&newRandomCircuit.Nodes[2], 5)
	newRandomCircuit.Nodes[1].AddConnection(&newRandomCircuit.Nodes[3], 10)

	newRandomCircuit.Nodes[2].AddConnection(&newRandomCircuit.Nodes[0], 10)

	newRandomCircuit.Nodes[3].AddConnection(&newRandomCircuit.Nodes[2], 10)
	newRandomCircuit.Nodes[3].AddConnection(&newRandomCircuit.Nodes[0], 5)

	return newRandomCircuit
}*/

func NewTestCircuit(size int) Circuit {
	newRandomCircuit := NewCircuit(size)

	// Generate simple solution
	for i := 0; i < len(newRandomCircuit.Nodes)-1; i++ {
		newRandomCircuit.Nodes[i].AddConnection(&newRandomCircuit.Nodes[i+1], 5)
	}
	newRandomCircuit.Nodes[size-1].AddConnection(&newRandomCircuit.Nodes[0], 5)

	// Generate random connections
	for i := 0; i < len(newRandomCircuit.Nodes); i++ {
		for j := 0; j < len(newRandomCircuit.Nodes); j++ {
			if (i != j) && (i != j-1) && !((i == size-1) && (j == 0)) {
				r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
				//if r1.Intn(10) > 5 {
				newRandomCircuit.Nodes[i].AddConnection(&newRandomCircuit.Nodes[j], float64(r1.Intn(10)+6))
				//}
			}
		}
	}

	return newRandomCircuit
}

// Verify if the parameters nodes are ok to be processed.
// They must be grater than zero, less than the number of nodes and different.
func (circuit *Circuit) verifyNodes(nodeA, nodeB int) error {
	if nodeA < 0 || nodeA >= len(circuit.Nodes) {
		log.Println("circuit.Circuit.verifyNodes")
		return ErrOutOfRange
	}

	if nodeB < 0 || nodeB >= len(circuit.Nodes) {
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
func (circuit *Circuit) SymmetricalConnectNodes(nodeA, nodeB int, weight float64) error {

	err := circuit.verifyNodes(nodeA, nodeB)
	if err != nil {
		log.Println("circuit.Circuit.SymmetricalConnectNodes")
		return err
	}

	circuit.Nodes[nodeA].AddConnection(&circuit.Nodes[nodeB], weight)
	circuit.Nodes[nodeB].AddConnection(&circuit.Nodes[nodeA], weight)

	return nil
}

// AsymmetricalConnectNodes connects two nodes with the given weight in only A to B direction.
func (circuit *Circuit) AsymmetricalConnectNodes(nodeA, nodeB int, weight float64) error {

	err := circuit.verifyNodes(nodeA, nodeB)
	if err != nil {
		log.Println("circuit.Circuit.AsymmetricalConnectNodes")
		return err
	}

	circuit.Nodes[nodeA].AddConnection(&circuit.Nodes[nodeB], weight)

	return nil
}

// GetConnectionWeight returns the weight of the connection between the two nodes in the A to B direction.
func (circuit *Circuit) GetConnectionWeight(nodeA, nodeB int) (float64, error) {
	if nodeA < 0 || nodeA >= len(circuit.Nodes) {
		log.Println("circuit.Circuit.GetConnectionWeight")
		return -1, ErrOutOfRange
	}

	if nodeB < 0 || nodeB >= len(circuit.Nodes) {
		log.Println("circuit.Circuit.GetConnectionWeight")
		return -1, ErrOutOfRange
	}

	connection, err := circuit.Nodes[nodeA].GetConnection(&circuit.Nodes[nodeB])
	if err != nil {
		log.Println("circuit.Circuit.GetConnectionWeight")
		return 0, err
	}

	return connection.GetWeight(), nil
}

// Returns the node number in the circuit.
func (circuit *Circuit) GetNodeNumber(node *place.Place) int {
	for i := 0; i < len(circuit.Nodes); i++ {
		if &circuit.Nodes[i] == node {
			return i
		}
	}
	return -1
}

// Print the circuit nodes and connections with weights.
func (circuit *Circuit) PrintCircuit() string {
	newStrBuider := strings.Builder{}

	for i := 0; i < len(circuit.Nodes); i++ {
		newStrBuider.WriteString("\n\r|" + strconv.Itoa(i) + "| =>")
		connections := circuit.Nodes[i].GetConnections()
		for j := 0; j < len(connections); j++ {
			destNode := connections[j].GetDestination()
			for k := 0; k < len(circuit.Nodes); k++ {
				if &circuit.Nodes[k] == destNode {
					newStrBuider.WriteString(" |" + strconv.Itoa(k))
					break
				}
			}
			weight := connections[j].GetWeight()
			newStrBuider.WriteString(" - " + strconv.FormatFloat(weight, 'f', 2, 64) + "|")
		}
	}

	return newStrBuider.String()
}
