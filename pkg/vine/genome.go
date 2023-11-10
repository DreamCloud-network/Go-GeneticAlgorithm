package vine

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

// Function codons
var SEED_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Beith, feda.Beith}
var MERISTEM_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Beith, feda.Luis}
var XYLEM_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Beith, feda.Fearn}
var PHLOEM_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Beith, feda.Saille}
var STALK_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Beith, feda.Nuin}
var LEAF_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Luis, feda.Beith}
var ROOT_CODON codons.Codon = [3]feda.Fid{feda.Gort, feda.Luis, feda.Luis}

type genesFunctions int

const (
	SEED genesFunctions = iota
	MERISTEM
	XYLEM
	PHLOEM
	STALK
	LEAF
	ROOT
)

type genesRegulations int

const (
	SEED_ENERGY genesRegulations = iota
	SEED_NUTRIENTS
	MERISTEM_ENERGY
	MERISTEM_NUTRIENTS
	XYLEM_ENERGY
	XYLEM_NUTRIENTS
	PHLOEM_ENERGY
	PHLOEM_NUTRIENTS
	LEAF_ENERGY
	LEAF_NUTRIENTS
	ROOT_ENERGY
	ROOT_NUTRIENTS
)

type Genome struct {
	functions   dnastrand.DNAStrand
	regulations dnastrand.DNAStrand
}

func NewGenome() *Genome {
	newGenome := &Genome{
		functions:   dnastrand.NewDNAStrand(),
		regulations: dnastrand.NewDNAStrand(),
	}

	newGenome.newFunctionGene()

	newGenome.newRegulationGene()

	return newGenome
}

// Generate a new basic function gene.
func (newGene *Genome) newFunctionGene() {

	//SEED_CODON
	genesBuilder := genes.NewBuilder()
	genesBuilder.AddCode(SEED_CODON)
	seedGene := genesBuilder.Finish()
	seedGene.Disable()
	newGene.functions.AddGene(*seedGene)

	//MERISTEM_CODON
	genesBuilder = genes.NewBuilder()
	genesBuilder.AddCode(MERISTEM_CODON)
	meristemGene := genesBuilder.Finish()
	meristemGene.Disable()
	newGene.functions.AddGene(*meristemGene)

	//XYLEM_CODON
	genesBuilder = genes.NewBuilder()
	genesBuilder.AddCode(XYLEM_CODON)
	xylemGene := genesBuilder.Finish()
	xylemGene.Disable()
	newGene.functions.AddGene(*xylemGene)

	//PHLOEM_CODON
	genesBuilder = genes.NewBuilder()
	genesBuilder.AddCode(PHLOEM_CODON)
	phloemGene := genesBuilder.Finish()
	phloemGene.Disable()
	newGene.functions.AddGene(*phloemGene)

	//STALK_CODON
	genesBuilder = genes.NewBuilder()
	genesBuilder.AddCode(STALK_CODON)
	stalkGene := genesBuilder.Finish()
	stalkGene.Disable()
	newGene.functions.AddGene(*stalkGene)

	//LEAF_CODON
	genesBuilder = genes.NewBuilder()
	genesBuilder.AddCode(LEAF_CODON)
	leafGene := genesBuilder.Finish()
	leafGene.Disable()
	newGene.functions.AddGene(*leafGene)

	//ROOT_CODON
	genesBuilder = genes.NewBuilder()
	genesBuilder.AddCode(ROOT_CODON)
	rootGene := genesBuilder.Finish()
	rootGene.Disable()
	newGene.functions.AddGene(*rootGene)

}

