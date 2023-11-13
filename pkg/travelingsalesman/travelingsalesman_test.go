package travelingsalesman

import (
	"log"
	"strconv"
	"testing"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/circuit"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/population"
	"github.com/DreamCloud-networkrkrkrk/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

func TestTravelingSalesMan(t *testing.T) {
	t.Log("Testing TravelingSalesMan")

	size := 10

	circuit := circuit.NewTestCircuit(size)

	t.Log("Circuit:", circuit.PrintCircuit())

	// Teste correct solution
	newSalesMan := NewTravelingSalesman(circuit)

	solutionDNA := dnastrand.NewDNAStrand()

	for i := 0; i < size; i++ {

		codons := codons.UintToCodons(0)

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
	/*
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
	*/
}

func TestEvolution(t *testing.T) {
	t.Log("Testing TestEvolution")

	size := 100
	circuit := circuit.NewTestCircuit(size)

	individuals := make([]population.Individual, 200)
	for i := range individuals {
		individuals[i] = NewTravelingSalesman(circuit)
	}
	population := population.NewPopulation(individuals)

	bestFitness := 0.0

	//for bestFitness < float64(size*10) {
	for (population.Generation < 50) && (bestFitness < float64(size*100)) {
		population.Evolve()

		bestFitness = population.Individuals[0].GetFitness()

		if population.HasAlive() && (population.Generation > 100) {
			break
		}
	}

	bestSalesMan := population.Individuals[0].(*TravelingSalesman)
	log.Println("Best SalesMan solution:", bestSalesMan.PrintSolution())

}
