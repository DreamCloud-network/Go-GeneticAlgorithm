package tobby

import (
	"log"
	"math/rand"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

const mutationChance = 0.00001 //0.000001 // /%
const actionStrandLength = 243

type TobbyGene struct {
	Actions dnastrand.DNAStrand
}

func NewEmptyGenes() *TobbyGene {

	newGenes := TobbyGene{
		Actions: dnastrand.NewDNAStrand(),
	}

	return &newGenes
}

func NewRandomGenes() *TobbyGene {

	newGenes := TobbyGene{
		Actions: dnastrand.NewDNAStrand(),
	}

	for cont := 0; cont < actionStrandLength; cont++ {
		newGenes.Actions = append(newGenes.Actions, *NewRandomActionGene())
	}

	return &newGenes
}

// Translate the gene code into an actions.
// Return error if the code doesn´t correspondt to any action
func translateGeneIntoAction(actionGene genes.Gene) (Action, error) {

	// All the actions have only three bases
	// First see if the gene size is ok - 1 + 2 (init and end)
	geneCode := actionGene.GetCode()
	if len(geneCode) != 3 {
		log.Println("tobby.TobbyGene.translateGeneIntoAction - Error translating gene into action.")
		return Action(codons.EMPTY_CODON), ErrorInvalidAction
	}

	// Read the codon that corresponds to the action.
	newAction := Action(geneCode[1])

	if newAction.IsValid() {
		return newAction, nil
	}

	return Action(codons.EMPTY_CODON), ErrorInvalidAction
}

// String returns a string representation of the Genes
func (genes *TobbyGene) ActionString() string {
	var str string

	for geneCont, gene := range genes.Actions {

		action, err := translateGeneIntoAction(gene)
		if err != nil {
			log.Println("tobby.TobbyGene.ActionString - Error translating gene into action.")
			return ""
		}
		if geneCont == 0 {
			str = action.String()
		} else {
			str = str + "|" + action.String()
		}
	}
	return str
}

// Return the action based in the robot activations in the position (positionSignature).
// The function takes in consideration that the action with grater number is dominant over the other.
func (genes *TobbyGene) GetAction(positionSignature int) Action {
	if positionSignature >= actionStrandLength {
		return DoNothing
	}

	action, err := translateGeneIntoAction(genes.Actions[positionSignature])
	if err != nil {
		log.Println("tobby.TobbyGene.GetAction - Error translating gene into action.")
		return DoNothing
	}

	return action
}

// Return a new genome from mating a father and a mother genomes.
func (mother *TobbyGene) Mate(father *TobbyGene) *TobbyGene {
	newChild1 := NewEmptyGenes()
	newChild2 := NewEmptyGenes()

	newChild1.Actions = make([]genes.Gene, len(father.Actions))
	copy(newChild1.Actions, father.Actions)

	newChild2.Actions = make([]genes.Gene, len(mother.Actions))
	copy(newChild2.Actions, mother.Actions)

	newChild1.Actions.Crossover(&newChild2.Actions)

	// Randomly select one child gene to generate a child
	if rand.Intn(2) == 0 {
		newChild1.Mutate()
		return newChild1
	} else {
		newChild2.Mutate()
		return newChild2
	}
}

func (genes *TobbyGene) Mutate() {
	for cont := 0; cont < actionStrandLength; cont++ {
		if rand.Float64() < mutationChance {
			log.Println("!!! MUTATION !!!")
			genes.Actions[cont] = *NewRandomActionGene()
		}
	}
}
