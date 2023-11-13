package travelingsalesman

import (
	"log"
	"strconv"
	"testing"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/population"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/travelingsalesman/circuit"
)

func TestTravelingSalesMan(t *testing.T) {
	t.Log("Testing TravelingSalesMan")

	circuit := circuit.NewCircuit(4)

	circuit.SymmetricalConnectNodes(0, 1, 1)
	circuit.SymmetricalConnectNodes(1, 2, 1)
	circuit.SymmetricalConnectNodes(2, 3, 1)
	circuit.SymmetricalConnectNodes(3, 0, 1)
	circuit.SymmetricalConnectNodes(0, 2, 5)

	t.Log("Circuit:", circuit)

	// Teste correct solution
	newSalesMan := NewTravelingSalesman(circuit, 0, 0)

	solutionDNA := dnastrand.NewDNAStrand()

	for i := 0; i < 4; i++ {

		val := i + 1
		if val >= 4 {
			val = 0
		}
		codons := codons.UintToCodons(uint(val))

		newGenesBuilder := genes.NewBuilder()

		for _, codon := range codons {
			newGenesBuilder.AddCode(codon)
		}

		newGene := newGenesBuilder.Finish()

		solutionDNA.AddGene(newGene)
	}

	newSalesMan.genome.nextPos = solutionDNA
	t.Log("New SalesMan solution:", newSalesMan.PrintSolution())

	err := newSalesMan.Run()
	if err != nil {
		t.Error("Run returned error:", err)
		return
	}

	if newSalesMan.alaive {
		t.Log("Objective reached")
		t.Log("Fitness:", strconv.FormatFloat(newSalesMan.fitness, 'f', 2, 64))
	}

	// Teste Random solution
	newSalesMan = NewTravelingSalesman(circuit, 0, 0)

	if newSalesMan == nil {
		t.Error("NewTravelingSalesMan returned nil")
		return
	}

	t.Log("New SalesMan solution:", newSalesMan.PrintSolution())

	err = newSalesMan.Run()
	if err != nil {
		t.Error("Run returned error:", err)
		return
	}

	if newSalesMan.alaive {
		t.Log("Objective reached")
		t.Log("Fitness:", strconv.FormatFloat(newSalesMan.fitness, 'f', 2, 64))
	}

}

func TestTestCircuit(t *testing.T) {
	t.Log("Testing TestCircuit")

	circuit := circuit.TestSolution(5)

	t.Log("Circuit:", circuit)
}

func TestEvolution(t *testing.T) {
	t.Log("Testing TestEvolution")

	circuitSize := 10

	circuit := circuit.TestSolution(circuitSize)

	individuals := make([]population.Individual, 200)
	for i := range individuals {
		individuals[i] = NewTravelingSalesman(circuit, 0, 0)
	}
	population := population.NewPopulation(individuals)

	bestFitness := 0.0

	for bestFitness < float64(circuitSize*100) {
		population.Evolve()

		bestFitness = population.Individuals[0].GetFitness()

		if population.HasAlive() && (population.Generation > 2000) {
			break
		}
	}

	bestSalesMan := population.Individuals[0].(*TravelingSalesman)
	log.Println("Best SalesMan solution:", bestSalesMan.PrintSolution())

}
