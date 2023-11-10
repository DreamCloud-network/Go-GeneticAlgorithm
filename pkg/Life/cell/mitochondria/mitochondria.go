package mitochondria

import (
	"log"
	"sync"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place/things"
	"github.com/google/uuid"
)

type MitochondriaCell struct {
	wg sync.WaitGroup

	id   uuid.UUID
	Type place.ThingType

	genome Genome

	cytoplasm     place.Place
	externalPlace *place.Place

	alive bool
}

func (mit *MitochondriaCell) GetID() uuid.UUID {
	return mit.id
}

func (mit *MitochondriaCell) GetType() place.ThingType {
	return mit.Type
}

func (mit *MitochondriaCell) GetPlace() *place.Place {
	return mit.externalPlace
}

func (mit *MitochondriaCell) SetPlace(newPlace *place.Place) {
	mit.externalPlace = newPlace
}

func NewMitochondriaCell() *MitochondriaCell {
	newMitochondria := &MitochondriaCell{
		wg:   sync.WaitGroup{},
		id:   uuid.New(),
		Type: place.Life,

		genome: newGenome(),

		cytoplasm:     place.NewPlace(),
		externalPlace: nil,

		alive: false,
	}

	for cont := 0; cont < 1; cont++ {
		newNutrient := things.NewNutrient()
		newMitochondria.cytoplasm.AddThing(newNutrient)

		newEnergy := things.NewEnergy()
		newMitochondria.cytoplasm.AddThing(newEnergy)
	}

	return newMitochondria
}

func (mit *MitochondriaCell) die() {
	mit.alive = false
}

func (mit *MitochondriaCell) generateEnergy() {
	conversionRate, err := mit.genome.getConversionRate()
	if err != nil {
		log.Println("mitochondria.Mitochondria.Run - error getting conversion rate from genome")
		mit.alive = false
		return
	}

	// converts nutrients into energy
	// Remove one nutrient
	thing := mit.cytoplasm.GetOneThingType(place.Nutrient)
	if thing != nil {
		// Generate energy to Mitochondria
		newEnergy := things.NewEnergy()
		mit.cytoplasm.AddThing(newEnergy)

		// Generate energy to host cell
		for cont := 0; cont < int(conversionRate); cont++ {
			newEnergy := things.NewEnergy()
			mit.externalPlace.AddThing(newEnergy)
		}
	}
}

func (mit *MitochondriaCell) absorbsNutrients() {
	nutrientsAbsorptionRate := 10
	maxNutrientsInCell := 100

	for mit.alive {
		numInternalNutrients := mit.cytoplasm.CountThingsType(place.Nutrient)
		numExternalNutrients := mit.externalPlace.CountThingsType(place.Nutrient)

		if (numInternalNutrients < numExternalNutrients) && (numInternalNutrients < maxNutrientsInCell) {
			thing := mit.externalPlace.GetOneThingType(place.Nutrient)
			if thing != nil {
				nutrient := thing.(*things.Nutrient)
				mit.cytoplasm.AddThing(nutrient)
			}
		}

		time.Sleep(time.Millisecond * time.Duration(nutrientsAbsorptionRate))
	}
}

func (mit *MitochondriaCell) basicFunctions() {

	for mit.alive {

		mit.generateEnergy()

		//consumes energy
		energy := mit.cytoplasm.GetOneThingType(place.Energy)
		if energy == nil {
			mit.die()
			return
		}

		//log.Println("Internal Nutrients: ", mit.cytoplasm.CountThingsType(place.Nutrient))
		//log.Println("External Nutrients: ", mit.externalPlace.CountThingsType(place.Nutrient))
		//log.Println("Internal Energy: ", mit.cytoplasm.CountThingsType(place.Energy))
		//log.Println("External Energy: ", mit.externalPlace.CountThingsType(place.Energy))

		time.Sleep(time.Millisecond * 10)
	}

}

func (mit *MitochondriaCell) Activate() {
	mit.alive = true

	// Start sbsorbing nutrients function.
	go mit.absorbsNutrients()

	// Start basic functions.
	go mit.basicFunctions()
	//mit.basicFunctions()
}
