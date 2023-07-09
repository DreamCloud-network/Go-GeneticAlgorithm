package mitchelrobot3

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/genome"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/artificial/dna"
)

const mutationChance = 1 // /1000 %
const strandLength = 243

type DoubleGenes struct {
	Strands [][]*dna.Gene
}

func NewGenes() *DoubleGenes {

	var newGenes DoubleGenes

	newGenes.Strands = make([][]*dna.Gene, 2)
	newGenes.Strands[0] = make([]*dna.Gene, strandLength)
	newGenes.Strands[1] = make([]*dna.Gene, strandLength)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Poppulate with random actions
	for _, strand := range newGenes.Strands {
		for i := 0; i < len(strand); i++ {
			// Define if the genes is recessive or dominant
			var newGene *dna.Gene
			if r1.Intn(2) == 0 {
				newGene = dna.NewGene(false)
			} else {
				newGene = dna.NewGene(true)
			}

			action := genome.Action(r1.Intn(int(genome.Pickup + 1)))
			newGene.Code = append(newGene.Code, dna.Codon(action))

			strand[i] = newGene
		}
	}
	return &newGenes
}

func (genes *DoubleGenes) Duplicate() genome.Genes {
	var newGenes DoubleGenes

	newGenes.Strands = make([][]*dna.Gene, 2)

	newGenes.Strands[0] = make([]*dna.Gene, strandLength)
	newGenes.Strands[1] = make([]*dna.Gene, strandLength)

	copy(newGenes.Strands[0], genes.Strands[0])
	copy(newGenes.Strands[1], genes.Strands[1])

	return &newGenes
}

// String returns a string representation of the Genes
func (genes *DoubleGenes) String() string {
	str := ""
	for _, strand := range genes.Strands {
		str += "\n\r-> "
		for _, gene := range strand {
			if gene.Dominance == dna.Recessive {
				str += "a-"
			} else {
				str += "A-"
			}
			str += genome.Action(gene.Code[0]).String() + "|"
		}
	}

	return str
}

// Sequence returns a string representation of the Genes into a sequence of numbers
func (genes *DoubleGenes) Sequence() string {
	str := ""
	for _, strand := range genes.Strands {
		str += "\n\r-> "
		for _, gene := range strand {
			if gene.Dominance == dna.Recessive {
				str += "a"
			} else {
				str += "A"
			}

			str += strconv.Itoa(int(gene.Code[0]))
		}
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

		var newGene *dna.Gene
		if r1.Intn(2) == 0 {
			newGene = dna.NewGene(false)
		} else {
			newGene = dna.NewGene(true)
		}

		action := genome.Action(r1.Intn(int(genome.Pickup + 1)))
		newGene.Code = append(newGene.Code, dna.Codon(action))

		log.Println("Mutated Gene: ", gene.Strands[strandToMutate][posMutation].String())
		gene.Strands[strandToMutate][posMutation] = newGene
		log.Println("New Gene: ", gene.Strands[strandToMutate][posMutation].String())
		log.Println("!!!!!!!!!!!!!!!!")
	}

}

func (genes *DoubleGenes) GetActions() []genome.Action {
	actions := make([]genome.Action, 200)

	for i := 0; i < len(actions); i++ {
		expression := genes.Strands[0][i].ReturnGeneExpressionFromAlleles(genes.Strands[1][i])
		actions[i] = genome.Action(expression.Code[0])
	}
	return actions
}

func (genes *DoubleGenes) Mate(genesPartner genome.Genes) []genome.Genes {

	genesP := genesPartner.(*DoubleGenes)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Create male and female
	maleGamete := make([]*dna.Gene, strandLength)
	femaleGamete := make([]*dna.Gene, strandLength)

	breakPointLenght := r1.Intn(strandLength/2) + 1
	actualAllele := r1.Intn(2)
	breakPointPos := 0

	for cont := range maleGamete {
		if breakPointPos >= breakPointLenght {
			actualAllele = r1.Intn(2)
			breakPointPos = 0
		}

		maleGamete[cont] = genes.Strands[actualAllele][cont].Duplicate()
		femaleGamete[cont] = genesP.Strands[actualAllele][cont].Duplicate()

		breakPointPos++
	}

	// Reproduction
	var newChild DoubleGenes

	newChild.Strands = make([][]*dna.Gene, 2)
	newChild.Strands[0] = maleGamete
	newChild.Strands[1] = femaleGamete

	newChild.mutate()

	return []genome.Genes{&newChild}

}
