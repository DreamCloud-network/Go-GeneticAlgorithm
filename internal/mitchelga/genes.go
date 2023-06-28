package mitchelga

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest"
)

type Genes struct {
	Actions []trashquest.Action
}

func newGenes() *Genes {

	newGenes := Genes{
		Actions: make([]trashquest.Action, 243),
	}

	// Poppulate with holda action
	/*for i := 0; i < len(newGenes.Actions); i++ {
		newGenes.Actions[i] = board.DoNothing
	}*/

	// Poppulate with random actions
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < len(newGenes.Actions); i++ {
		newGenes.Actions[i] = trashquest.Action(r1.Intn(int(trashquest.Pickup + 1)))
	}
	return &newGenes
}

// String returns a string representation of the Genes
func (genes *Genes) String() string {
	var str string
	for i := 0; i < len(genes.Actions); i++ {
		if i == 0 {
			str, _ = genes.Actions[i].String()
		} else {
			newAction, _ := genes.Actions[i].String()
			str = str + "|" + newAction
		}

	}
	return str
}

// Sequence returns a string representation of the Genes into a sequence of numbers
func (genes *Genes) Sequence() string {
	var str string
	for i := 0; i < len(genes.Actions); i++ {
		newAction := int(genes.Actions[i])
		str += strconv.Itoa(newAction)

	}
	return str
}