// GEnerate a new basic regulation gene.
func (newGene *Genome) newRegulationGene() {
	// SEED_ENERGY
	energyCodons := codons.UintToCodons(uint(1000))
	genesBuilder := genes.NewBuilder()
	for _, codon := range energyCodons {
		genesBuilder.AddCode(codon)
	}
	energyGene := genesBuilder.Finish()
	newGene.regulations.AddGene(*energyGene)

	// SEED_NUTRIENTS
	nutrientsCodons := codons.UintToCodons(uint(1000))
	genesBuilder = genes.NewBuilder()
	for _, codon := range nutrientsCodons {
		genesBuilder.AddCode(codon)
	}
	nutrientGenes := genesBuilder.Finish()
	newGene.regulations.AddGene(*nutrientGenes)

	// MERISTEM_ENERGY
	energyCodons = codons.UintToCodons(uint(100))
	genesBuilder = genes.NewBuilder()
	for _, codon := range energyCodons {
		genesBuilder.AddCode(codon)
	}
	energyGene = genesBuilder.Finish()
	newGene.regulations.AddGene(*energyGene)

	// MERISTEM_NUTRIENTS
	nutrientsCodons = codons.UintToCodons(uint(100))
	genesBuilder = genes.NewBuilder()
	for _, codon := range nutrientsCodons {
		genesBuilder.AddCode(codon)
	}
	nutrientGenes = genesBuilder.Finish()
	newGene.regulations.AddGene(*nutrientGenes)

	// Starts with 10 energy and nutrients for all other kind of cells.
	for cont := XYLEM_ENERGY; cont <= ROOT_NUTRIENTS; cont++ {
		valueCodons := codons.UintToCodons(uint(10))
		genesBuilder = genes.NewBuilder()
		for _, codon := range valueCodons {
			genesBuilder.AddCode(codon)
		}
		valueGenes := genesBuilder.Finish()
		newGene.regulations.AddGene(*valueGenes)
	}
}

func (gen Genome) Duplicate() Genome {
	return Genome{
		functions:   gen.functions.Duplicate(),
		regulations: gen.regulations.Duplicate(),
	}
}

// Return all the active functions in the DNA strand.
func (gen *Genome) GetActiveFunctions() []genes.Gene {
	activeFunctions := make([]genes.Gene, 0)

	for _, gene := range gen.functions {
		if gene.IsEnabled() {
			activeFunctions = append(activeFunctions, gene)
		}
	}

	return activeFunctions
}

func (gen *Genome) ActivateFunction(function codons.Codon) {
	for cont := range gen.functions {
		if gen.functions[cont].GetRawCode()[0] == function {
			gen.functions[cont].Enable()
		}
	}
}

func (gen *Genome) ActivateOnlyFunction(function codons.Codon) {
	for cont := range gen.functions {
		if gen.functions[cont].GetRawCode()[0] == function {
			gen.functions[cont].Enable()
		} else {
			gen.functions[cont].Disable()
		}
	}
}

func (gen *Genome) DisableFunction(function codons.Codon) {
	for cont := range gen.functions {
		if gen.functions[cont].GetRawCode()[0] == function {
			gen.functions[cont].Disable()
		}
	}
}

func (gen *Genome) GetRegulationvalue(function genesRegulations) uint {
	if (function > ROOT_NUTRIENTS) || (function < SEED_ENERGY) {
		return 0
	}

	regulationCodons := gen.regulations[function].GetRawCode()
	energy, err := codons.CodonsToUint(regulationCodons)
	if err != nil {
		log.Println("vine.Genome.getRegulationvalue - Error reading energy to store in seeds from genome.")
		return 0
	}

	return energy
}

/*
func (gen *Genome) ActivateLeafFunctions() {

}

func (gen *Genome) ActivateStalkFunctions() {
	// Configure function gene
	genesBuilder := genes.NewBuilder()
	genesBuilder.AddCode(STALK_CODON)
	stalkFunctionGene := genesBuilder.Finish()

	gen.regulations[FUNCTION] = *stalkFunctionGene

	// Configure movement gene
	/*genesBuilder = genes.NewBuilder()
	connectionNumberCodons := codons.UintToCodons(uint(connectionNumber))
	for _, codon := range connectionNumberCodons {
		genesBuilder.AddCode(codon)
	}

	connectionGene := genesBuilder.Finish()

	gen.regulations[MOVEMENT] = *connectionGene*/
//}
/*
func (gen *Genome) ReadConnectionNumber() (int, error) {
	if gen.regulations[MOVEMENT].IsDisabled() {
		return -1, nil
	}

	connectionNumber := gen.regulations[MOVEMENT].GetRawCode()
	number, err := codons.CodonsToUint(connectionNumber)
	if err != nil {
		log.Println("vine.Genome.ReadConnectionNumber - Error reading connection number from genome.")
		return -1, err
	}

	return int(number), nil
}

func (gen *Genome) DisableMovementGene() {
	gen.regulations[MOVEMENT].Disable()
}

// EnableReproductionGene - Enables the reproduction gene.
func (gen *Genome) EnableReproductionGene() {
	gen.regulations[REPRODUCTION].Enable()
}

func (gen *Genome) DisableReproductionGene() {
	gen.regulations[REPRODUCTION].Disable()
}

func (gen *Genome) ReproductionGeneEnabled() bool {
	return !gen.regulations[REPRODUCTION].IsDisabled()
}
*/
