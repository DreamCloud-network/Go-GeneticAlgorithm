package circuit

import (
	"log"
	"testing"
)

func TestCircuit(t *testing.T) {
	circuit := NewCircuit(10)

	log.Println("Circuit: ", circuit)

	circuit.SymmetricalConnectNodes(0, 1, 10)

	log.Println("Circuit: ", circuit)

}
