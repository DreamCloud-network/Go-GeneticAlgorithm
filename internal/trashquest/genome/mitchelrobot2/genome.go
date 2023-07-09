package mitchelrobot2

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/genome"
)

const mutationChance = 1 // /1000 %
const strandLength = 243

type DoubleGenes struct {
	Strands [][]genome.Action
}

func NewGenes() *DoubleGenes {

	var newGenes DoubleGenes

	newGenes.Strands = make([][]genome.Action, 2)
	newGenes.Strands[0] = make([]genome.Action, strandLength)
	newGenes.Strands[1] = make([]genome.Action, strandLength)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Define if the genes is recessive or dominant
	newGenes.Strands[0][0] = genome.Action(r1.Intn(2))
	newGenes.Strands[1][0] = newGenes.Strands[0][0]

	// Poppulate with random actions
	for _, strand := range newGenes.Strands {
		for i := 0; i < len(strand); i++ {
			strand[i] = genome.Action(r1.Intn(int(genome.Pickup + 1)))
		}
	}
	return &newGenes
}

func (genes *DoubleGenes) Duplicate() genome.Genes {
	var newGenes DoubleGenes

	newGenes.Strands = make([][]genome.Action, 2)
	newGenes.Strands[0] = make([]genome.Action, strandLength)
	newGenes.Strands[1] = make([]genome.Action, strandLength)

	copy(newGenes.Strands[0], genes.Strands[0])
	copy(newGenes.Strands[1], genes.Strands[1])

	return &newGenes
}

// String returns a string representation of the Genes
func (genes *DoubleGenes) String() string {
	str := "||"
	for _, strand := range genes.Strands {
		for cont, action := range strand {
			if cont == 0 {
				if action == 0 {
					str += "Recessive:"
				} else {
					str += "Dominant:"
				}
			} else {
				str += str + "|" + action.String()
			}
		}
		str += "||"
	}

	return str
}

// Sequence returns a string representation of the Genes into a sequence of numbers
func (genes *DoubleGenes) Sequence() string {
	str := "||"
	for _, strand := range genes.Strands {
		for cont, action := range strand {
			if cont == 0 {
				if action == 0 {
					str += "Recessive:"
				} else {
					str += "Dominant:"
				}
			} else {
				str += strconv.Itoa(int(action))
			}
		}
		str += "||"
	}

	return str
}

func (gene *DoubleGenes) mutate() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Mutate with (mutationChance/1000)% chance. Should it be more? Should it be less?
	// 0,01
	if r1.Intn(10000) < mutationChance {
		log.Println("!!! MUTATION !!!")

		strandToMutate := r1.Intn(2)

		posMutation := r1.Intn(len(gene.Strands[strandToMutate])-1) + 1

		gene.Strands[strandToMutate][posMutation] = genome.Action(r1.Intn(int(genome.Pickup + 1)))
	}

}

func (genes *DoubleGenes) Mate(genesPartner genome.Genes) []genome.Genes {

	genesP := genesPartner.(*DoubleGenes)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Create malke gametes
	maleGametes := make([][]genome.Action, 4)

	for cont := 0; cont < 4; cont++ {
		splitPosition := r1.Intn(len(genes.Strands[0])-2) + 1

		maleGametes[cont] = make([]genome.Action, len(genes.Strands[0]))
		for pos := 0; pos < splitPosition; pos++ {
			maleGametes[cont][pos] = genes.Strands[0][pos]
		}

		for pos := splitPosition; pos < len(maleGametes[cont]); pos++ {
			maleGametes[cont][pos] = genes.Strands[1][pos]
		}

		cont++
		splitPosition = r1.Intn(len(genes.Strands[0])-2) + 1

		maleGametes[cont] = make([]genome.Action, len(genes.Strands[0]))
		for pos := 0; pos < splitPosition; pos++ {
			maleGametes[cont][pos] = genes.Strands[1][pos]
		}

		for pos := splitPosition; pos < len(maleGametes[cont]); pos++ {
			maleGametes[cont][pos] = genes.Strands[0][pos]
		}
	}

	// Create female gamete
	femaleGamete := make([]genome.Action, strandLength)

	strandChosen := r1.Intn(2)
	splitPosition := r1.Intn(len(genesP.Strands[0])-2) + 1

	for pos := 0; pos < splitPosition; pos++ {
		femaleGamete[pos] = genesP.Strands[strandChosen][pos]
	}

	if strandChosen == 0 {
		strandChosen = 1
	} else {
		strandChosen = 0
	}

	for pos := splitPosition; pos < len(genesP.Strands[strandChosen]); pos++ {
		femaleGamete[pos] = genesP.Strands[strandChosen][pos]
	}

	// Reproduction
	maleChosen := r1.Intn(4)

	var newChild DoubleGenes

	newChild.Strands = make([][]genome.Action, 2)
	newChild.Strands[0] = maleGametes[maleChosen]
	newChild.Strands[1] = femaleGamete

	newChild.mutate()

	return []genome.Genes{&newChild}

}

func (genes *DoubleGenes) GetActions() []genome.Action {
	if genes.Strands[0][0] == 0 {
		return genes.Strands[0][1:201]
	} else if genes.Strands[1][0] == 0 {
		return genes.Strands[1][1:201]
	}
	return genes.Strands[0][1:201]
}
