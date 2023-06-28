package main

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/mitchelga"
)

func main() {
	// Generate the inital population
	population := mitchelga.GenerateInitialPopulation(200)

	err := population.Evolve(30)
	if err != nil {
		log.Println("Error evolving population")
		log.Println(err)
		return
	}

	bestOne := population.GetBestIndividual()

	log.Println("Best fitness: ", bestOne.Fintness)
	log.Println("Sequence: ", bestOne.Genes.String())
}
