package genes

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
)

// The decoder struct helps to recode genes from a sequence of fedas or codons.
type Decoder struct {
	geneBuilder  *Builder
	codonBuilder *codons.Builder
	codonDecoder codons.Decoder
	state        int
}

// Returns a new decoder already initiaded.
func NewDecoder() Decoder {
	return Decoder{
		geneBuilder:  nil,
		codonBuilder: nil,
		codonDecoder: codons.NewDecoder(),
		state:        0,
	}
}

// Returns true when the decoder is assembling a gene.
// It means that an initiator codon was found and the decoder is waiting for a end codon.
func (d *Decoder) AssemblingGene() bool {
	return d.state > 0
}

// Receive a feda and decode the sequence.
// Returns the gene when the sequence is complete, until there, returns nil.
func (d *Decoder) ReceiveFeda(feda fedas.Feda) (*Gene, error) {

	err := d.codonDecoder.NewFeda(feda)
	if err != nil {
		log.Println("genes.Decoder.ReceiveFeda - Error receiving feda.")
		return nil, err
	}

	switch d.state {
	case 0:
		found := d.codonDecoder.INIT_CODON()
		if found {
			newGeneBuilder := NewBuilder()
			d.geneBuilder = &newGeneBuilder

			newCodonBuilder := codons.NewBuilder()
			d.codonBuilder = &newCodonBuilder

			d.state++
		}
	case 1:
		found := d.codonDecoder.END_CODON()

		if found {
			gene := d.geneBuilder.GetGene()
			d.Reset()
			return gene, nil
		} else {
			err := d.codonBuilder.AddFeda(feda)
			if err != nil {
				// In this case, it was received an symbol that is not a feda.
				if (feda != fedas.SPACE) && (feda != fedas.TERMINATOR) {
					log.Println("genes.Decoder.ReceiveFeda - Error adding feda to codon.")
					return nil, err
				}
			}

			if d.codonBuilder.IsComplete() {
				newCodon := d.codonBuilder.GetCodon()
				d.geneBuilder.AddCodon(*newCodon)
				d.codonBuilder.Reset()
			}
		}
	}

	// It should be impossible to reach this return.
	return nil, nil
}

// Reset the decoder state.
func (d *Decoder) Reset() {
	d.codonDecoder = codons.NewDecoder()
	d.geneBuilder = nil
	d.codonBuilder = nil
	d.state = 0
}
