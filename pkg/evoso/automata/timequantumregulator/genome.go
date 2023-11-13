package timequantumregulator

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

var TimeQuantumRegulatorId_Codon codons.Codon = [3]feda.Fid{feda.Onn, feda.Muin, feda.Idad}
var TimeQuantumProducer_Codon codons.Codon = [3]feda.Fid{feda.Nuin, feda.Idad, feda.Onn}

func NewBasicTimeQuantumRegulatorGenome() dnastrand.DNAStrand {
	newDna := dnastrand.NewDNAStrand()

	// Create and add the id gene
	idGene := genes.NewGene()
	idGene.AddCodon(TimeQuantumRegulatorId_Codon)
	idGene.Enable()

	newDna.AddGene(idGene)

	// Create and add the quantum producer gene
	timeQuantumProducerGene := genes.NewGene()
	timeQuantumProducerGene.AddCodon(TimeQuantumRegulatorId_Codon)
	timeQuantumProducerGene.AddCodon(TimeQuantumProducer_Codon)
	timeQuantumProducerGene.AddCodon(codons.EMPTY_CODON)
	codonsSleepTime := codons.UintToCodons(100)
	for _, codon := range codonsSleepTime {
		timeQuantumProducerGene.AddCodon(codon)
	}
	timeQuantumProducerGene.Enable()

	newDna.AddGene(timeQuantumProducerGene)

	return newDna
}
