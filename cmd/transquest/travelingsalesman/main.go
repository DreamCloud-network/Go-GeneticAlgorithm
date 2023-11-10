package main

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place/circuit"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/population"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/travelingsalesman"
)

func main() {
	circuitSize := 100

	circuit := circuit.NewTestCircuit(circuitSize)

	individuals := make([]population.Individual, 200)
	for i := range individuals {
		individuals[i] = travelingsalesman.NewTravelingSalesman(circuit)
	}
	population := population.NewPopulation(individuals)

	bestFitness := 0.0

	for bestFitness < float64(circuitSize*100) {
		population.Evolve()

		bestFitness = population.Individuals[0].GetFitness()

		//if population.HasAlive() && (population.Generation > 2000) {
		//break
		//}
	}

	bestSalesMan := population.Individuals[0].(*travelingsalesman.TravelingSalesman)
	log.Println("Best SalesMan solution:", bestSalesMan.PrintSolution())
}
