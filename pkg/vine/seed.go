package vine

import "github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"

func NewVineSeed() *VineCell {
	newVineSeed := NewVineCell()

	newVineSeed.genome.ActivateFunction(SEED_CODON)

	newVineSeed.energy = float64(newVineSeed.genome.GetRegulationvalue(SEED_ENERGY))
	newVineSeed.nutrients = float64(newVineSeed.genome.GetRegulationvalue(SEED_NUTRIENTS))

	return newVineSeed
}

// Transform the seed in a meristem cell.
func (seed *VineCell) transformIntoMeristemCell() {
	seed.genome.DisableFunction(SEED_CODON)
	seed.genome.ActivateFunction(MERISTEM_CODON)
	seed.genome.ActivateFunction(XYLEM_CODON)
	seed.genome.ActivateFunction(PHLOEM_CODON)
	//seed.genome.ActivateFunction(LEAF_CODON)
	//seed.genome.ActivateFunction(ROOT_CODON)
}

func (seed *VineCell) seedReadyToHatch() bool {
	// Verify if the seed is in the earth.
	return seed.detectThing(place.Earth)
}

// Execute the seed functions.
func (seed *VineCell) seedRun() {
	// Verify if the seed is in the earth.
	if seed.seedReadyToHatch() {
		// Create a Meristem cell in the same place.
		seed.transformIntoMeristemCell()
	}
}
