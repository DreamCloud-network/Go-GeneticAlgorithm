package main

import (
	"log"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/internal/trashquest/tobby"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/population"
	"github.com/DreamCloud-network/Go-Utils/pkg/fileutils"
)

func main() {
	// Create csv File
	currentTime := time.Now()
	csfFile, err := fileutils.CSVNewFile(currentTime.Format("2006-02-01_15:04:05") + "_" + "TobbyRobot")
	if err != nil {
		log.Println("Error creating csv file")
		log.Println(err)
		return
	}
	defer csfFile.Flush()

	header := []string{"Generation", "Average Fitness", "Best Fitness", "Worst Fitness"}
	fileutils.CSVAppendLine(csfFile, header)

	// Generate the inital population
	individuals := make([]population.Individual, 200)
	for i := range individuals {
		individuals[i] = tobby.NewRandomTobby()
	}
	population := population.NewPopulation(individuals)

	initialBestFitness := 0.0

	for initialBestFitness <= 500 {
		csvLine := population.Evolve()

		initialBestFitness = population.Individuals[0].GetFitness()

		fileutils.CSVAppendLine(csfFile, csvLine)
	}

	bestTobby := population.Individuals[0].(*tobby.Tobby)

	//GenesStr := []string{"Best robot DNA: " + bestTobby.Genes.ActionString()}
	//fileutils.CSVAppendLine(csfFile, GenesStr)

	bestTobby.ReplayASCII()
}
