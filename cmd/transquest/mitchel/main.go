package main

import (
	"log"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/robby"
	"github.com/GreenMan-Network/Go-Utils/pkg/fileutils"
)

func main() {
	// Create csv File
	currentTime := time.Now()
	csfFile, err := fileutils.CSVNewFile(currentTime.Format("2006-02-01_15:04:05") + "_" + "mitchelRobot")
	if err != nil {
		log.Println("Error creating csv file")
		log.Println(err)
		return
	}

	header := []string{"Generation", "Average Fitness", "Best Fitness", "Worst Fitness"}
	fileutils.CSVAppendLine(csfFile, header)

	// Generate the inital population
	population := robby.PrepareInitialPopulation(200)

	initialBestFitness := 0.0

	for initialBestFitness < 16 {
		csvLine := population.Evolve()

		initialBestFitness = population.GetIndividuals()[0].Fitness

		fileutils.CSVAppendLine(csfFile, csvLine)
	}

	GenesStr := []string{"Best robot DNA: " + population.GetIndividuals()[0].Genes.String()}
	fileutils.CSVAppendLine(csfFile, GenesStr)

	population.GetIndividuals()[0].ReplayASCII()
}
