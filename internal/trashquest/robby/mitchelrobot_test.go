package robby

import (
	"log"
	"testing"
)

func TestEvolve(t *testing.T) {
	t.Log("TestEvolve")

	// Generate the inital population
	population := PrepareInitialPopulation(200)

	err := population.Evolve(200)
	if err != nil {
		log.Println("Error evolving population")
		log.Println(err)
		return
	}

	population.GetIndividuals()[0].ReplayASCII()
}

func TestMate(t *testing.T) {
	t.Log("TestMate")

	robot1 := NewMitchelRobot()
	robot2 := NewMitchelRobot()

	child := robot1.Mate(robot2)

	log.Println("Robot 1: ", robot1.Genes.Sequence())
	log.Println("Robot 2: ", robot2.Genes.Sequence())
	log.Println("Child: ", child[0].Genes.Sequence())
}
