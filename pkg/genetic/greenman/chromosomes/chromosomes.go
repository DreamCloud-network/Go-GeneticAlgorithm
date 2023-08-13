package chromosomes

import (
	"log"
	"sort"
	"strings"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

type Chromosome struct {
	Strand []fedas.Feda
}

func NewChromosome() Chromosome {
	return Chromosome{
		Strand: make([]fedas.Feda, 0),
	}
}

func (chromosome *Chromosome) String() string {
	var cStr strings.Builder

	for _, feda := range chromosome.Strand {
		cStr.WriteString(feda.String())
	}

	return cStr.String()
}

// Add the fedas of a gene to the chromosome.
func (c *Chromosome) AddGene(gene *genes.Gene) {
	geneFedas := gene.GetFedas()

	c.Strand = append(c.Strand, geneFedas...)
}

// Add the fedas of a gene to the chromosome.
func (c *Chromosome) AddCodon(codon codons.Codon) {
	codonFedas := codon.GetFedas()

	c.Strand = append(c.Strand, codonFedas...)
}

// Return the code of the gene in a specific position, initiating from 0 (first position is 0).
// If a position out of range is passed, the function returns nil.
func (chromosome *Chromosome) ReadGeneInPosition(position int) (*genes.Gene, error) {
	actualPosition := 0

	geneDecoder := genes.NewDecoder()

	for _, feda := range chromosome.Strand {
		gene, err := geneDecoder.ReceiveFeda(feda)
		if err != nil {
			log.Println("chromosomes.Chromosome.ReadGeneInPosition - Error reading feda.")
			return nil, err
		}

		if gene != nil {
			if actualPosition == position {
				return gene, nil
			} else {
				actualPosition++
			}
		}
	}

	return nil, nil
}

// Returns a exact copy of the chromosome.
func (chromosome *Chromosome) Duplicate() *Chromosome {
	newChromosome := Chromosome{
		Strand: make([]fedas.Feda, len(chromosome.Strand)),
	}

	copy(newChromosome.Strand, chromosome.Strand)

	return &newChromosome
}

// Tokenize the chromossome strand in tokens composed by a sequence of codons and genes.
// Each gene end defines a new token.
// The Chiasm codon is not included in the tokens, but the function returns the token index where the chiasm was found.
func (chromosome *Chromosome) Tokenize() ([][]fedas.Feda, []int, error) {

	tokenizedChromosome := make([][]fedas.Feda, 1)
	chiasmPositions := make([]int, 0)

	geneDecoder := genes.NewDecoder()
	codonDecoder := codons.NewDecoder()

	tokenCount := 0

	for _, feda := range chromosome.Strand {

		if tokenCount == len(tokenizedChromosome) {
			tokenizedChromosome = append(tokenizedChromosome, make([]fedas.Feda, 0))
			geneDecoder.Reset()
		}

		err := codonDecoder.NewFeda(feda)
		if err != nil {
			log.Println("chromosomes.Chromosome.Tokenize - Error reading feda in the codon decoder.")
			return nil, nil, err
		}
		gene, err := geneDecoder.ReceiveFeda(feda)
		if err != nil {
			log.Println("chromosomes.Chromosome.Tokenize - Error reading feda in the gene decoder.")
			return nil, nil, err
		}

		if codonDecoder.CHIASM_CODON() {
			chiasmPositions = append(chiasmPositions, tokenCount)
		} else {
			if feda.IsFeda() || feda.IsSignalization() {
				tokenizedChromosome[tokenCount] = append(tokenizedChromosome[tokenCount], feda)
			}
		}

		if gene != nil {
			tokenCount++
		}
	}

	return tokenizedChromosome, chiasmPositions, nil
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
	sort.SliceStable(finalChiasmPositions, func(i, j int) bool {
		return finalChiasmPositions[i] < finalChiasmPositions[j]
	})

	// Eliminate the chiasms that are equal or too close
	chiasmQuantity := len(finalChiasmPositions)
	genesPerChiasm := genesQuantity / chiasmQuantity

	// Elimitate the chiasms in the same position or that are too close (less than 10% genes distance).
	// 25% is an arbitrary value that must be tested or changed in the future.
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

	// Reposition the chiasms that are closer than 50% of ganes count per chiasm to balance the distance between the genes
	// 50% is an arbitrary value that must be tested or changed in the future.
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

	// Eliminate any chiasm that is out of the chromosome
	for pos := 0; pos < len(finalChiasmPositions); pos++ {
		if finalChiasmPositions[pos] >= genesQuantity {
			finalChiasmPositions = finalChiasmPositions[:pos]
			break
		}
	}

	return finalChiasmPositions
}

// Returns a new chromosome crossedver with another chromosome.
func (motherChromosome *Chromosome) Crossover(fatherChromosome *Chromosome) error {
	motherChromosomeTokenized, motherChromosomeChiasmPositions, err := motherChromosome.Tokenize()
	if err != nil {
		log.Println("chromosomes.Chromosome.Crossover - Error tokenizing the chromosome.")
		return err
	}

	fatherChromosomeTokenized, fatherChromosomeChiasmPositions, err := fatherChromosome.Tokenize()
	if err != nil {
		log.Println("chromosomes.Chromosome.Crossover - Error tokenizing the other chromosome.")
		return err
	}

	// Verify if the chromosomes have the same gene count
	if len(motherChromosomeTokenized) != len(fatherChromosomeTokenized) {
		log.Println("chromosomes.Chromosome.Crossover - The chromosomes have different gene counts.")
		return ErrDifferentGeneCount
	}

	finalCrossPositions := crossoverRegulation(motherChromosomeChiasmPositions, fatherChromosomeChiasmPositions, len(motherChromosomeTokenized))

	newMotherChormosome := NewChromosome()
	newFatherChormosome := NewChromosome()

	strandSwitch := 0
	actualGene := 0

	for actualGene < len(motherChromosomeTokenized) {
		if strandSwitch == 0 {
			newMotherChormosome.Strand = append(newMotherChormosome.Strand, motherChromosomeTokenized[actualGene]...)
			newFatherChormosome.Strand = append(newFatherChormosome.Strand, fatherChromosomeTokenized[actualGene]...)

			for _, positions := range finalCrossPositions {
				if positions == actualGene {
					newMotherChormosome.AddCodon(codons.CHIASM_CODON)
					newFatherChormosome.AddCodon(codons.CHIASM_CODON)

					strandSwitch = 1
					break
				}
			}
		} else {
			newMotherChormosome.Strand = append(newMotherChormosome.Strand, fatherChromosomeTokenized[actualGene]...)
			newFatherChormosome.Strand = append(newFatherChormosome.Strand, motherChromosomeTokenized[actualGene]...)

			for _, positions := range finalCrossPositions {
				if positions == actualGene {
					newMotherChormosome.AddCodon(codons.CHIASM_CODON)
					newFatherChormosome.AddCodon(codons.CHIASM_CODON)

					strandSwitch = 0
					break
				}
			}
		}

		actualGene++
	}

	motherChromosome.Strand = newMotherChormosome.Strand
	fatherChromosome.Strand = newFatherChormosome.Strand

	return nil
}

/*
func NewChromosome() Chromosome {
	return Chromosome{
		Genes: make([]Gene, 0),
	}
}

func (chromosome *Chromosome) Duplicate() *Chromosome {
	newDNA := Chromosome{
		Genes: make([]Gene, len(chromosome.Genes)),
	}

	for i := range chromosome.Genes {
		newDNA.Genes[i] = *chromosome.Genes[i].Duplicate()
	}

	return &newDNA
}

func (chromosome *Chromosome) String() string {
	var dnaStr strings.Builder

	//dnaStr.WriteString("᚛  ")
	for num := range chromosome.Genes {
		dnaStr.WriteString(chromosome.Genes[num].String())
	}

	//dnaStr.WriteString("  ")

	return dnaStr.String()
}

// Execute the crossover of two DNA strands.
func Crossover(crossoverPoints int, chromosomeFather, chromosomeMother *Chromosome) error {
	if (crossoverPoints < 0) || (crossoverPoints > len(chromosomeFather.Genes)) {
		return ErrOutOfRange
	}
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate the crossover points.
	crossoverPositions := make([]int, crossoverPoints)

	// Initialize the positions with -1.
	for i := range crossoverPositions {
		crossoverPositions[i] = -1
	}

	for i := range crossoverPositions {
		// Garantee that the position is not repeated.
		newPosition := r1.Intn(len(chromosomeFather.Genes))
		for cont := 0; cont < len(crossoverPositions); cont++ {
			if crossoverPositions[cont] == newPosition {
				newPosition = r1.Intn(len(chromosomeFather.Genes))
				cont = 0
			}
		}

		crossoverPositions[i] = newPosition

		//Exchanges the genes between the two strands.
		chromosomeFather.Genes[newPosition], chromosomeMother.Genes[newPosition] = chromosomeMother.Genes[newPosition], chromosomeFather.Genes[newPosition]
	}

	return nil
}

type HomologousChromosomes struct {
	Father *Chromosome
	Mother *Chromosome
}

func NewHomologousChromosomes(father, mother *Chromosome) HomologousChromosomes {
	return HomologousChromosomes{
		Father: father,
		Mother: mother,
	}
}

// Duplicate the chromosomes.
func (chromosomes HomologousChromosomes) Duplicate() (HomologousChromosomes, error) {
	var homologousChromosomes HomologousChromosomes

	// Duplicate the chromosomes.
	fatherChromossomeReplica := homologousChromosomes.Father.Duplicate()
	motherChromossomeReplica := homologousChromosomes.Mother.Duplicate()

	return NewHomologousChromosomes(fatherChromossomeReplica, motherChromossomeReplica), nil
}

// Duplicate and execute de crossover to generate the spermatids dna material.
func (chromosomes HomologousChromosomes) GenerateSpermatidsGenes(crossOverPoints int) ([4]Chromosome, error) {
	fatherReplica := chromosomes.Father.Duplicate()
	motherReplica := chromosomes.Mother.Duplicate()

	// Execute the crossover.
	err := Crossover(crossOverPoints, fatherReplica, motherReplica)
	if err != nil {
		log.Println("greenman.GenerateGametes - Error generating crossover.")
		return [4]Chromosome{}, err
	}

	// Generate the gametes.
	return [4]Chromosome{
		*chromosomes.Father,
		*fatherReplica,
		*chromosomes.Mother,
		*motherReplica,
	}, nil
}

// Duplicate and execute de crossover to generate the ootid dna material.
func (chromosomes HomologousChromosomes) GenerateOotidGenes(crossOverPoints int) (Chromosome, error) {
	dnaMaterial, err := chromosomes.GenerateSpermatidsGenes(crossOverPoints)
	if err != nil {
		log.Println("greenman.GenerateOotidGenes - Error generating spermatids.")
		return Chromosome{}, err
	}

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select a random spermatid.
	return dnaMaterial[r1.Intn(len(dnaMaterial))], nil
}

// Save the DNA to file. Replaces any file in the directory with the same name.
/*func (dna *DNA) SaveToFile(fileNamePath string) error {
	dna.Genes = append(dna.Genes, gene)
}
*/
