package basicmachinery

import (
	"sync"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

type RRNA struct {
	sync.Mutex

	environment.Thing

	gene genes.Gene
}

func (rrna *RRNA) GetGene() *genes.Gene {
	rrna.Lock()
	defer rrna.Unlock()

	return &rrna.gene
}

type ribossome struct {
	Organelle

	idleSleepTime int //in milisseconds
}

func NewRibossome(basicUnit *BasicMachinery) {
	newRibossome := &ribossome{
		Organelle: *NewOrganelle(environment.Ribossome, basicUnit),
	}

	newRibossome.SetObject(newRibossome)

	basicUnit.GetInternalEnvironment().AddThing(&newRibossome.Thing)

	//newRibossome.Activate()
}

/*
func (rib *ribossome) Activate() {
	if rib.Thing.GetEnvironment() == nil {
		log.Println("basicunit.ribossome.Activate - No place to activate")
		return
	}

	// Loads the following values from the DNA
	rib.idleSleepTime = 100 //in milisseconds

	rib.alive = true

	rib.Run()
}

func (rib *ribossome) Run() {
	rib.basicUnit.Add(1)
	defer rib.basicUnit.Done()

	for rib.alive {
		// Get one available time quantum to execute job.
		timeQuantum := rib.basicUnit.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)
		if timeQuantum != nil {
			rrnaThing := rib.GetEnvironment().GetOneThingType(environment.RRNA)
			if rrnaThing != nil {
				rrna := rrnaThing.GetObject().(*RRNA)
				// Decode the RRNA into a function (protein).
				geneCode := rrna.GetGene().ReadCode()
				if geneCode[0] == codons.INIT_CODON {
					switch geneCode[1] {
					case ACTIVE_RECEPTOR_CODON:
						NewActiveReceptor(rib.basicUnit, rrna.GetGene())
					}
				}

				// Give back the time quantum
				rib.basicUnit.Thing.GetEnvironment().AddThing(timeQuantum)
			} else {
				// Give back the time quantum
				rib.basicUnit.Thing.GetEnvironment().AddThing(timeQuantum)

				// Sleep for a while
				time.Sleep(time.Duration(rib.idleSleepTime) * time.Millisecond)

			}
		}
	}
}
*/
