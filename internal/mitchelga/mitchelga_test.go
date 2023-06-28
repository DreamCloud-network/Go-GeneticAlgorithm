package mitchelga

import (
	"log"
	"testing"
)

func TestGenesGeneration(t *testing.T) {
	t.Log("Testing Genes Generation")

	genes1 := newGenes()

	t.Log("Genes names: ", genes1.String())
	t.Log("Genes sequence: ", genes1.Sequence())

	genes2 := newGenes()

	t.Log("Genes2 sequence: ", genes2.Sequence())

	if genes2 == genes1 {
		t.Error("Genes are equal")
	}
}

func TestMate(t *testing.T) {
	t.Log("Testing Mate")

	individual1 := NewIndividual()
	individual2 := NewIndividual()

	t.Log("Genes1 sequence: ", individual1.Genes.Sequence())
	t.Log("Genes2 sequence: ", individual2.Genes.Sequence())

	child := individual1.Mate(individual2)

	t.Log("Child sequence: ", child.Genes.Sequence())

	if child.Genes == individual1.Genes || child.Genes == individual2.Genes {
		t.Error("Child is equal to parent")
	}
}

func TestExecuteSession(t *testing.T) {
	t.Log("Testing ExecuteSession")

	individual := NewIndividual()

	err := individual.ExecuteSession()
	if err != nil {
		log.Println("Error executing session")
		log.Println(err)
		return
	}

	log.Println("Individual fitness: ", individual.Fintness)
	log.Println("Individual sequence: ", individual.Genes.String())
}

func TestEvaluate(t *testing.T) {
	t.Log("Testing Evaluate")

	population := GenerateInitialPopulation(200)

	err := population.Evaluate()
	if err != nil {
		log.Println("Error evaluating population")
		log.Println(err)
		return
	}

	population.PrintIndividualsFintess()
	avarageFitness := population.GetAverageFitness()
	bestOne := population.GetBestIndividual()

	log.Println("Avarage fitness: ", avarageFitness)
	log.Println("Best fitness: ", bestOne.Fintness)
}

func TestEvolve(t *testing.T) {
	t.Log("Testing Evolve")

	population := GenerateInitialPopulation(200)

	//population.PrintIndividualsFintess()
	avarageFitness := population.GetAverageFitness()
	bestOne := population.GetBestIndividual()

	t.Log("Initial population")
	t.Log("Avarage fitness: ", avarageFitness)
	t.Log("Best fitness: ", bestOne.Fintness)

	err := population.Evolve(1)
	if err != nil {
		t.Log("Error evolving population")
		t.Error(err)
		return
	}

	//population.PrintIndividualsFintess()
	avarageFitness = population.GetAverageFitness()
	bestOne = population.GetBestIndividual()

	t.Log("Evolved population")
	t.Log("Avarage fitness: ", avarageFitness)
	t.Log("Best fitness: ", bestOne.Fintness)
}
