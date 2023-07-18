package robby

import (
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

func NewMitchelRobot() *Robby {
	newRobot := NewRobot()

	newRobot.Genes.PopulateRandomActions()

	return &newRobot
}

// CreateNewPopulation returns a new population with initialPopulation number of individuals
func PrepareInitialPopulation(initialPopulation int) *RobotPopulation {

	newPopulation := GenerateInitialPopulation()

	for i := 0; i < initialPopulation; i++ {

		newPopulation.AddIndividual(NewMitchelRobot())
	}

	return &newPopulation
}

type RobotPopulation struct {
	robots            []*Robby
	InitialPopulation int
	Generation        int
}

// CreateNewPopulation returns a new population with initialPopulation number of individuals
func GenerateInitialPopulation() RobotPopulation {

	newPopulation := RobotPopulation{
		robots:            nil,
		InitialPopulation: 0,
		Generation:        0,
	}

	return newPopulation
}

func (population *RobotPopulation) SetInitialPopulation(initialPopulation []*Robby) {
	population.robots = initialPopulation
	population.InitialPopulation = len(population.robots)
}

func (population *RobotPopulation) AddIndividual(individual *Robby) {
	population.robots = append(population.robots, individual)
	population.InitialPopulation++
}

func (population *RobotPopulation) GetIndividuals() []*Robby {
	return population.robots
}

// GetTotalFitness returns the total fitness of the population
func (population *RobotPopulation) GetTotalAbsFitness() float64 {
	totalAbsFitness := float64(0)
	for _, individual := range population.robots {
		totalAbsFitness += math.Abs(individual.Fitness)
	}

	return totalAbsFitness
}

// GetTotalFitness returns the total fitness of the population
func (population *RobotPopulation) GetTotalFitness() float64 {
	totalFitness := 0.0
	for _, individual := range population.robots {
		totalFitness += individual.Fitness
		//log.Println(cont, " - Individual fitness: ", individual.Fintness)
		//log.Println(cont, " - Individual sequence: ", individual.Genes.String())
	}

	return totalFitness
}

// GetAverageFitness returns the average fitness of the population
func (population *RobotPopulation) GetAverageFitness() float64 {

	return (population.GetTotalFitness() / float64(len(population.robots)))
}

// Evaluate executes one session for all the individuals in the population.
func (population *RobotPopulation) Evaluate() {
	var wg sync.WaitGroup

	for robotNum := range population.robots {
		wg.Add(1)
		go func(robot *Robby) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			for cont := 0; cont < 100; cont++ {
				robot.Run()
			}

			robot.Fitness = robot.Fitness / 100

		}(population.robots[robotNum])
	}

	wg.Wait()

	/*for robotNum := range population.robots {
		for cont := 0; cont < 100; cont++ {
			population.robots[robotNum].Run()
		}

		population.robots[robotNum].Fitness = population.robots[robotNum].Fitness / 100
	}*/
}

// SortPopulation sorts the population by fitness
func (population *RobotPopulation) sortPopulation() {
	sort.SliceStable(population.robots, func(i, j int) bool {
		return population.robots[i].Fitness > population.robots[j].Fitness
	})
}

// PrintIndividualsFintess prints the fitness of every individuals in the population
func (population *RobotPopulation) PrintIndividualsFintess() {
	for ind := 0; ind < len(population.robots); ind++ {
		log.Println(ind, " - Points: ", population.robots[ind].Fitness)
	}
}

