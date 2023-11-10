package population

import (
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Interface the defines the methods that a individual must have
type Individual interface {
	Run() error
	GetFitness() float64
	Mate(individual Individual) (Individual, error)
	IsAlive() bool
}

// Population struct
type Population struct {
	// Individuals that compose the population
	Individuals []Individual

	// Initial size of the population
	InitialSize int

	// Current generation of the population
	Generation int
}

// Initialize a new population empty population
func NewEmptyPopulation() Population {
	return Population{
		Individuals: nil,
		InitialSize: 0,
		Generation:  0,
	}
}

// Returns a new population with initialPopulation number of individuals randomly generated
func NewPopulation(individuals []Individual) *Population {

	newPopulation := NewEmptyPopulation()

	newPopulation.InitialSize = len(individuals)

	newPopulation.Individuals = individuals

	newPopulation.Generation = 0

	return &newPopulation
}

// Evaluate executes one session for all the individuals in the population.
func (population *Population) Evaluate() {

	var wg sync.WaitGroup

	for individualNum := range population.Individuals {
		wg.Add(1)
		go func(individual Individual) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			individual.Run()

			//err := individual.Run()
			//if err != nil {
			//log.Println("Error running individual.")
			//}

		}(population.Individuals[individualNum])
	}

	wg.Wait()

	/*
		for individualNum := range population.Individuals {
			err := population.Individuals[individualNum].Run()
			if err != nil {
				//log.Println("Error running individual: ", individualNum)
				continue
			}
		}
	*/
}

func (population *Population) selectPartner(fitnessPoints []float64) Individual {
	bestRobbot := len(fitnessPoints) - 1
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select the best of 15 random individuals
	for cont := 0; cont < 15; cont++ {
		randomProbability := r1.Float64()

		accumulatedWeight := 0.0
		index := 0

		for _, weight := range fitnessPoints {
			accumulatedWeight += weight
			if randomProbability < accumulatedWeight {
				break
			}
			index++
		}

		if index < bestRobbot {
			bestRobbot = index
		}

	}

	//log.Println("Index: ", bestRobbot)

	return population.Individuals[bestRobbot]
}

// Prepare the fitness points for the fitness weightened roulette wheel
func (population *Population) prepareFitnessPointsVector() []float64 {
	fitnessPoints := make([]float64, len(population.Individuals))
	minFitness := 0.0
	for i := range fitnessPoints {
		fitnessPoints[i] = population.Individuals[i].GetFitness()
		if fitnessPoints[i] < minFitness {
			minFitness = fitnessPoints[i]
		}
	}

	// Square the fitness points
	totalFitnessPoints := 0.0
	for i := range fitnessPoints {
		if minFitness < 0 {
			fitnessPoints[i] += math.Abs(minFitness)
		} else {
			fitnessPoints[i] -= minFitness
		}
		//fitnessPoints[i] = math.Pow(fitnessPoints[i], 2)
		totalFitnessPoints += fitnessPoints[i]
	}

	// Normalize the fitness points
	for i := range fitnessPoints {
		fitnessPoints[i] = fitnessPoints[i] / totalFitnessPoints
	}

	return fitnessPoints
}

// This Mate function implements the same algorithm described by Mitchell in his book.
// It generate the next generations with the same size of the initial population.
func (population *Population) Mate() {

	// Prepare the fitness points for the roulette wheel
	fitnessPoints := population.prepareFitnessPointsVector()

	// Generate a new population
	newIndividuals := make([]Individual, 0, len(population.Individuals))

	for len(newIndividuals) < population.InitialSize {

		// Select first partner
		//log.Println("Partner 1:")
		partner1 := population.selectPartner(fitnessPoints)

		// Select second partner
		partner2 := partner1

		// Make sure that the partners are different
		//log.Println("Partner 2:")
		for partner1 == partner2 {
			partner2 = population.selectPartner(fitnessPoints)
		}

		child, err := partner1.Mate(partner2)
		if (err != nil) || (child == nil) {
			log.Println("Error mating individuals.")
		} else {
			newIndividuals = append(newIndividuals, child)
		}
	}

	population.Individuals = newIndividuals
}

// Returns the number o alive individuals in the population
func (population *Population) NumAlive() int {
	numAlive := 0
	for _, individual := range population.Individuals {
		if individual.IsAlive() {
			numAlive++
		}
	}

	return numAlive
}

// Return true if the population has at least one alive individual
func (population *Population) HasAlive() bool {
	for _, individual := range population.Individuals {
		if individual.IsAlive() {
			return true
		}
	}
	return false
}

// Removes all the dead individuals from the population
func (population *Population) removeDeads() {
	livePopulation := make([]Individual, 0, len(population.Individuals))

	for _, individual := range population.Individuals {
		if individual.IsAlive() {
			livePopulation = append(livePopulation, individual)
		}
	}

	population.Individuals = livePopulation
}

// SortPopulation sorts the population by fitness
func (population *Population) sortPopulation() {
	sort.SliceStable(population.Individuals, func(i, j int) bool {
		return population.Individuals[i].GetFitness() > population.Individuals[j].GetFitness()
	})
}

// GetTotalFitness returns the total fitness of the population
func (population *Population) GetTotalFitness() float64 {
	totalFitness := 0.0
	for _, individual := range population.Individuals {
		totalFitness += individual.GetFitness()
		//log.Println(cont, " - Individual fitness: ", individual.Fintness)
		//log.Println(cont, " - Individual sequence: ", individual.Genes.String())
	}

	return totalFitness
}

// GetAverageFitness returns the average fitness of the population
func (population *Population) GetAverageFitness() float64 {

	return (population.GetTotalFitness() / float64(len(population.Individuals)))
}

// Evolve generates a new population based on the better individuals of the current population. This function returns a string with the information of the generation.
// The string contains the generation number, the average fitness of the population, the best fitness and the worst fitness:
// GenerationNumber, AverageFitness, BestFitness, WorstFitness
func (population *Population) Evolve() []string {

	// Just execute the mate after the first generation
	start := time.Now()
	if population.Generation > 0 {
		population.Mate()
	}
	mateTime := time.Since(start)

	start = time.Now()
	population.Evaluate()
	evaluateTime := time.Since(start)

	population.removeDeads()

	population.sortPopulation()

	lastIndividual := population.Individuals[len(population.Individuals)-1]
	record := []string{strconv.Itoa(population.Generation), strconv.FormatFloat(population.GetAverageFitness(), 'f', 2, 64),
		strconv.FormatFloat(population.Individuals[0].GetFitness(), 'f', 2, 64), strconv.FormatFloat(lastIndividual.GetFitness(), 'f', 2, 64)}

	log.Println("=====================")
	log.Println("Generation: ", record[0])
	log.Println("Population size: ", len(population.Individuals), " - Alive: ", population.NumAlive())
	log.Println("Avarage fitness: ", record[1])
	log.Println("Best fitness: ", record[2])
	log.Println("Worst fitness: ", record[3])
	log.Println("=====================")
	log.Printf("Evaluate took %s", evaluateTime)
	log.Printf("Mate population took %s", mateTime)
	log.Println("=====================")

	population.Generation++

	return record
}
