package tobby

import (
	"log"
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
)

const mutationChance = 0.000001 // /%
const actionStrandLength = 243
const lifeSpanStrandLength = 1
const numSpermatids = 100000000

type TobbyGene struct {
	Regulator greenman.HomologousChromosomes
	Actions   greenman.HomologousChromosomes
}

type TobbyGametes struct {
	Regulator *greenman.Chromosome
	Actions   *greenman.Chromosome
}

// Create and return a new Regulator chromosome with 1 point of crossover
func newRegulatorChromosome() *greenman.Chromosome {

	newGenes := greenman.NewChromosome()

	// CrossOverPoints Gene
	newGenes.Genes = append(newGenes.Genes, greenman.NewGene())
	newGenes.Genes[0].AppendFeda(fedas.Feda(1))

	return &newGenes
}

func newActionsChromosome() *greenman.Chromosome {

	newGenes := greenman.NewChromosome()

	// CrossOverPoints Gene
	newGenes.Genes = make([]greenman.Gene, actionStrandLength)
	for genePos := range newGenes.Genes {
		newGenes.Genes[genePos] = greenman.NewGene()
	}

	return &newGenes
}

func NewGenes() *TobbyGene {

	newGenes := TobbyGene{
		Regulator: greenman.NewHomologousChromosomes(newRegulatorChromosome(), newRegulatorChromosome()),
		Actions:   greenman.NewHomologousChromosomes(newActionsChromosome(), newActionsChromosome()),
	}

	return &newGenes
}

// Initialize the genome code with ramdom values for all chromosomes
func (genes *TobbyGene) RamdomizeGenes() {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Regulator - CrossOverPoints
	genes.Regulator.Father.Genes[0].ResetCode()
	genes.Regulator.Father.Genes[0].AppendFeda(fedas.Feda(r1.Intn(actionStrandLength-1) + 1))

	genes.Regulator.Mother.Genes[0].ResetCode()
	genes.Regulator.Mother.Genes[0].AppendFeda(fedas.Feda(r1.Intn(actionStrandLength-1) + 1))

	// Actions
	for genePos := range genes.Actions.Father.Genes {
		genes.Actions.Father.Genes[genePos].ResetCode()
		genes.Actions.Father.Genes[genePos].AppendFeda(fedas.Feda(r1.Intn(int(Pickup)) + 1))
	}

	for genePos := range genes.Actions.Mother.Genes {
		genes.Actions.Mother.Genes[genePos].ResetCode()
		genes.Actions.Mother.Genes[genePos].AppendFeda(fedas.Feda(r1.Intn(int(Pickup)) + 1))
	}
}

/*
// String returns a string representation of the Genes
func (genes *TobbyGene) ActionString() string {
	var str string

	for geneCont := range genes.Actions.Genes {
		code, err := genes.Actions.Genes[geneCont].ReadCode(1)
		if err != nil {
			code = fedas.Peith
		}

		if geneCont == 0 {
			str = Action(code).String()
		} else {
			newAction := Action(code).String()
			str = str + "|" + newAction
		}
	}
	return str
}
*/

// Return the action based in the robot activations in the position (positionSignature).
// The function takes in consideration that the action with grater number is dominant over the other.
func (genes *TobbyGene) GetAction(positionSignature int) Action {
	if positionSignature >= actionStrandLength {
		return 0
	}

	fatherGeneAction, err := genes.Actions.Father.Genes[positionSignature].ReadCode(1)
	if err != nil {
		log.Println("tobby.GetAction - Error reading father gene")
		return 0
	}

	motherGeneAction, err := genes.Actions.Mother.Genes[positionSignature].ReadCode(1)
	if err != nil {
		log.Println("tobby.GetAction - Error reading mother gene")
		return 0
	}

	if fatherGeneAction > motherGeneAction {
		return Action(fatherGeneAction)
	}

	return Action(motherGeneAction)
}

// Return the crossover point taking in account that the a crossover point value grater is dominant.
func (genes *TobbyGene) getCrossOverPoint() int {
	fatherCrossOverPoint, err := genes.Regulator.Father.Genes[0].ReadCode(1)
	if err != nil {
		log.Println("tobby.GetCrossOverPoint - Error reading father gene")
		return 0
	}

	motherCrossOverPoint, err := genes.Regulator.Mother.Genes[0].ReadCode(1)
	if err != nil {
		log.Println("tobby.GetCrossOverPoint - Error reading mother gene")
		return 0
	}

	if fatherCrossOverPoint > motherCrossOverPoint {
		return int(fatherCrossOverPoint)
	}

	return int(motherCrossOverPoint)
}

// Generate the Tobby spermatids with two chromosomes (Regulator and Actions) each.
func (genes *TobbyGene) GenerateSpermatidsGenes() []TobbyGametes {
	spermatids := make([]TobbyGametes, 0, numSpermatids)

	for len(spermatids) < numSpermatids {
		regulatorChromatids, err := genes.Regulator.GenerateSpermatidsGenes(genes.getCrossOverPoint())
		if err != nil {
			log.Println("tobby.Mate - Error generating spermatoids for regulator.")
			return nil
		}

		actionsChromatids, err := genes.Actions.GenerateSpermatidsGenes(genes.getCrossOverPoint())
		if err != nil {
			log.Println("tobby.Mate - Error generating spermatoids for actions.")
			return nil
		}

		for cont := range regulatorChromatids {
			newSpermatid := TobbyGametes{
				Regulator: &regulatorChromatids[cont],
				Actions:   &actionsChromatids[cont],
			}

			spermatids = append(spermatids, newSpermatid)
		}
	}

	return spermatids
}

// Generate a Tobby ootide with two chromosomes (Regulator and Actions) each.
func (genes *TobbyGene) GenerateOotidGenes() *TobbyGametes {
	regulatorOotid, err := genes.Regulator.GenerateOotidGenes(genes.getCrossOverPoint())
	if err != nil {
		log.Println("tobby.GenerateOotidGenes - Error generating ootid for regulator.")
		return nil
	}

	actionOotid, err := genes.Actions.GenerateOotidGenes(genes.getCrossOverPoint())
	if err != nil {
		log.Println("tobby.GenerateOotidGenes - Error generating ootid for actions.")
		return nil
	}

	return &TobbyGametes{
		Regulator: &regulatorOotid,
		Actions:   &actionOotid,
	}
}

// Return a new genome from mating a father and a mother genomes.
func (father *TobbyGene) Mate(mother *TobbyGene) *TobbyGene {

	spermatids := father.GenerateSpermatidsGenes()

	ootid := mother.GenerateOotidGenes()

	if spermatids == nil || ootid == nil {
		log.Println("tobby.Mate - Error generating spermatoids or ootid.")
		return nil
	}

	// Generate the new TobbyGene

	// Select one spermatoid
	// This is selected randomly, but in the future I should think of a way to select the "best" spermatoid
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	spermatid := spermatids[r1.Intn(len(spermatids))]

	// Generate new Regulator chromosome
	return &TobbyGene{
		Regulator: greenman.NewHomologousChromosomes(spermatid.Regulator, ootid.Regulator),
		Actions:   greenman.NewHomologousChromosomes(spermatid.Actions, ootid.Actions),
	}
}