/*
// This Mate function version kills the half worst population and keep exactly the same best half for the next iteration.
func (population *RobotPopulation) Mate() {
	// Generate a new population
	var newRobots []*Robby

	fitnessPoints := make([]float64, len(population.robots)/2)
	minFitness := 0.0
	for i := range fitnessPoints {
		fitnessPoints[i] = population.robots[i].Fitness
		if fitnessPoints[i] < minFitness {
			minFitness = fitnessPoints[i]
		}
	}

	// Squaring normalization of the fitness points
	totalFitnessPoints := 0.0
	for i := range fitnessPoints {
		fitnessPoints[i] -= minFitness
		fitnessPoints[i] = math.Pow(fitnessPoints[i], 2)
		totalFitnessPoints += fitnessPoints[i]
	}

	for i := range fitnessPoints {
		fitnessPoints[i] = fitnessPoints[i] / totalFitnessPoints
	}

	// Fills the new population with the best half actual individuals
	for i := 0; i < len(population.robots)/2; i++ {
		newRobots = append(newRobots, population.robots[i])
	}

	for len(newRobots) < population.InitialPopulation {

		positionPartner1 := 0
		positionPartner2 := 0

		// Select first partner
		bestRobbot := len(fitnessPoints) - 1
		for cont := 0; cont < 15; cont++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r := r1.Float64()
			accumulatedWeight := 0.0

			for i, weight := range fitnessPoints {
				accumulatedWeight += weight
				if r < accumulatedWeight {
					if i < bestRobbot {
						bestRobbot = i
					}
					break
				}
			}

		}

		positionPartner1 = bestRobbot

		// Select second partner
		bestRobbot = len(fitnessPoints) - 1
		for cont := 0; cont < 15; cont++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r := r1.Float64()
			accumulatedWeight := 0.0
			for i, weight := range fitnessPoints {
				accumulatedWeight += weight
				if r < accumulatedWeight {
					if i < bestRobbot {
						bestRobbot = i
					}
					break
				}
			}
		}
		positionPartner2 = bestRobbot

		if positionPartner1 == positionPartner2 {
			positionPartner2++
		}

		//log.Println("Partner 1: ", positionPartner1, "- Partner 2: ", positionPartner2)

		partner1 := population.robots[positionPartner1]
		partner2 := population.robots[positionPartner2]

		children := partner1.Mate(partner2)
		newRobots = append(newRobots, children...)
	}

	population.robots = newRobots
}
*/

// This Mate function implements the same algorithm described by Mitchell in his book.
func (population *RobotPopulation) Mate() {
	// Generate a new population
	var newRobots []*Robby

	fitnessPoints := make([]float64, len(population.robots))
	minFitness := 0.0
	for i := range fitnessPoints {
		fitnessPoints[i] = population.robots[i].Fitness
		if fitnessPoints[i] < minFitness {
			minFitness = fitnessPoints[i]
		}
	}

	// Squaring normalization of the fitness points
	totalFitnessPoints := 0.0
	for i := range fitnessPoints {
		fitnessPoints[i] -= minFitness
		fitnessPoints[i] = math.Pow(fitnessPoints[i], 2)
		totalFitnessPoints += fitnessPoints[i]
	}

	for i := range fitnessPoints {
		fitnessPoints[i] = fitnessPoints[i] / totalFitnessPoints
	}

	for len(newRobots) < population.InitialPopulation {

		positionPartner1 := 0
		positionPartner2 := 0

		// Select first partner
		bestRobbot := len(fitnessPoints) - 1
		for cont := 0; cont < 15; cont++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r := r1.Float64()
			accumulatedWeight := 0.0

			for i, weight := range fitnessPoints {
				accumulatedWeight += weight
				if r < accumulatedWeight {
					if i < bestRobbot {
						bestRobbot = i
					}
					break
				}
			}

		}

		positionPartner1 = bestRobbot

		// Select second partner
		bestRobbot = len(fitnessPoints) - 1
		for cont := 0; cont < 15; cont++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r := r1.Float64()
			accumulatedWeight := 0.0
			for i, weight := range fitnessPoints {
				accumulatedWeight += weight
				if r < accumulatedWeight {
					if i < bestRobbot {
						bestRobbot = i
					}
					break
				}
			}
		}
		positionPartner2 = bestRobbot

		if positionPartner1 == positionPartner2 {
			positionPartner2++
		}

		//log.Println("Partner 1: ", positionPartner1, "- Partner 2: ", positionPartner2)

		partner1 := population.robots[positionPartner1]
		partner2 := population.robots[positionPartner2]

		children := partner1.Mate(partner2)
		newRobots = append(newRobots, children...)
	}

	population.robots = newRobots
}

// Evolve generates a new population based on the better individuals of the current population. This function returns a string with the information of the generation.
// The string contains the generation number, the average fitness of the population, the best fitness and the worst fitness:
// GenerationNumber, AverageFitness, BestFitness, WorstFitness
func (population *RobotPopulation) Evolve() []string {

	// Just execute the mate after the first generation
	start := time.Now()
	if population.Generation > 0 {
		population.Mate()
	}
	mateTime := time.Since(start)

	start = time.Now()
	population.Evaluate()
	evaluateTime := time.Since(start)

	population.sortPopulation()

	record := []string{strconv.Itoa(population.Generation), strconv.FormatFloat(population.GetAverageFitness(), 'f', 2, 64),
		strconv.FormatFloat(population.robots[0].Fitness, 'f', 2, 64), strconv.FormatFloat(population.robots[199].Fitness, 'f', 2, 64)}

	log.Println("=====================")
	log.Println("Generation: ", record[0])
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
