package vine

func (meristem *VineCell) newPhloemCell() *VineCell {
	newPhloem := NewVineCell()

	newPhloem.SetPosition(meristem.place)

	duplicatedGenome := meristem.genome.Duplicate()
	newPhloem.genome = &duplicatedGenome

	newPhloem.genome.ActivateOnlyFunction(PHLOEM_CODON)

	// Send energy and nutrients to the new Phloem cell.
	newPhloem.energy = float64(meristem.genome.GetRegulationvalue(PHLOEM_ENERGY))
	meristem.energy -= newPhloem.energy

	newPhloem.nutrients = float64(meristem.genome.GetRegulationvalue(PHLOEM_NUTRIENTS))
	meristem.nutrients -= newPhloem.nutrients

	newPhloem.connectCells(meristem.phloem)

	meristem.connectPhloem(newPhloem)

	//go newPhloem.Run()

	return newPhloem
}

// Return true if the cell is a phloem cell
func (xy *VineCell) isPhloem() bool {
	functions := xy.genome.GetActiveFunctions()
	if len(functions) != 1 {
		return false
	}

	function := functions[0].GetRawCode()

	return (function[0] == PHLOEM_CODON)

}

// Receive energy from the leaf cells.
func (ph *VineCell) receiveEnergy(energy float64) {
	ph.mu.Lock()
	defer ph.mu.Unlock()

	ph.energy += energy
}

// Send energy for the requesting cell
func (ph *VineCell) getEnergy() float64 {
	ph.mu.Lock()
	defer ph.mu.Unlock()

	// Receive nutients from the root.
	if ph.energy > 10 {
		ph.energy--
		return 1
	}

	return 0
}

func (ph *VineCell) phloemRun() {
	// Get all cylem cells connected
	connPh := ph.getAllCellsPrimaryFunctionConnected(PHLOEM_CODON)

	// find connection with less nutrients
	var phloemCellsToSendNutrients *VineCell
	lessEnergy := -1.0
	for _, actualCell := range connPh {
		if actualCell.energy < lessEnergy {
			lessEnergy = actualCell.energy
			phloemCellsToSendNutrients = actualCell
		}
	}

	if phloemCellsToSendNutrients == nil {
		return
	}

	// Verify if there is less nutrients that the actual xylem cell
	if phloemCellsToSendNutrients.energy < ph.energy {
		// Send nutrients to the xylem cell
		phloemCellsToSendNutrients.receiveEnergy(1)
		ph.energy--
	}

}
