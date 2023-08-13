package tobby

import (
	"testing"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/population"
)

func TestTobby(t *testing.T) {
	t.Log("TestTobby")

	tobby := NewTobby()

	tobby.Genes.RamdomizeGenes()

	err := tobby.ReplayASCII()
	if err != nil {
		t.Error("Error replaying the robot: ", err)
	}
}

func TestEvolve(t *testing.T) {
	t.Log("TestEvolve")

	// Generate the inital population
	individuals := make([]population.Individual, 200)
	for i := range individuals {
		individuals[i] = NewRandomTobby()
	}
	population := population.NewPopulation(individuals)

	for cont := 0; cont < 100; cont++ {
		population.Evolve()
	}
}

/*
func TestMate(t *testing.T) {
	t.Log("TestMate")

	robot1 := NewMitchelRobot()
	robot2 := NewMitchelRobot()

	child := robot1.Mate(robot2)

	log.Println("Robot 1: ", robot1.Genes.Sequence())
	log.Println("Robot 2: ", robot2.Genes.Sequence())
	log.Println("Child: ", child[0].Genes.Sequence())
}
*/
