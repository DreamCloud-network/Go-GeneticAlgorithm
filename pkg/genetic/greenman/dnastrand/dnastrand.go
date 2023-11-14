package dnastrand

import (
	"sort"
	"strings"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

type DNAStrand []genes.Gene

func NewDNAStrand() DNAStrand {
	return make([]genes.Gene, 0)
}

func (strand DNAStrand) String() string {
	var strStr strings.Builder

	for pos := 0; pos < len(strand); pos++ {
		strStr.WriteString(strand[pos].String())
		strStr.WriteString(" ")
	}

	return strStr.String()
}

// Add one or more genes to the chromosome.
func (strand *DNAStrand) AddGenes(genes []genes.Gene) {
	*strand = append(*strand, genes...)
}

// Add one or more genes to the chromosome.
func (strand *DNAStrand) AddGene(gene *genes.Gene) {
	*strand = append(*strand, *gene.Duplicate())
}

// Replace a gene in the chromosome.
func (strand *DNAStrand) ReplaceGene(pos int, gene *genes.Gene) {
	(*strand)[pos] = *gene.Duplicate()
}

// Returns a exact copy of the chromosome.
func (strand DNAStrand) Duplicate() DNAStrand {
	newStrand := make([]genes.Gene, len(strand))
	copy(newStrand, strand)
	return newStrand
}

// This function regulates the crossover between two chromosomes.
// It receives the chiasm positions of the two chromosomes and the quantity of genes in the chromosomes to assembly a final list of chiasmas
// that will be used to crossover the chromosomes.
// Based on the biological crossover, the chiasmas are distributed to balance the distance between the genes.
// It returns the chiasm positions of the new chromosomes.
func crossoverRegulation(c1ChiasmPositions, c2ChiasmPositions []int, genesQuantity int) []int {
	// Create the chiasm joint positions
	finalChiasmPositions := make([]int, len(c1ChiasmPositions))
	copy(finalChiasmPositions, c1ChiasmPositions)
	finalChiasmPositions = append(finalChiasmPositions, c2ChiasmPositions...)

	// Sort the chiasm positions
	sort.Ints(finalChiasmPositions)

	// Eliminate the chiasms that are equal or too close
	//chiasmQuantity := len(finalChiasmPositions)
	//genesPerChiasm := genesQuantity / chiasmQuantity

	// Eliminate the chiasms in the same position
	for pos := 0; pos < len(finalChiasmPositions); pos++ {
		if pos > 0 {
			if finalChiasmPositions[pos] == finalChiasmPositions[pos-1] {
				finalChiasmPositions = append(finalChiasmPositions[:pos], finalChiasmPositions[pos+1:]...)
				pos = 0
			}
		}
	}

	// Elimitate the chiasms in the same position or that are too close (less than 10% genes distance).
	// 25% is an arbitrary value that must be tested or changed in the future.
	/*
		eliminateGenes := genesPerChiasm / 4
		if eliminateGenes < 1 {
			eliminateGenes = 1
		}
		for pos := 0; pos < len(finalChiasmPositions); pos++ {
			if pos > 0 {
				if finalChiasmPositions[pos]-finalChiasmPositions[pos-1] <= eliminateGenes {
					finalChiasmPositions = append(finalChiasmPositions[:pos], finalChiasmPositions[pos+1:]...)
					pos = 0
				}
			}
		}
	*/

	// Reposition the chiasms that are closer than 50% of ganes count per chiasm to balance the distance between the genes
	// 50% is an arbitrary value that must be tested or changed in the future.
	/*
		genesPerChiasm = genesQuantity / len(finalChiasmPositions)
		repositionGenes := genesPerChiasm / 2
		if repositionGenes < 1 {
			repositionGenes = 1
		}

		for pos := 0; pos < len(finalChiasmPositions); pos++ {
			if pos > 0 {
				if finalChiasmPositions[pos]-finalChiasmPositions[pos-1] <= repositionGenes {
					valueToBallance := genesPerChiasm - (finalChiasmPositions[pos] - finalChiasmPositions[pos-1])
					for pos2 := pos; pos2 < len(finalChiasmPositions); pos2++ {
						finalChiasmPositions[pos2] += valueToBallance
					}

					pos = 0
				}
			}
		}
	*/
	// Eliminate any chiasm that is out of the chromosome
	for pos := 0; pos < len(finalChiasmPositions); pos++ {
		if finalChiasmPositions[pos] >= genesQuantity {
			finalChiasmPositions = finalChiasmPositions[:pos]
			break
		}
	}

	return finalChiasmPositions
}

// Return the positions of the genes that have chiasm.
func (strand DNAStrand) findChiasmPositions() []int {
	chiasmPositions := make([]int, 0)

	for pos := 0; pos < len(strand); pos++ {
		if strand[pos].HasChiasm() {
			chiasmPositions = append(chiasmPositions, pos)
		}
	}

	return chiasmPositions
}

/*
// Returns a new chromosome crossedover with another chromosome.
func (motherStrand *DNAStrand) Crossover(fatherStrand *DNAStrand) error {

	strandsLenght := len(*motherStrand)
	// Verify if the chromosomes have the same gene count
	if strandsLenght != len(*fatherStrand) {
		log.Println("dnastrand.Chromosome.Crossover - The chromosomes have different gene counts.")
		return ErrDifferentGeneCount
	}

	// Find the chiasm positions for the mother and father chromosomes
	motherChromosomeChiasmPositions := motherStrand.findChiasmPositions()
	fatherChromosomeChiasmPositions := fatherStrand.findChiasmPositions()

	finalCrossPositions := crossoverRegulation(motherChromosomeChiasmPositions, fatherChromosomeChiasmPositions, strandsLenght)

	//log.Println("Final chiasm positions:", finalCrossPositions)

	newMotherDNAStrand := NewDNAStrand()
	newFatherDNAStrand := NewDNAStrand()

	fatherRead := *fatherStrand
	motherRead := *motherStrand

	actualGene := 0
	nextCrossPosition := 0
	//addChiasmFlag := false

	for actualGene < strandsLenght {

		if (nextCrossPosition < len(finalCrossPositions)) && (actualGene == finalCrossPositions[nextCrossPosition]) {
			// Swap the chromosomes origins
			helperGenePointer := motherRead
			motherRead = fatherRead
			fatherRead = helperGenePointer

			nextCrossPosition++

			//addChiasmFlag = true

		}

		newMotherDNAStrand = append(newMotherDNAStrand, motherRead[actualGene])
		newFatherDNAStrand = append(newFatherDNAStrand, fatherRead[actualGene])

		/*
			if addChiasmFlag {
				newMotherDNAStrand[actualGene].AddChiasm()
				newFatherDNAStrand[actualGene].AddChiasm()
				addChiasmFlag = false
			} else {
				newMotherDNAStrand[actualGene].RemoveChiasm()
				newFatherDNAStrand[actualGene].RemoveChiasm()
			}
*/ /*
		newMotherDNAStrand[actualGene].RemoveChiasm()
		newFatherDNAStrand[actualGene].RemoveChiasm()
		actualGene++
	}

	*motherStrand = newMotherDNAStrand
	*fatherStrand = newFatherDNAStrand

	return nil
}
*/
