package terminal

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/things"
)

type TerminalAutomata struct {
	sync.Mutex

	environment.Thing

	internalEnv environment.Environment

	alive bool
}

func NewTerminalAutomata() *TerminalAutomata {
	newTerminal := &TerminalAutomata{
		Mutex: sync.Mutex{},

		Thing: *environment.NewThing(environment.TerminalAutomata, nil),

		internalEnv: *environment.NewEnvironment(),

		alive: false,
	}

	newTerminal.SetObject(newTerminal)

	return newTerminal
}

func (ta *TerminalAutomata) getTerminalOutput() {
	for ta.alive {
		// Get one available time quantum to execute job.
		timeQuantum := ta.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)

		if timeQuantum != nil {
			// Get the oldest keyboard output
			terminalOutputThings := ta.Thing.GetEnvironment().LookThingsType(environment.TerminalOutput)
			if len(terminalOutputThings) > 0 {
				terminalOutput := terminalOutputThings[0].GetObject().(*things.TerminalOutput)

				if len(terminalOutputThings) > 1 {
					// Find the oldest keyboard output
					for _, terminalOutputThing := range terminalOutputThings {
						actualTerminalOutput := terminalOutputThing.GetObject().(*things.TerminalOutput)
						if actualTerminalOutput.GetTime().Before(terminalOutput.GetTime()) {
							terminalOutput = actualTerminalOutput
						}
					}
				}

				// Get Thing from environment
				if ta.Thing.GetEnvironment().GetThing(&terminalOutput.Thing) {
					ta.internalEnv.AddThing(&terminalOutput.Thing)
				}
			}

			// Give back the time quantum
			ta.Thing.GetEnvironment().AddThing(timeQuantum)
		}

	}
}

func (ta *TerminalAutomata) cleanTerminalOutputs() {
	for ta.alive {
		// Get one available time quantum to execute job.
		timeQuantum := ta.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)

		if timeQuantum != nil {

			// Get all keyboard outputs
			termOutputThings := ta.internalEnv.GetAllThingsType(environment.TerminalOutput)

			if len(termOutputThings) > 0 {
				termOutputObjects := make([]*things.TerminalOutput, len(termOutputThings))
				for i, thing := range termOutputThings {
					termOutputObjects[i] = thing.GetObject().(*things.TerminalOutput)
				}

				// Sort the vector in order of time received
				if len(termOutputObjects) > 1 {
					sort.SliceStable(termOutputObjects, func(i, j int) bool {
						return termOutputObjects[i].GetTime().Before(termOutputObjects[j].GetTime())
					})
				}

				// Create a new keyboar input with all data
				outputData := ""
				for _, terminalOutput := range termOutputObjects {
					outputData += terminalOutput.GetValue()
				}

				// Send to terminal
				fmt.Printf("%s", outputData)
			}

			// Give back the time quantum
			ta.Thing.GetEnvironment().AddThing(timeQuantum)
		}
		time.Sleep(time.Millisecond)
	}
}

func (ta *TerminalAutomata) Activate() {
	if ta.Thing.GetEnvironment() == nil {
		log.Println("terminal.TerminalAutomata.Activate - No place to activate")
		return
	}
	ta.alive = true

	go ta.getTerminalOutput()
	go ta.cleanTerminalOutputs()
}
