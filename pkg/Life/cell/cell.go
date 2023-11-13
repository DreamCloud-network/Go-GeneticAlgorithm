package cell

import (
	"sync"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/things"
	"github.com/google/uuid"
)

type Cell struct {
	wg sync.WaitGroup

	id   uuid.UUID
	Type place.ThingType

	genome Genome

	cytoplasm     place.Place
	externalPlace *place.Place

	alive bool
}

func NewCell() *Cell {
	newCell := &Cell{
		wg: sync.WaitGroup{},

		id:   uuid.New(),
		Type: place.Life,

		genome: newGenome(),

		cytoplasm:     place.NewPlace(),
		externalPlace: nil,

		alive: false,
	}

	for cont := 0; cont < 1; cont++ {
		newNutrient := things.NewNutrient()
		newCell.cytoplasm.AddThing(newNutrient)

		newEnergy := things.NewEnergy()
		newCell.cytoplasm.AddThing(newEnergy)
	}

	return newCell
}

func (cell *Cell) GetID() uuid.UUID {
	return cell.id
}

func (cell *Cell) GetType() place.ThingType {
	return cell.Type
}

func (cell *Cell) GetPlace() *place.Place {
	return cell.externalPlace
}

func (cell *Cell) SetPlace(newPlace *place.Place) {
	cell.externalPlace = newPlace
}

func (cell *Cell) Die() {
	cell.alive = false
}

func (cell *Cell) absorbsNutrients() {
	nutrientsAbsorptionRate := 1
	maxNutrientsInCell := 100

	for cell.alive {
		numInternalNutrients := cell.cytoplasm.CountThingsType(place.Nutrient)
		numExternalNutrients := cell.externalPlace.CountThingsType(place.Nutrient)

		if (numInternalNutrients < numExternalNutrients) && (numInternalNutrients < maxNutrientsInCell) {
			thing := cell.externalPlace.GetOneThingType(place.Nutrient)
			if thing != nil {
				nutrient := thing.(*things.Nutrient)
				cell.cytoplasm.AddThing(nutrient)
			}
		}

		time.Sleep(time.Millisecond * time.Duration(nutrientsAbsorptionRate))
	}
}

func (cell *Cell) basicFunctions() {

	for cell.alive {

		//consumes energy
		energy := cell.cytoplasm.GetOneThingType(place.Energy)
		if energy == nil {
			cell.Die()
			return
		}

		time.Sleep(time.Second)
	}

}

func (cell *Cell) Activate() {
	cell.alive = true

	// Start sbsorbing nutrients function.
	go cell.absorbsNutrients()

	//go cell.basicFunctions()
	cell.basicFunctions()
}
