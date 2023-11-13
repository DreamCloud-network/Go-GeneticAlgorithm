package organelle

import (
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/automata/basicmachinery"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

// Genome structure
//FUNCTION - SPACE - IDLESLEEPTIME

// Represents the basic struct of any automata internal organelle.
type Organelle struct {
	environment.Thing

	Machinary *basicmachinery.BasicMachinery

	IdleSleepTime int //in milisseconds

	MaxInternalSize int
}

func NewOrganelle(thingType environment.ThingType, mach *basicmachinery.BasicMachinery) *Organelle {
	newOrganelle := &Organelle{
		Thing: *environment.NewThing(thingType, nil),

		Machinary: mach,

		IdleSleepTime: 100, //in milisseconds

		MaxInternalSize: 100,
	}

	return newOrganelle
}

// Open and publicate a new channel for the organelle.
func (organelle *Organelle) OpenNewChannel(thingType environment.ThingType) *chan *environment.Thing {
	// Create the new channel
	newChannel := make(chan *environment.Thing)

	// Start listenning to the new channel
	go organelle.activateChannel(newChannel, thingType)

	return &newChannel
}

func (organelle *Organelle) activateChannel(channel chan *environment.Thing, thingTypeChannel environment.ThingType) {
	for organelle.Machinary.IsAlive() {
		thing := <-channel

		if thing != nil {
			if thing.GetType() == thingTypeChannel {
				if organelle.GetInternalEnvironment().NumThingsType(thingTypeChannel) < organelle.MaxInternalSize {
					organelle.GetInternalEnvironment().AddThing(thing)
				} else {
					organelle.GetExternalEnvironment().AddThing(thing)
				}
			} else {
				organelle.GetExternalEnvironment().AddThing(thing)
			}
		}
	}
}
