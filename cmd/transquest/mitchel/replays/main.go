package main

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/robby"
)

//var fileName = "mitchelRobot_original.genes"

//var fileName = "mitchelRobot_WithoutDoNotihing.genes"

var fileName = "mitchelRobot_Death.genes"

func main() {
	loadedGenes, err := robby.LoadGnomeFile(fileName)
	if err != nil {
		log.Println("Error reading file: ", err)
		return
	}

	robot := robby.NewMitchelRobot()
	robot.Genes.Strand = loadedGenes.Strand

	robot.ReplayASCII()
}
