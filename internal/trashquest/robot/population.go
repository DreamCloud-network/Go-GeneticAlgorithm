package robot

import (
	"log"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type RobotPopulation struct {
	robots            []*Robot
	InitialPopulation int
}

// CreateNewPopulation returns a new population with initialPopulation number of individuals
func GenerateInitialPopulation() RobotPopulation {

	newPopulation := RobotPopulation{
		robots:            nil,
		InitialPopulation: 0,
	}

	return newPopulation
}

func (population *RobotPopulation) SetInitialPopulation(initialPopulation []*Robot) {
	population.robots = initialPopulation
	population.InitialPopulation = len(population.robots)
}

func (population *RobotPopulation) AddIndividual(individual *Robot) {
	population.robots = append(population.robots, individual)
	population.InitialPopulation++
}

func (population *RobotPopulation) GetIndividuals() []*Robot {
	return population.robots
}

// GetTotalFitness returns the total fitness of the population
func (population *RobotPopulation) GetTotalAbsFitness() float64 {
	cont := 0
	totalAbsFitness := float64(0)
	for _, individual := range population.robots {
		totalAbsFitness += math.Abs(float64(individual.Points))
		cont++
	}

	return totalAbsFitness
}

// GetTotalFitness returns the total fitness of the population
func (population *RobotPopulation) GetTotalFitness() int {
	cont := 0
	averageFitness := 0
	for _, individual := range population.robots {
		averageFitness += individual.Points
		cont++
		//log.Println(cont, " - Individual fitness: ", individual.Fintness)
		//log.Println(cont, " - Individual sequence: ", individual.Genes.String())
	}

	return averageFitness
}

// GetAverageFitness returns the average fitness of the population
func (population *RobotPopulation) GetAverageFitness() int {

	return (population.GetTotalFitness() / len(population.robots))
}

// Evaluate executes one session for all the individuals in the population.
func (population *RobotPopulation) Evaluate() error {
	var wg sync.WaitGroup

	for _, robot := range population.robots {
		wg.Add(1)
		go func(robot *Robot) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			robot.Evaluate()
		}(robot)
	}

	wg.Wait()

	/*for _, robot := range population.robots {
		robot.Evaluate()
	}*/

	return nil
}

// SortPopulation sorts the population by fitness
func (population *RobotPopulation) sortPopulation() {
	sort.SliceStable(population.robots, func(i, j int) bool {
		return population.robots[i].Points > population.robots[j].Points
	})
}

// PrintIndividualsFintess prints the fitness of every individuals in the population
func (population *RobotPopulation) PrintIndividualsFintess() {
	for ind := 0; ind < len(population.robots); ind++ {
		log.Println(ind, " - Points: ", population.robots[ind].Points)
	}
}

func (population *RobotPopulation) Mate() {
	// Generate a new population
	var newRobots []*Robot

	newRobots = make([]*Robot, 0)

	totalPoints := population.GetTotalAbsFitness()

	pointsVector := make([]int, len(population.robots))
	minPoints := 0

	for i, robot := range population.robots {
		pointsVector[i] = robot.Points
		if robot.Points < minPoints {
			minPoints = robot.Points
		}
	}

	if minPoints < 0 {
		for i, _ := range population.robots {
			posMinPoints := minPoints * -1
			pointsVector[i] = pointsVector[i] + posMinPoints
		}
	}

	// Select a random individual. More points means more probrability to be selected
	probabilidades := make([]float64, len(population.robots))

	for i := range probabilidades {

		if totalPoints > 0 {
			probabilidades[i] = float64(pointsVector[i]) / totalPoints
		} else {
			probabilidades[i] = float64(1) / float64(len(population.robots))
		}
	}

	for len(newRobots) < population.InitialPopulation {

		var partner1 *Robot
		var partner2 *Robot

		numeroAleatorio := rand.Float64()
		somaProbabilidade := 0.0
		position := 0
		for _, p := range probabilidades {
			somaProbabilidade += p
			if numeroAleatorio <= somaProbabilidade {
				break
			}
			position++
		}

		if position >= len(population.robots) {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			partner1 = population.robots[r1.Intn(len(population.robots))]
		} else {
			partner1 = population.robots[position]
		}

		partner2 = partner1

		for partner1.GetID() == partner2.GetID() {
			numeroAleatorio = rand.Float64()
			somaProbabilidade = 0.0
			position = 0
			for _, p := range probabilidades {
				somaProbabilidade += p
				if numeroAleatorio <= somaProbabilidade {
					break
				}
				position++
			}

			if position >= len(population.robots) {
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				partner2 = population.robots[r1.Intn(len(population.robots))]
			} else {
				partner2 = population.robots[position]
			}
		}

		children := partner1.Mate(partner2)
		newRobots = append(newRobots, children...)
	}

	population.robots = newRobots
}

// Evolve generates a new population based on the better individuals of the current population
func (population *RobotPopulation) Evolve(generations int) error {
	for i := 0; i < generations; i++ {

		start := time.Now()
		err := population.Evaluate()
		if err != nil {
			log.Println("mitchelga.Population.Evolve - Error evaluating population")
			return err
		}
		evaluateTime := time.Since(start)

		population.sortPopulation()

		log.Println("=====================")
		log.Println("Generation: ", i)
		log.Println("Avarage fitness: ", population.GetAverageFitness())
		log.Println("Best fitness: ", population.robots[0].Points)
		log.Println("Best trash collected: ", population.robots[0].NumTrashCollected())
		log.Println("=====================")

		start = time.Now()
		population.Mate()
		mateTime := time.Since(start)

		log.Printf("Evaluate took %s", evaluateTime)
		log.Printf("Mate population took %s", mateTime)
		log.Println("=====================")
	}

	population.sortPopulation()

	return nil
}
