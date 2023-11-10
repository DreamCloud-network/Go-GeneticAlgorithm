package vine

func (meristem *VineCell) newXylemCell() *VineCell {
	newXylem := NewVineCell()

	newXylem.SetPosition(meristem.place)

	duplicatedGenome := meristem.genome.Duplicate()
	newXylem.genome = &duplicatedGenome

	newXylem.genome.ActivateOnlyFunction(XYLEM_CODON)

	// Send energy and nutrients to the new Xylem cell.
	newXylem.energy = float64(meristem.genome.GetRegulationvalue(XYLEM_ENERGY))
	meristem.energy -= newXylem.energy

	newXylem.nutrients = float64(meristem.genome.GetRegulationvalue(XYLEM_NUTRIENTS))
	meristem.nutrients -= newXylem.nutrients

	newXylem.connectCells(meristem.xylem)

	meristem.connectXylem(newXylem)

	//go newXylem.Run()

	return newXylem
}

// Return true if the cell is a xylem cell
func (xy *VineCell) isXylem() bool {
	functions := xy.genome.GetActiveFunctions()
	if len(functions) != 1 {
		return false
	}

	function := functions[0].GetRawCode()

	return (function[0] == XYLEM_CODON)

}

// Receive nutrients from the root cells.
func (xy *VineCell) receiveNutrients(nutients float64) {
	xy.mu.Lock()
	defer xy.mu.Unlock()

	xy.nutrients += nutients
}

// Send nutrients for the requesting cell
func (xy *VineCell) getNutrients() float64 {
	xy.mu.Lock()
	defer xy.mu.Unlock()

	// Receive nutients from the root.
	if xy.nutrients > 10 {
		xy.nutrients--
		return 1
	}

	return 0
}

func (xy *VineCell) xylemRun() {
	// Get all cylem cells connected
	connXy := xy.getAllCellsPrimaryFunctionConnected(XYLEM_CODON)

	// find connection with less nutrients
	var xylemCellsToSendNutrients *VineCell
	lessNutrinets := -1.0
	for _, actualCell := range connXy {
		if actualCell.nutrients < lessNutrinets {
			lessNutrinets = actualCell.nutrients
			xylemCellsToSendNutrients = actualCell
		}
	}

	if xylemCellsToSendNutrients == nil {
		return
	}

	// Verify if there is less nutrients that the actual xylem cell
	if xylemCellsToSendNutrients.nutrients < xy.nutrients {
		// Send nutrients to the xylem cell
		xylemCellsToSendNutrients.receiveNutrients(1)
		xy.nutrients--
	}

}
