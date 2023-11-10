package timequantumregulator

import (
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/automata/basicunit"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type TimeQuantumProducer struct {
	basicunit.Organelle
}

func NewTimeQuantumProducer(basicUnit *basicunit.BasicUnit) {
	newTQP := &TimeQuantumProducer{
		Organelle: *basicunit.NewOrganelle(environment.TimeQuantumProducer, basicUnit),
	}

	newTQP.SetObject(newTQP)

	basicUnit.GetInternalEnvironment().AddThing(&newTQP.Thing)

	go newTQP.Run()
}

// Add one time quantum to the environment if it is needed
func (tqp *TimeQuantumProducer) Run() {
	tqp.BasicUnit.Add(1)
	defer tqp.BasicUnit.Done()

	for tqp.BasicUnit.IsAlive() {
		// Verifies if there is time quantum available in the environment
		numTimeQuantums := tqp.BasicUnit.GetEnvironment().NumThingsType(environment.TimeQuantum)

		if numTimeQuantums <= 0 {
			// Create a new time quantum
			newTimeQuantum := environment.NewThing(environment.TimeQuantum, nil)

			// Send to the environment
			tqp.BasicUnit.GetEnvironment().AddThing(newTimeQuantum)
		}

		// Sleep for a while
		time.Sleep(time.Duration(tqp.IdleSleepTime) * time.Millisecond)
	}
}
