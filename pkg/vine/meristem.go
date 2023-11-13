package vine

import (
	"log"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
)

/*func (meristem *VineCell) newMeristemCell() *VineCell {
	newMeristem := NewVineCell()

	newMeristem.SetPosition(meristem.place)

	duplicatedGenome := meristem.genome.Duplicate()
	newMeristem.genome = &duplicatedGenome

	newMeristem.genome.ActivateFunction(MERISTEM_CODON)

	newMeristem.previusCell = meristem

	return newMeristem
}*/

func (meristem *VineCell) newMeristemRootCell() {
	newMeristemRoot := NewVineCell()

	newMeristemRoot.SetPosition(meristem.place)

	duplicatedGenome := meristem.genome.Duplicate()
	newMeristemRoot.genome = &duplicatedGenome

	newMeristemRoot.genome.ActivateOnlyFunction(MERISTEM_CODON)
	newMeristemRoot.genome.ActivateFunction(ROOT_CODON)

	// Send energy and nutrients to the new Root cell.
	newMeristemRoot.energy = float64(meristem.genome.GetRegulationvalue(MERISTEM_ENERGY))
	meristem.energy -= newMeristemRoot.energy

	newMeristemRoot.nutrients = float64(meristem.genome.GetRegulationvalue(MERISTEM_NUTRIENTS))
	meristem.nutrients -= newMeristemRoot.nutrients

	newMeristemRoot.connectXylem(meristem.xylem)

	newMeristemRoot.connectPhloem(meristem.phloem)

	newMeristemRoot.connectCells(meristem)

	//go newMeristemRoot.Run()
}

func (meristem *VineCell) newMeristemLeafCell() {
	newMeristemRoot := NewVineCell()

	newMeristemRoot.SetPosition(meristem.place)

	duplicatedGenome := meristem.genome.Duplicate()
	newMeristemRoot.genome = &duplicatedGenome

	newMeristemRoot.genome.ActivateOnlyFunction(MERISTEM_CODON)
	newMeristemRoot.genome.ActivateFunction(LEAF_CODON)

	// Send energy and nutrients to the new Root cell.
	newMeristemRoot.energy = float64(meristem.genome.GetRegulationvalue(MERISTEM_ENERGY))
	meristem.energy -= newMeristemRoot.energy

	newMeristemRoot.nutrients = float64(meristem.genome.GetRegulationvalue(MERISTEM_NUTRIENTS))
	meristem.nutrients -= newMeristemRoot.nutrients

	newMeristemRoot.connectXylem(meristem.xylem)
	newMeristemRoot.connectPhloem(meristem.phloem)

	//go newMeristemRoot.Run()
}

func (meristem *VineCell) detectMeristemLeafCell() bool {
	for _, cell := range meristem.connectedCells {
		functions := cell.genome.GetActiveFunctions()
		if len(functions) >= 2 {
			function1 := functions[0].GetRawCode()
			function2 := functions[1].GetRawCode()
			if function1[0] == MERISTEM_CODON && function2[0] == LEAF_CODON {
				return true
			}
		}
	}

	return false
}

func (meristem *VineCell) detectMeristemRootCell() bool {
	for _, cell := range meristem.connectedCells {
		functions := cell.genome.GetActiveFunctions()
		if len(functions) >= 2 {
			function1 := functions[0].GetRawCode()
			function2 := functions[1].GetRawCode()
			if function1[0] == MERISTEM_CODON && function2[0] == ROOT_CODON {
				return true
			}
		}
	}

	return false
}

func (meristem *VineCell) detectRootCell() bool {
	for _, cell := range meristem.connectedCells {
		functions := cell.genome.GetActiveFunctions()
		if len(functions) > 0 {
			function := functions[0].GetRawCode()
			if function[0] == ROOT_CODON {
				return true
			}
		}
	}

	return false
}

func (meristem *VineCell) basicMeristemRun() {
	// Dectect if there if light in the place
	// Detect if there is a meristem leaf cell
	// If there is not, create a new meristem leaf cell
	if meristem.detectThing(place.Light) {
		if !meristem.detectMeristemLeafCell() {
			meristem.newMeristemLeafCell()
		}
	}

	// Dectect if there if earth in the place
	// Detect if there is a meristem root cell
	// If there is not, create a new meristem root cell
	if meristem.detectThing(place.Earth) {
		if !meristem.detectMeristemRootCell() {
			meristem.newMeristemRootCell()
		}
	}

}

// Run the Meristem cell functions related to the xylem cells.
func (meristem *VineCell) meristemXylemRun() {
	// Create a new Xylem cell to connect to the meristem.
	neededEnergy := float64(meristem.genome.GetRegulationvalue(XYLEM_ENERGY))
	neededNutrients := float64(meristem.genome.GetRegulationvalue(XYLEM_NUTRIENTS))

	if meristem.energy >= neededEnergy && meristem.nutrients >= neededNutrients {
		meristem.xylem = meristem.newXylemCell()
		meristem.genome.DisableFunction(XYLEM_CODON)
	} else {
		meristem.requetNutrients()
		meristem.requetEnergy()
	}
}

// Run the Meristem cell functions related to the xylem cells.
func (meristem *VineCell) meristemPhloemRun() {
	// Create a new Phloem cell to connect to the meristem.
	neededEnergy := float64(meristem.genome.GetRegulationvalue(PHLOEM_ENERGY))
	neededNutrients := float64(meristem.genome.GetRegulationvalue(PHLOEM_NUTRIENTS))

	if meristem.energy >= neededEnergy && meristem.nutrients >= neededNutrients {
		meristem.phloem = meristem.newPhloemCell()

		meristem.xylem.connectPhloem(meristem.phloem)
		meristem.phloem.connectXylem(meristem.xylem)

		meristem.genome.DisableFunction(PHLOEM_CODON)
	} else {
		meristem.requetNutrients()
		meristem.requetEnergy()
	}
}

func (meristem *VineCell) mersitemLeafRun() {
	if !meristem.detectThing(place.Light) {
		meristem.genome.DisableFunction(PHLOEM_CODON)
		meristem.newMeristemLeafCell()
		return
	}
	// Create a new Phloem cell to connect to the meristem.
	//neededEnergy := float64(meristem.genome.GetRegulationvalue(LEAF_ENERGY))
	//neededNutrients := float64(meristem.genome.GetRegulationvalue(LEAF_NUTRIENTS))
}

func (meristem *VineCell) mersitemRootRun() {
	// Check if the cell is in earth
	if !meristem.detectThing(place.Earth) {
		meristem.die()
		return
	}
	// Check if there already is a root cell in the place.
	if meristem.detectRootCell() {
		return
	}

	// Create a new root cell.
	meristem.newRootCell()
}

// Execute the meristem functions.
func (meristem *VineCell) meristemRun(functions []genes.Gene) {

	secondaryFunction := functions[0].GetRawCode()
	if len(functions) > 1 {
		secondaryFunction = functions[1].GetRawCode()
	}

	switch secondaryFunction[0] {
	case XYLEM_CODON:
		// In this case the meristem is creating a new Xylem cell.
		meristem.meristemXylemRun()
	case PHLOEM_CODON:
		// In this case the meristem is creating a new Phloem cell.
		meristem.meristemPhloemRun()
	case LEAF_CODON:
		meristem.mersitemLeafRun()
	case ROOT_CODON:
		meristem.mersitemRootRun()
	case MERISTEM_CODON:
		meristem.basicMeristemRun()
	default:
		log.Println("Meristem cell with unknown function.")
		meristem.die()
		return
	}

}
