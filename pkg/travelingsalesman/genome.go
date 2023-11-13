package travelingsalesman

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/circuit"
)

const mutationChance = 0.00001 // 0.00001 /%
const crossoverNumGenes = 10

type Genome struct {
	nextPos dnastrand.DNAStrand
}

// Create a new genome with random values based in the number of the nodes.
func NewRandomGenome(circuit *circuit.Circuit) *Genome {
	newGenome := Genome{
		nextPos: dnastrand.NewDNAStrand(),
	}

	for i := 0; i < len(circuit.Nodes); i++ {

		nodeConnections := circuit.Nodes[i].GetConnections()

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		val := r1.Intn(len(nodeConnections))

		codons := codons.UintToCodons(uint(val))

		newGenesBuilder := genes.NewBuilder()

		for _, codon := range codons {
			newGenesBuilder.AddCode(codon)
		}

		newGene := newGenesBuilder.Finish()

		newGenome.nextPos.AddGene(newGene)
	}

	return &newGenome
}

// Get the next position based in the actual position.
func (genome *Genome) GetNextPosition(actualPosition int) (int, error) {
	if actualPosition < 0 || actualPosition >= len(genome.nextPos) {
		return -1, ErrInvalidPosition
	}
	nexPosGene := genome.nextPos[actualPosition]

	nexPosCodons := nexPosGene.GetRawCode()

	nexPosVal, err := codons.CodonsToUint(nexPosCodons)
	if err != nil {
		log.Println("travelingsalesman.Genome.GetNextPosition")
		return -1, err
	}

	return int(nexPosVal), nil
}

// Generate a string witht the solution.
func (genome *Genome) String() string {
	var genomeStrBuilder strings.Builder

	for pos, gene := range genome.nextPos {
		genomeStrBuilder.WriteString("\n\r|" + strconv.Itoa(pos) + "| --> |")
		code := gene.GetRawCode()
		destPos, err := codons.CodonsToUint(code)
		if err != nil {
			log.Println("travelingsalesman.Genome.String")
			return ""
		}
		genomeStrBuilder.WriteString(strconv.Itoa(int(destPos)) + "|")
	}

	return genomeStrBuilder.String()
}

// Insert one crossover point at each crossoverNumGenes genes
func (gene *Genome) inserCrossOverPoints() {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	numCrossovers := (len(gene.nextPos) / crossoverNumGenes)

	if numCrossovers == 0 {
		numCrossovers = 1
	}

	for i := 0; i < numCrossovers; i++ {
		relposCrossover := r1.Intn(crossoverNumGenes)
		genePointToCrossover := i*crossoverNumGenes + relposCrossover
		if genePointToCrossover < len(gene.nextPos) {
			gene.nextPos[genePointToCrossover].AddChiasm()
		}

	}
}

// Mate two genomes.
func (mother *Genome) Mate(father *Genome) (*Genome, error) {
	// Insert one crossover point at each 10 genes
	mother.inserCrossOverPoints()

	// Duplicate the genes
	motherDup := mother.nextPos.Duplicate()
	fatherDup := father.nextPos.Duplicate()

	motherDup.Crossover(&fatherDup)

	// Randomly select one child gene to generate a child
	newGenome := &Genome{
		nextPos: nil,
	}

	if rand.Intn(2) == 0 {
		newGenome.nextPos = motherDup
		newGenome.Mutate()
		return newGenome, nil
	} else {
		newGenome.nextPos = fatherDup
		newGenome.Mutate()
		return newGenome, nil
	}
}

// Mutate the genome.
// I have questions if the mutation should really be done this way.
// In life the mutations have great chance to generate dnas that can´t be alive and the mutations that could
// in fact generate a working dna is very very very low.
// There are mutations caused by environemnt factors the act in specific genes, so in this case it can be good,
// but the way it is being done, generating a random gene, appears to be too much artificial.
func (genome *Genome) Mutate() {
	for pos := range genome.nextPos {
		if rand.Float64() < mutationChance {
			val := rand.Intn(len(genome.nextPos))

			codons := codons.UintToCodons(uint(val))

			newGenesBuilder := genes.NewBuilder()

			for _, codon := range codons {
				newGenesBuilder.AddCode(codon)
			}

			newGene := newGenesBuilder.Finish()

			genome.nextPos[pos] = *newGene
		}
	}
}
