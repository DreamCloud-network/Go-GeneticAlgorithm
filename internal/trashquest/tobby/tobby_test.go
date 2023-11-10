package tobby

import (
	"log"
	"testing"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/population"
)

func TestTobby(t *testing.T) {
	t.Log("TestTobby")

	tobby := NewRandomTobby()

	log.Println("Tobby Actions: " + tobby.Genes.ActionString())

	err := tobby.ReplayASCII()
	if err != nil {
		t.Error("Error replaying the robot: ", err)
	}
}

func TestMate(t *testing.T) {
	t.Log("TestMate")

	robot1 := NewRandomTobby()
	robot2 := NewRandomTobby()

	child, err := robot1.Mate(robot2)
	if err != nil {
		t.Error("Error mating the robots: ", err)
		return
	}

	robotChild := child.(*Tobby)

	log.Println("Robot 1: ", robot1.Genes.Actions.String())
	log.Println("Robot 2: ", robot2.Genes.Actions.String())
	log.Println("Child: ", robotChild.Genes.Actions.String())
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
