package eukaryotic

import (
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/Life/cell/mitochondria"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place/things"
	"github.com/google/uuid"
)

type EukaryoticCell struct {
	id   uuid.UUID
	Type place.ThingType

	genome Genome

	cytoplasm     place.Place
	externalPlace *place.Place

	alive bool
}

func (cell *EukaryoticCell) GetID() uuid.UUID {
	return cell.id
}

func (cell *EukaryoticCell) GetType() place.ThingType {
	return cell.Type
}

func (cell *EukaryoticCell) GetPlace() *place.Place {
	return cell.externalPlace
}

func (cell *EukaryoticCell) SetPlace(newPlace *place.Place) {
	cell.externalPlace = newPlace
}

func (cell *EukaryoticCell) die() {
	cell.alive = false
}

func NewEukaryoticCell() *EukaryoticCell {
	newEukaryoticCell := &EukaryoticCell{
		id:   uuid.New(),
		Type: place.Life,

		genome: newGenome(),

		cytoplasm:     place.NewPlace(),
		externalPlace: nil,

		alive: false,
	}

	// Add a mitochondria
	newEukaryoticCell.cytoplasm.AddThing(mitochondria.NewMitochondriaCell())

	for cont := 0; cont < 1; cont++ {
		newNutrient := things.NewNutrient()
		newEukaryoticCell.cytoplasm.AddThing(newNutrient)

		newEnergy := things.NewEnergy()
		newEukaryoticCell.cytoplasm.AddThing(newEnergy)
	}

	return newEukaryoticCell
}

func (eukCell *EukaryoticCell) getMitochondrias() []*mitochondria.MitochondriaCell {
	mitochondrias := make([]*mitochondria.MitochondriaCell, 0)

	for _, thing := range eukCell.cytoplasm.LookThingsType(place.Life) {
		mitochondrias = append(mitochondrias, thing.(*mitochondria.MitochondriaCell))
	}

	return mitochondrias
}

func (eukCell *EukaryoticCell) activateMitochondrias() {
	mitochondrias := eukCell.getMitochondrias()
	for _, mitochondria := range mitochondrias {
		mitochondria.Activate()
	}
}

func (cell *EukaryoticCell) basicFunctions() {

	for cell.alive {

		//consumes energy
		energy := cell.cytoplasm.GetOneThingType(place.Energy)
		if energy == nil {
			cell.die()
			return
		}

		//log.Println("Internal Nutrients: ", cell.cytoplasm.CountThingsType(place.Nutrient))
		//log.Println("External Nutrients: ", cell.externalPlace.CountThingsType(place.Nutrient))
		//log.Println("Internal Energy: ", cell.cytoplasm.CountThingsType(place.Energy))
		//log.Println("External Energy: ", cell.externalPlace.CountThingsType(place.Energy))
		time.Sleep(time.Millisecond * 100)
	}

}

func (cell *EukaryoticCell) absorbsNutrients() {
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
func (eukCell *EukaryoticCell) Activate() {
	eukCell.alive = true

	// Start sbsorbing nutrients function.
	go eukCell.absorbsNutrients()

	// Activa Mitochondrias
	eukCell.activateMitochondrias()

	// Start basic functions.
	//go mit.basicFunctions()
	go eukCell.basicFunctions()
}
