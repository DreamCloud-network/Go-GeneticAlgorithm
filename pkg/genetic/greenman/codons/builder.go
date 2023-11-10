package codons

import (
	"log"
	"math/rand"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
)

// Helps to build a new codon by adding one feda at a time.
// Also helps to check if the codon is complete and use only the values that corresponds to the fedas.
type Builder struct {
	codon          Codon
	pos            int
	fedaCodon      bool
	forfedaCodon   bool
	reverseBuilder bool
}

func NewBuilder() Builder {
	return Builder{
		codon:          GetEmptyCodon(),
		pos:            0,
		fedaCodon:      false,
		forfedaCodon:   false,
		reverseBuilder: false,
	}
}

// Configure the builder to build the codon in reverse order.
// The default is false.
func (b *Builder) SetReverseBuilder(reverseBuilder bool) {
	b.reverseBuilder = reverseBuilder
}

// Add a new valid fid to the codon if it is not full.
func (b *Builder) AddFid(fid feda.Fid) error {
	if !fid.IsFid() && (fid != feda.SPACE) {
		return ErrorInvalidFid
	}

	if (b.pos > 0) && (!b.fedaCodon) {
		return ErrorInvalidAction
	}

	if b.pos >= 3 {
		return ErrorCodonFull
	}

	if b.reverseBuilder {
		b.codon[2-b.pos] = fid
	} else {
		b.codon[b.pos] = fid
	}
	b.pos++

	b.fedaCodon = true

	return nil
}

// Add a new valid forfid to the codon if it is not full.
func (b *Builder) AddForfid(forfid feda.Fid) error {
	if !forfid.IsForfid() {
		return ErrorInvalidFid
	}

	if (b.pos > 0) && (!b.forfedaCodon) {
		return ErrorInvalidAction
	}

	if b.pos >= 3 {
		return ErrorCodonFull
	}

	if b.reverseBuilder {
		b.codon[2-b.pos] = forfid
	} else {
		b.codon[b.pos] = forfid
	}
	b.pos++

	b.forfedaCodon = true

	return nil
}

// Returns true codon if it is complete, otherwise returns false.
func (b *Builder) IsComplete() bool {
	return b.pos >= 3
}

// Returns true if the codon has at least one fid, otherwise returns false.
func (b *Builder) HasFid() bool {
	return b.pos > 0
}

// Returns the codon if it is complete, otherwise returns nil.
func (b *Builder) GetCodon() Codon {
	if b.IsComplete() {
		return b.codon
	}

	return GetEmptyCodon()
}

// Create a random codon.
func (b *Builder) GenerateRandomFedaCodon() int {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	maxFedaNum := int(feda.Idad) - int(feda.Beith) + 1

	for i := 0; i < 3; i++ {
		feda := feda.Fid(r1.Intn(maxFedaNum) + int(feda.Beith))
		err := b.AddFid(feda)
		if err != nil {
			// This error should never happens.
			log.Println("codons.Builder.AddRandomCodon - Error adding feda to codon.")
			panic(err)
		}
	}
	return b.pos
}

// Returns the codon that represents the begining of a gene.
func GetInitiatiorCodon() Codon {
	return INIT_CODON
}

// Returns the codon that represents the end of a gene.
/*
func GetEndCodon() Codon {
	return END_CODON
}
*/
// Returns the codon that represents a chiasm.
func GetChiasmCodon() Codon {
	return CHIASM_CODON
}

// Returns the codon that represents a empty codon.
func GetEmptyCodon() Codon {
	return EMPTY_CODON
}
