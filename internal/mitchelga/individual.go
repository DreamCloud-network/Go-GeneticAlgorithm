package mitchelga

import (
	"log"
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest"
)

type Individual struct {
	Genes    *Genes
	Fintness float64
}

func NewIndividual() *Individual {

	newIndividual := Individual{
		Genes:    newGenes(),
		Fintness: 0,
	}

	return &newIndividual
}

// Mate returns a new Individual with the genes of the parents
func (individual *Individual) Mate(partner *Individual) *Individual {

	newIndividual := Individual{
		Genes:    newGenes(),
		Fintness: (individual.Fintness + partner.Fintness) / 2, // I don´t know if this is true. Should it be zero?
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	firstHalf := r1.Intn(2)

	for i := 0; i < len(newIndividual.Genes.Actions); i++ {

		if firstHalf == 0 {
			if i < len(newIndividual.Genes.Actions)/2 {
				newIndividual.Genes.Actions[i] = individual.Genes.Actions[i]
			} else {
				newIndividual.Genes.Actions[i] = partner.Genes.Actions[i]
			}
		} else {
			if i < len(newIndividual.Genes.Actions)/2 {
				newIndividual.Genes.Actions[i] = partner.Genes.Actions[i]
			} else {
				newIndividual.Genes.Actions[i] = individual.Genes.Actions[i]
			}
		}

		// Mutate with 0.1% chance. Should it be more? Souhld it be less?
		if r1.Intn(1000) == 0 {
			individual.Genes.Actions[i] = trashquest.Action(r1.Intn(int(trashquest.Pickup + 1)))
		}
	}

	return &newIndividual
}

// ExecuteSession executes a session of 100 games and returns the fitness, that is the average of the points.
func (individual *Individual) ExecuteSession() error {

	trashBoard := trashquest.NewTrashBoard(10)
	individual.Fintness = 0

	for sessionNum := 0; sessionNum < 100; sessionNum++ {

		trashBoard.Board.CleamItems()
		trashBoard.PopulateBoardWithTrash()

		player := trashquest.NewTrashPlayer(trashBoard)

		err := player.MoveSequence(individual.Genes.Actions, false)
		if err != nil {
			log.Println("mitchelga.Individual.ExecuteSession - Error moving player with sequence")
			return err
		}

		individual.Fintness += float64(player.Points)
	}

	individual.Fintness = individual.Fintness / 100
	return nil
}
