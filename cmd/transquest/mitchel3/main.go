package main

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/genome/mitchelrobot3"
)

func main() {
	// Generate the inital population
	population := mitchelrobot3.PrepareInitialPopulation(200)

	err := population.Evolve(100)
	if err != nil {
		log.Println("Error evolving population")
		log.Println(err)
		return
	}

	population.GetIndividuals()[0].ReplayASCII()
	log.Println(population.GetIndividuals()[0].Genes.String())
}
