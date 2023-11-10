package vine

import (
	"log"
	"sync"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place"
	"github.com/google/uuid"
)

// This verctor contains all the vine cells.
// It is used only for debug purposes.
var AllVineCells []*VineCell

type VineCell struct {
	mu sync.Mutex

	id   uuid.UUID
	Type place.ThingType

	genome *Genome

	place *place.Place

	xylem          *VineCell
	phloem         *VineCell
	connectedCells []*VineCell

	alive     bool
	energy    float64
	nutrients float64
}

func NewVineCell() *VineCell {
	newViceCll := &VineCell{
		mu: sync.Mutex{},

		id:   uuid.New(),
		Type: place.Life,

		genome: NewGenome(),

		place: nil,

		xylem:          nil,
		phloem:         nil,
		connectedCells: make([]*VineCell, 0),

		alive:     true,
		energy:    0,
		nutrients: 0,
	}

	AllVineCells = append(AllVineCells, newViceCll)

	return newViceCll
}

func (vc *VineCell) GetID() uuid.UUID {
	return vc.id
}
func (vc *VineCell) GetType() place.ThingType {
	return vc.Type
}

func (vc *VineCell) GetPosition() *place.Place {
	return vc.place
}

func (vc *VineCell) SetPosition(newPlace *place.Place) error {
	err := place.MoveThingToPlace(vc, newPlace)
	if err != nil {
		log.Println("vine.VineCell.SetPosition - error moving vine cell to new place")
		return err
	}
	vc.place = newPlace
	return nil
}

// Return if the cell is alive or not.
func (vc *VineCell) IsAlive() bool {
	return vc.alive
}

// Connect two cells.
func (vc *VineCell) connectCells(otherVc *VineCell) {
	if otherVc == nil {
		return
	}

	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.connectedCells = append(vc.connectedCells, otherVc)
	otherVc.connectedCells = append(otherVc.connectedCells, vc)
}

func (vc *VineCell) connectXylem(xylem *VineCell) {
	if !xylem.isXylem() {
		return
	}
	vc.xylem = xylem
	vc.connectCells(xylem)
}

func (vc *VineCell) connectPhloem(phloem *VineCell) {
	if !phloem.isPhloem() {
		return
	}
	vc.phloem = phloem
	vc.connectCells(phloem)
}

// Disconnect two cells.
func (vc *VineCell) disconnectCell(otherVc *VineCell) {
	if otherVc == nil {
		return
	}

	vc.mu.Lock()
	defer vc.mu.Unlock()

	if vc.xylem == otherVc {
		vc.xylem = nil
	}

	if vc.phloem == otherVc {
		vc.phloem = nil
	}

	for i, cell := range vc.connectedCells {
		if cell == otherVc {
			vc.connectedCells = append(vc.connectedCells[:i], vc.connectedCells[i+1:]...)
			break
		}
	}
}

func (vc *VineCell) detectThing(thing place.ThingType) bool {
	// Verify if there is earth in the place.
	if vc.place != nil {
		earthThings := vc.place.GetThingsType(thing)
		return len(earthThings) > 0
	}

	return false
}

// Return all the cell with the specified primary function connected to this cell.
func (xy *VineCell) getAllCellsPrimaryFunctionConnected(codon codons.Codon) []*VineCell {
	xy.mu.Lock()
	defer xy.mu.Unlock()

	connCells := make([]*VineCell, 0)
	for _, cell := range xy.connectedCells {
		functions := cell.genome.GetActiveFunctions()
		if len(functions) >= 1 {
			if functions[0].GetRawCode()[0] == codon {
				connCells = append(connCells, cell)
			}
		}
	}
	return connCells
}

func (vc *VineCell) requetNutrients() {
	if vc.xylem != nil {
		vc.nutrients += vc.xylem.getNutrients()
	}
}

func (vc *VineCell) requetEnergy() {
	if vc.phloem != nil {
		vc.energy += vc.phloem.getEnergy()
	}
}

// This functions tries to request energy from the xylem and phloem, and ballance the nutients and energy.
func (vc *VineCell) requestNutrientsAndEnergy() {
	// Request energy and nutrients from the Xylem and Phloem cell.
	// 10 is an arbitrary value that should be well analysed in the future.
	if vc.nutrients < 10 {
		vc.requetNutrients()
	}

	if vc.energy < 10 {
		vc.requetEnergy()
	}
}

// COnverts nutrients into energy
func (vc *VineCell) convertNutrientsIntoEnergy() {
	if vc.nutrients >= 2 {
		vc.nutrients -= 2
		vc.energy++
	}
}

func (vc *VineCell) die() {
	// remove the connections with all other cells.
	vc.xylem = nil
	vc.phloem = nil

	vc.connectedCells = nil

	vc.alive = false
}

// Disconnect from all dead cells.
func (vc *VineCell) disconnectFromDeadCells() {
	for _, cell := range vc.connectedCells {
		if !cell.IsAlive() {
			vc.disconnectCell(cell)
		}
	}
}

// Execute the basic functions common to all kind of cells.
func (vc *VineCell) basicFunctions() {
	// Disconnect from all dead cells.
	vc.disconnectFromDeadCells()

	// Request nutrients and energy for basic functions.
	vc.requestNutrientsAndEnergy()

	// If the cell has no energy, it will try to convert nutrients into energy.
	if vc.energy <= 1 {
		vc.convertNutrientsIntoEnergy()
	}
	// Discount the energy and nutrients for the basic functions.

	vc.energy--
	//vc.nutrients--

	// The cell dies if there is no nutrients and energy.
	if vc.energy <= 0 {
		vc.die()
	}
}

func (vc *VineCell) Run() {

	if vc.alive {
		vc.basicFunctions()

		functions := vc.genome.GetActiveFunctions()
		if len(functions) == 0 {
			log.Println("vine.VineCell.Run - no function to execute")
			vc.die()
			return
		}

		firstFunction := functions[0].GetRawCode()
		switch firstFunction[0] {
		case SEED_CODON:
			vc.seedRun()
		case MERISTEM_CODON:
			vc.meristemRun(functions)
		case XYLEM_CODON:
			vc.xylemRun()
		case PHLOEM_CODON:
			vc.phloemRun()
		//case LEAF_CODON:
		case ROOT_CODON:
			vc.rootRun()
		default:
			log.Println("vine.VineCell.Run - unknown function to execute")
			vc.die()
			return
		}
	}
}

func RunVine() {
	for i := range AllVineCells {
		AllVineCells[i].Run()
	}
}
