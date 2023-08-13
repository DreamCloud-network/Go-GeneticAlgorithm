package codons

import (
	"log"
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
)

// Helps to build a new codon by adding one feda at a time.
// Also helps to check if the codon is complete and use only the values that corresponds to the fedas.
type Builder struct {
	codons *Codon
	pos    int
}

func NewBuilder() Builder {
	newCodon := EMPTY_CODON
	return Builder{
		codons: &newCodon,
		pos:    0,
	}
}

// Add a new valid feda to the codon if it is not full.
func (b *Builder) AddFeda(feda fedas.Feda) error {
	if !feda.IsFeda() {
		//log.Println("codons.Builder.AddFeda - Error adding feda to codon.")
		return ErrorInvalidFeda
	}

	if b.pos >= 3 {
		return ErrorCodonFull
	}

	b.codons[b.pos] = feda
	b.pos++

	return nil
}

// Returns true codon if it is complete, otherwise returns false.
func (b *Builder) IsComplete() bool {
	return b.pos >= 3
}

// Returns the codon if it is complete, otherwise returns nil.
func (b *Builder) GetCodon() *Codon {
	if b.IsComplete() {
		return b.codons
	}
	return nil
}

// Reset the builder to start a new codon.
func (b *Builder) Reset() {
	newCodon := EMPTY_CODON
	b.codons = &newCodon
	b.pos = 0
}

// Create a random codon.
func (b *Builder) GenerateRandomCodon() int {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	maxFedaNum := int(fedas.Idad) - int(fedas.Beith)

	for i := 0; i < 3; i++ {
		feda := fedas.Feda(r1.Intn(maxFedaNum) + int(fedas.Beith))
		err := b.AddFeda(feda)
		if err != nil {
			// This error should never happens.
			log.Println("codons.Builder.AddRandomCodon - Error adding feda to codon.")
			panic(err)
		}
	}
	return b.pos
}
