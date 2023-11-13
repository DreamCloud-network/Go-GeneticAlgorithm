package basicmachinery

import (
	"log"
	"sync"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
)

// Contains the basic functions and genome that must be present in all automata
type BasicMachinery struct {
	sync.Mutex
	sync.WaitGroup

	environment.Thing

	genome dnastrand.DNAStrand

	alive bool
}

func NewBasicUnit(thingType environment.ThingType) *BasicMachinery {
	newBasicUint := &BasicMachinery{
		Mutex: sync.Mutex{},

		Thing: *environment.NewThing(thingType, nil),

		alive: false,
	}

	return newBasicUint
}

func (basicUnit *BasicMachinery) IsAlive() bool {
	return basicUnit.alive
}

// This function regulates the number of rimossome in the automata
func (basicUnit *BasicMachinery) ribossomomeRun() {
	basicUnit.Add(1)
	defer basicUnit.Done()

	// Loads the following values from the DNA
	maxRibossomes := 10

	numRRNAforNewRibossome := 50

	sleepTime := 100 //in milisseconds

	for basicUnit.alive {
		// Get one available time quantum to execute job.
		timeQuantum := basicUnit.GetExternalEnvironment().GetOneThingType(environment.TimeQuantum)
		if timeQuantum != nil {

			// Get the number of rrnas in the automata
			numRRNAs := basicUnit.GetInternalEnvironment().NumThingsType(environment.RRNA)

			if numRRNAs > numRRNAforNewRibossome {
				// Get the number of ribossomes in the automata
				numRibossomes := basicUnit.GetInternalEnvironment().NumThingsType(environment.Ribossome)

				if numRibossomes < maxRibossomes {
					// Create a new ribossome
					NewRibossome(basicUnit)
				}
			}

			// Give back the time quantum
			basicUnit.GetExternalEnvironment().AddThing(timeQuantum)
		}

		// Sleep for a while
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	}
}

func (basicUnit *BasicMachinery) Activate() {
	if basicUnit.GetExternalEnvironment() == nil {
		log.Println("basicunit.BasicUnit.Activate - No place to activate")
		return
	}
	basicUnit.alive = true

	//go basicUnit.ribossomomeRun()
}
