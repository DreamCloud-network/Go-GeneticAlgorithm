package robby

import (
	"math/rand"
	"strconv"
	"time"
)

const mutationChance = 0.01 // /%
const strandLength = 243

type MitchelGene struct {
	Strand [strandLength]Action
}

func NewGenes() MitchelGene {

	newGenes := MitchelGene{
		Strand: [strandLength]Action{},
	}
	return newGenes
}

func (genes *MitchelGene) PopulateRandomActions() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < len(genes.Strand); i++ {
		genes.Strand[i] = Action(r1.Intn(int(DoNothing + 1)))
	}
}

// String returns a string representation of the Genes
func (genes *MitchelGene) String() string {
	var str string
	for i := 0; i < len(genes.Strand); i++ {
		if i == 0 {
			str = Action(genes.Strand[i]).String()
		} else {
			newAction := Action(genes.Strand[i]).String()
			str = str + "|" + newAction
		}

	}
	return str
}

func (genes *MitchelGene) GetDNA() [strandLength]Action {
	return genes.Strand
}

// Sequence returns a string representation of the Genes into a sequence of numbers
func (genes *MitchelGene) Sequence() string {
	var str string
	for _, gene := range genes.Strand {
		str += strconv.Itoa(int(gene))

	}
	return str
}

func (gene *MitchelGene) mutate() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for pos := range gene.Strand {
		// Mutate with mutationChance% chance. Should it be more? Should it be less?
		// 0,01%
		if r1.Float64() < mutationChance {
			//log.Println("!!! MUTATION !!!")
			gene.Strand[pos] = Action(r1.Intn(int(DoNothing + 1)))
		}
	}

}

func (genes *MitchelGene) Mate(genesPartner *MitchelGene) []MitchelGene {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	genesChild1 := NewGenes()
	genesChild2 := NewGenes()

	crossOverPoint := r1.Intn(len(genes.Strand)-2) + 1

	// Cross over
	for cont := 0; cont < crossOverPoint; cont++ {
		genesChild1.Strand[cont] = genes.Strand[cont]
		genesChild2.Strand[cont] = genesPartner.Strand[cont]
	}

	for cont := crossOverPoint; cont < len(genes.Strand); cont++ {
		genesChild1.Strand[cont] = genesPartner.Strand[cont]
		genesChild2.Strand[cont] = genes.Strand[cont]
	}

	genesChild1.mutate()
	genesChild2.mutate()

	return []MitchelGene{genesChild1, genesChild2}
}

// Return the action based in the robot activations in the position (positionSignature).
func (genes *MitchelGene) GetAction(positionSignature int) Action {
	if positionSignature >= len(genes.Strand) {
		return DoNothing
	}

	return Action(genes.Strand[positionSignature])
}
