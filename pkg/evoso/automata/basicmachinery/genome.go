package basicmachinery

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
)

type basicMachinaryGenomePositions int

const (
	COLECTORS basicMachinaryGenomePositions = iota
)

var SEPARATOR_CODON codons.Codon = [3]feda.Fid{feda.SPACE, feda.Uath, feda.SPACE}

var COLECTOR_CODON codons.Codon = [3]feda.Fid{feda.Nuin, feda.Coll, feda.Gort}
var TIME_QUANTUM_REGULATOR_CODON codons.Codon = [3]feda.Fid{feda.Onn, feda.Muin, feda.Idad}

var TimeQuantumProducer_Codon codons.Codon = [3]feda.Fid{feda.Nuin, feda.Idad, feda.Onn}
