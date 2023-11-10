package circuit

import (
	"log"
	"testing"
)

func TestCircuit(t *testing.T) {
	circuit := NewTestCircuit(10)

	log.Println("Circuit:\n\r", circuit.PrintCircuit())

}
