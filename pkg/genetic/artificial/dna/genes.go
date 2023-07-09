package dna

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Gene struct {
	ID        uuid.UUID
	Dominance Codon
	Code      []Codon
}

func NewGene(dominant bool) *Gene {
	newGene := Gene{
		ID:        uuid.New(),
		Dominance: Dominant,
		Code:      make([]Codon, 0),
	}

	if !dominant {
		newGene.Dominance = Recessive
	}

	return &newGene
}

func (gene *Gene) String() string {
	str := ""
	if gene.Dominance == Recessive {
		str += "a-"
	} else {
		str += "A-"
	}

	for _, codon := range gene.Code {
		str += strconv.Itoa(int(codon))
	}

	return str
}

func (gene *Gene) Duplicate() *Gene {
	newGene := Gene{
		ID:        gene.ID,
		Dominance: gene.Dominance,
		Code:      make([]Codon, len(gene.Code)),
	}

	copy(newGene.Code, gene.Code)

	return &newGene
}

func (gene *Gene) ReturnGeneExpressionFromAlleles(gene2 *Gene) *Gene {

	//log.Println("INI - EXPRESSING GENES")
	//defer log.Println("END - EXPRESSING GENES")

	if gene.Dominance == Dominant {
		if gene2.Dominance == Dominant {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			if r1.Intn(2) == 0 {
				return gene
			} else {
				return gene2
			}
		} else {
			return gene
		}
	} else {
		if gene2.Dominance == Dominant {
			return gene2
		} else {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)

			if r1.Intn(2) == 0 {
				return gene
			} else {
				return gene2
			}
		}
	}
}
