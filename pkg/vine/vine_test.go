package vine

import (
	"testing"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/circuit"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/things"
)

func TestVin(t *testing.T) {
	t.Log("TestVin")

	environment := circuit.NewTestCircuit(5)

	t.Log("Environment: " + environment.PrintCircuit())

	newEarth := things.NewEarth(1)
	environment.Nodes[0].AddThing(newEarth)

	vineSeed := NewVineSeed()

	vineSeed.SetPosition(&environment.Nodes[0])

	for vineSeed.alive {
		RunVine()
	}
}
