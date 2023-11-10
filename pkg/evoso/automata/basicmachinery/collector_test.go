package basicmachinery

import (
	"testing"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

func TestCollectorGene(t *testing.T) {
	t.Log("TestCollectorGene")

	newCollectorGene := newCollectorGene(environment.Unknown)
	t.Log("Collector Gene: ", newCollectorGene.String())

	sleepTime, thing, err := decodeCollectorGene(newCollectorGene)
	if err != nil {
		t.Error("Error decoding collector gene: ", err)
		return
	}

	t.Log("Sleep time: ", sleepTime)
	t.Log("Thing: ", thing.String())

}
