package dna

type DNA struct {
	Genes []Gene `json:"genes"`
}

func NewDNA() DNA {
	return DNA{
		Genes: make([]Gene, 0),
	}
}

func (dna *DNA) Duplicate() *DNA {
	newDNA := DNA{
		Genes: make([]Gene, len(dna.Genes)),
	}

	for i := range dna.Genes {
		newDNA.Genes[i] = *dna.Genes[i].Duplicate()
	}

	return &newDNA
}

func (dna *DNA) String() string {
	str := ""
	for num := range dna.Genes {
		str += "\n\r" + dna.Genes[num].String()
	}

	return str
}
