package vine

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
)

func (meristem *VineCell) newRootCell() {
	newRoot := NewVineCell()

	newRoot.SetPosition(meristem.place)

	duplicatedGenome := meristem.genome.Duplicate()
	newRoot.genome = &duplicatedGenome

	newRoot.genome.ActivateOnlyFunction(ROOT_CODON)

	// Send energy and nutrients to the new Root cell.
	newRoot.energy = float64(meristem.genome.GetRegulationvalue(ROOT_ENERGY))
	meristem.energy -= newRoot.energy

	newRoot.nutrients = float64(meristem.genome.GetRegulationvalue(ROOT_NUTRIENTS))
	meristem.nutrients -= newRoot.nutrients

	newRoot.connectXylem(meristem.xylem)

	newRoot.connectPhloem(meristem.phloem)

	newRoot.connectCells(meristem)

	//go newRoot.Run()
}

// Receive nutrients from the earth .
func (vc *VineCell) rootRun() {

	// Receive nutrients from the earth.
	thingsVector := vc.place.GetThingsType(place.Earth)

	if len(thingsVector) > 0 {
		// 5 is an arbitrary value that corresponds to the maximum amount of nutrients that one root cell can receive.
		vc.nutrients += 5
	}

	// If there is enought nutrients for survival and a little more, send nutients for the xylem.
	// I used 10 arbitrarily, but it can be changed or adjusted by the DNA.
	if vc.nutrients > 10 {
		vc.xylem.receiveNutrients(1)
		vc.nutrients--
	}
}
