package tobby

import (
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

type Action codons.Codon

var RandomMove Action = [3]feda.Fid{feda.Beith, feda.Beith, feda.Beith}
var StepNorth Action = [3]feda.Fid{feda.Beith, feda.Beith, feda.Luis}
var StepWest Action = [3]feda.Fid{feda.Beith, feda.Beith, feda.Fearn}
var StepSouth Action = [3]feda.Fid{feda.Beith, feda.Beith, feda.Saille}
var StepEast Action = [3]feda.Fid{feda.Beith, feda.Beith, feda.Nuin}
var Pickup Action = [3]feda.Fid{feda.Uath, feda.Uath, feda.Uath}

var DoNothing Action = [3]feda.Fid{feda.Peith, feda.Peith, feda.Peith}

func (m Action) String() string {
	switch m {
	case StepNorth:
		return "StepNorth"
	case StepSouth:
		return "StepSouth"
	case StepEast:
		return "StepEast"
	case StepWest:
		return "StepWest"
	case RandomMove:
		return "RandomMove"
	case Pickup:
		return "Pickup"

	default:
		err := ErrorInvalidAction
		return err.Error()
	}
}

func (action Action) IsValid() bool {
	switch action {
	case StepNorth:
		return true
	case StepSouth:
		return true
	case StepEast:
		return true
	case StepWest:
		return true
	case RandomMove:
		return true
	case Pickup:
		return true

	default:
		return false
	}
}

func GetRandomAction() Action {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch r1.Intn(6) {
	case 0:
		return StepNorth
	case 1:
		return StepSouth
	case 2:
		return StepEast
	case 3:
		return StepWest
	case 4:
		return RandomMove
	case 5:
		return Pickup
	}

	return Pickup
}

func GetRandomMove() Action {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch r1.Intn(4) {
	case 0:
		return StepNorth
	case 1:
		return StepSouth
	case 2:
		return StepEast
	case 3:
		return StepWest
	}

	return StepNorth
}

func NewRandomActionGene() *genes.Gene {
	newGeneBuilder := genes.NewBuilder()

	newGeneBuilder.AddCode(codons.Codon(GetRandomAction()))

	newGene := newGeneBuilder.Finish()

	// I used 5% of chance to generate a crossover in this gene
	// TO DO: think better about this value
	//r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	//if r1.Intn(100) < 5 {
	newGene.AddChiasm()
	//}

	return newGeneBuilder.Finish()
}
