package main

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/robby"
)

func main() {
	// Generate the inital population
	population := robby.PrepareInitialPopulation(200)

	err := population.Evolve(2000)
	if err != nil {
		log.Println("Error evolving population")
		log.Println(err)
		return
	}

	population.GetIndividuals()[0].ReplayASCII()
	//log.Println(population.GetIndividuals()[0].Genes.String())
}
