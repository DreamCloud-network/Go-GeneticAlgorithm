package mitchelrobot

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/internal/trashquest/genome"
)

const mutationChance = 1 // /1000 %

type MitchelGene struct {
	Actions []genome.Action
}

func NewGenes() *MitchelGene {

	newGenes := MitchelGene{
		Actions: make([]genome.Action, 243),
	}

	// Poppulate with random actions
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < len(newGenes.Actions); i++ {
		newGenes.Actions[i] = genome.Action(r1.Intn(int(genome.Pickup + 1)))
	}
	return &newGenes
}

// String returns a string representation of the Genes
func (genes *MitchelGene) String() string {
	var str string
	for i := 0; i < len(genes.Actions); i++ {
		if i == 0 {
			str = genes.Actions[i].String()
		} else {
			newAction := genes.Actions[i].String()
			str = str + "|" + newAction
		}

	}
	return str
}

func (genes *MitchelGene) Duplicate() genome.Genes {
	var newGenes MitchelGene

	newGenes.Actions = make([]genome.Action, 243)
	copy(newGenes.Actions, genes.Actions)

	return &newGenes
}

// Sequence returns a string representation of the Genes into a sequence of numbers
func (genes *MitchelGene) Sequence() string {
	var str string
	for _, action := range genes.Actions {
		str += strconv.Itoa(int(action))

	}
	return str
}

func (gene *MitchelGene) mutate() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Mutate with (mutationChance/1000)% chance. Should it be more? Should it be less?
	// 0,01
	if r1.Intn(10000) < mutationChance {
		log.Println("!!! MUTATION !!!")
		posMutation := r1.Intn(len(gene.Actions))

		gene.Actions[posMutation] = genome.Action(r1.Intn(int(genome.Pickup + 1)))
	}

}

func (genes *MitchelGene) Mate(genesPartner genome.Genes) []genome.Genes {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	splitPosition := r1.Intn(len(genes.Actions) - 1)

	genesP := genesPartner.(*MitchelGene)

	genesChild1 := MitchelGene{
		Actions: make([]genome.Action, 243),
	}

	genesChild2 := MitchelGene{
		Actions: make([]genome.Action, 243),
	}

	for cont := 0; cont < splitPosition; cont++ {
		genesChild1.Actions[cont] = genes.Actions[cont]
		genesChild2.Actions[cont] = genesP.Actions[cont]
	}

	for cont := splitPosition; cont < len(genes.Actions); cont++ {
		genesChild1.Actions[cont] = genesP.Actions[cont]
		genesChild2.Actions[cont] = genes.Actions[cont]
	}

	genesChild1.mutate()
	genesChild2.mutate()

	return []genome.Genes{&genesChild1, &genesChild2}
}

func (genes *MitchelGene) GetActions() []genome.Action {
	return genes.Actions[0:200]
}
