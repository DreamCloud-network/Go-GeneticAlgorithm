package keyboardechoterminal

import (
	"log"
	"sort"
	"sync"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/things"
)

type KeyboardEchoTerminalAutomata struct {
	sync.Mutex
	sync.WaitGroup

	environment.Thing

	internalEnv environment.Environment

	life  int
	alive bool
}

func NewKeyboardEchoTerminalAutomata() *KeyboardEchoTerminalAutomata {
	newKeyboardEcho := &KeyboardEchoTerminalAutomata{
		Mutex:     sync.Mutex{},
		WaitGroup: sync.WaitGroup{},

		Thing: *environment.NewThing(environment.KeyboardEchoTerminalAutomata, nil),

		internalEnv: *environment.NewEnvironment(),

		// TO DO: This initial life value may be in the DNA
		life:  25,
		alive: false,
	}

	newKeyboardEcho.Thing.SetObject(newKeyboardEcho)

	return newKeyboardEcho
}

func (keta *KeyboardEchoTerminalAutomata) receiveKeyboardInput() {
	keta.Add(1)
	defer keta.Done()

	for keta.alive {
		// Get one available time quantum to execute job.
		timeQuantum := keta.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)

		if timeQuantum != nil {

			// Get the oldest keyboard input
			keyboardThings := keta.Thing.GetEnvironment().LookThingsType(environment.KeyboardInput)

			if len(keyboardThings) > 0 {

				keybInputObject := keyboardThings[0].GetObject().(*things.KeyboardInput)
				if len(keyboardThings) > 1 {
					// Find the oldest keyboard input
					for _, keyboardThing := range keyboardThings {
						actualInputObject := keyboardThing.GetObject().(*things.KeyboardInput)
						if actualInputObject.GetTime().Before(keybInputObject.GetTime()) {
							keybInputObject = actualInputObject
						}
					}
				}

				// Get Thing from environment
				if keta.Thing.GetEnvironment().GetThing(&keybInputObject.Thing) {
					// Create a new keyboard output
					newTerminalOutput := things.NewTerminalOutput(keybInputObject.GetValue())
					newTerminalOutput.SetTime(keybInputObject.GetTime())

					keta.internalEnv.AddThing(&newTerminalOutput.Thing)

					// TO DO: This value should be in the DNA
					keta.life += 5
				}
			}

			// Give back the time quantum
			keta.Thing.GetEnvironment().AddThing(timeQuantum)
		}

		if keta.life > 0 {
			keta.life--
		}
	}
}

// Clean the internal environment sending all the received inputs to the external environment
func (keta *KeyboardEchoTerminalAutomata) cleanReadInputs() {
	keta.Add(1)
	defer keta.Done()

	for keta.alive {
		// Get one available time quantum to execute job.
		timeQuantum := keta.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)

		if timeQuantum != nil {

			// Get all keyboard outputs
			termOutputThings := keta.internalEnv.GetAllThingsType(environment.TerminalOutput)

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

					// Create a new keyboar input with all data
					outputData := ""
					for _, terminalOutput := range termOutputObjects {
						outputData += terminalOutput.GetValue()
					}

					termOutputObjects[0].SetValue(outputData)
				}

				// Add the new keyboard input to the environment
				keta.Thing.GetEnvironment().AddThing(&termOutputObjects[0].Thing)
			}

			// Give back the time quantum
			keta.Thing.GetEnvironment().AddThing(timeQuantum)
		}
		if keta.life > 0 {
			keta.life--
		}
		//time.Sleep(time.Millisecond)
	}
}

func (keta *KeyboardEchoTerminalAutomata) die() {
	keta.alive = false

	keta.Wait()

	keta.Thing.GetEnvironment().GetThing(&keta.Thing)
}

func (keta *KeyboardEchoTerminalAutomata) basic() {
	keta.Add(1)
	defer keta.Done()

	for keta.alive {
		// Get one available time quantum to execute job.
		timeQuantum := keta.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)
		if timeQuantum != nil {
			if keta.life <= 0 {
				// Verify if this is the last Keyboardecho
				if keta.Thing.GetEnvironment().NumThingsType(environment.KeyboardEchoTerminalAutomata) > 1 {
					go keta.die()
					return
				} else {
					keta.life = 0
				}
			} else if keta.life > 100 {
				// Duplicate
				newKeyboardEcho := NewKeyboardEchoTerminalAutomata()
				keta.GetEnvironment().AddThing(&newKeyboardEcho.Thing)
				newKeyboardEcho.Activate()
				keta.life = 25
			}

			if keta.life > 0 {
				keta.life--
			}

			// Give back the time quantum
			keta.Thing.GetEnvironment().AddThing(timeQuantum)
		}
	}
}

func (keta *KeyboardEchoTerminalAutomata) Activate() {
	if keta.Thing.GetEnvironment() == nil {
		log.Println("keyboardecho.KeyboardEcho.Activate - No place to activate")
		return
	}
	keta.alive = true

	go keta.receiveKeyboardInput()
	go keta.cleanReadInputs()
	//go keta.basic()
}
