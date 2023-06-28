package mitchelga

import (
	"log"
	"math"
	"strconv"
)

type Population struct {
	Individuals []*Individual
}

// CreateNewPopulation returns a new population with initialPopulation number of individuals
func GenerateInitialPopulation(initialPopulation int) Population {

	newPopulation := Population{
		Individuals: make([]*Individual, initialPopulation),
	}

	for i := 0; i < initialPopulation; i++ {
		newPopulation.Individuals[i] = NewIndividual()
	}

	return newPopulation
}

// Evaluate executes one session for all the individuals in the population.
func (population *Population) Evaluate() error {

	for _, individual := range population.Individuals {
		err := individual.ExecuteSession()
		if err != nil {
			log.Println("mitchelga.Population.Evaluate - Error executing session")
			return err
		}
	}

	return nil
}

// PrintIndividualsFintess prints the fitness of every individuals in the population
func (population *Population) PrintIndividualsFintess() {
	for ind := 0; ind < len(population.Individuals); ind++ {
		log.Printf(strconv.Itoa(ind)+" - Fitness: %f", population.Individuals[ind].Fintness)
	}
}

// GetAverageFitness returns the average fitness of the population
func (population *Population) GetAverageFitness() float64 {
	var averageFitness float64

	cont := 0

	for _, individual := range population.Individuals {
		averageFitness += individual.Fintness
		cont++
		//log.Println(cont, " - Individual fitness: ", individual.Fintness)
		//log.Println(cont, " - Individual sequence: ", individual.Genes.String())
	}

	averageFitness = averageFitness / float64(len(population.Individuals))

	return averageFitness
}

// PullBetterIndividual returns the individual with the better fitness and remove from the population
func (population *Population) PullFirstIndividual() *Individual {
	betterIndividual := population.Individuals[0]
	population.Individuals = population.Individuals[1 : len(population.Individuals)-1]

	return betterIndividual
}

// GetBetterIndividual returns the individual with the better fitness
func (population *Population) GetBestIndividual() *Individual {
	betterIndividual := 0
	betterFitness := population.Individuals[betterIndividual].Fintness

	for ind := 0; ind < len(population.Individuals); ind++ {
		if population.Individuals[ind].Fintness > betterFitness {
			betterIndividual = ind
			betterFitness = population.Individuals[ind].Fintness
		}
	}

	return population.Individuals[betterIndividual]
}

// SortPopulation sorts the population by fitness
func (population *Population) sortPopulation() {
	for i := 0; i < len(population.Individuals); i++ {
		for j := 0; j < len(population.Individuals)-1; j++ {
			if population.Individuals[j].Fintness < population.Individuals[j+1].Fintness {
				population.Individuals[j], population.Individuals[j+1] = population.Individuals[j+1], population.Individuals[j]
			}
		}
	}
}

func (population *Population) Mate() {
	// Generate a new population
	newPopulation := GenerateInitialPopulation(len(population.Individuals))

	// Pull the better individual
	population.sortPopulation()
	betterIndividual := population.PullFirstIndividual()

	for cont := 0; cont < len(newPopulation.Individuals); cont++ {
		betterIndividualMate := population.GetBestIndividual()

		newPopulation.Individuals[cont] = betterIndividual.Mate(betterIndividualMate)

		if math.Abs(betterIndividualMate.Fintness) > math.Abs(betterIndividual.Fintness*2) {
			betterIndividual = population.PullFirstIndividual()
		}
	}

	*population = newPopulation
}

// Evolve generates a new population based on the better individuals of the current population
func (population *Population) Evolve(generations int) error {
	for i := 0; i < generations; i++ {

		//start = time.Now()
		population.Mate()
		//elapsed = time.Since(start)
		//log.Printf("Mate population took %s", elapsed)

		//start := time.Now()
		err := population.Evaluate()
		if err != nil {
			log.Println("mitchelga.Population.Evolve - Error evaluating population")
			return err
		}
		//elapsed := time.Since(start)
		//log.Printf("Evaluate took %s", elapsed)

		log.Println("Avarage fitness: ", population.GetAverageFitness())
	}

	return nil
}
