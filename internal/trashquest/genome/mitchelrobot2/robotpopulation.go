package mitchelrobot2

import "github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/robot"

func NewMitchelRobot() *robot.Robot {
	newRobot := robot.NewRobot()

	newRobot.Genes = NewGenes()

	return newRobot
}

// CreateNewPopulation returns a new population with initialPopulation number of individuals
func PrepareInitialPopulation(initialPopulation int) *robot.RobotPopulation {

	newPopulation := robot.GenerateInitialPopulation()

	for i := 0; i < initialPopulation; i++ {

		newPopulation.AddIndividual(NewMitchelRobot())
	}

	return &newPopulation
}
