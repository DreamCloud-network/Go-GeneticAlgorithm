package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/things"
)

// This automata is responsible for reaidng the keyboard input and send it in string format to the environment
type KeyboardAutomata struct {
	sync.Mutex

	environment.Thing

	internalEnv environment.Environment

	alive bool
}

func NewKeyboardAutomata() *KeyboardAutomata {
	newKeyboard := &KeyboardAutomata{
		Mutex: sync.Mutex{},

		internalEnv: *environment.NewEnvironment(),

		alive: false,
	}

	newKeyboard.Thing = *environment.NewThing(environment.KeyboardAutomata, newKeyboard)

	return newKeyboard
}

func (kba *KeyboardAutomata) readInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite algo: ")
	for kba.alive {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		} else {
			newKeyboardInput := things.NewKeyboardInput(string(char))
			kba.internalEnv.AddThing(&newKeyboardInput.Thing)
		}
	}
}

func (kba *KeyboardAutomata) readInputTest() {
	testString := string("\n\rKeyboard input test.")
	for kba.alive {
		for _, char := range testString {
			newKeyboardInput := things.NewKeyboardInput(string(char))
			kba.internalEnv.AddThing(&newKeyboardInput.Thing)

			time.Sleep(time.Microsecond)
		}
	}
}

// Clean the internal environment sending all the received inputs to the external environment
func (kba *KeyboardAutomata) cleanReadInputs() {
	for kba.alive {
		// Get one available time quantum to execute job.
		timeQuantum := kba.Thing.GetEnvironment().GetOneThingType(environment.TimeQuantum)

		if timeQuantum != nil {

			// Get all keyboard inputs
			keyboardProduction := kba.internalEnv.GetAllThingsType(environment.KeyboardInput)

			if len(keyboardProduction) > 0 {
				inputObjVector := make([]*things.KeyboardInput, len(keyboardProduction))
				for i, keyboardInput := range keyboardProduction {
					inputObjVector[i] = keyboardInput.GetObject().(*things.KeyboardInput)
				}

				// Sort the vector in order of time received
				if len(inputObjVector) > 1 {
					sort.SliceStable(inputObjVector, func(i, j int) bool {
						return inputObjVector[i].GetTime().Before(inputObjVector[j].GetTime())
					})
				}

				// Create a new keyboar input with all data
				inputData := ""
				for _, keyboardInput := range inputObjVector {
					inputData += keyboardInput.GetValue()
				}
				exitData := things.NewKeyboardInput(inputData)

				// Add the new keyboard input to the environment
				kba.Thing.GetEnvironment().AddThing(&exitData.Thing)
			}

			// Give back the time quantum
			kba.Thing.GetEnvironment().AddThing(timeQuantum)
		}

		time.Sleep(time.Millisecond)
	}
}

func (kba *KeyboardAutomata) Activate() {
	if kba.Thing.GetEnvironment() == nil {
		log.Println("keyboard.KeyboardAutomata.Activate - No place to activate")
		return
	}
	kba.alive = true

	//go kba.readInput()
	go kba.readInputTest()

	go kba.cleanReadInputs()
}

/*
func NewKeyboardAutomata() *KeyboardAutomata {
	return &KeyboardAutomata{
		Mutex: sync.Mutex{},

		id:    uuid.New(),
		Type:  place.KeyboardAutomata,
		place: nil,

		internalPlace: place.NewPlace(),

		alive: false,
	}
}

func (kba *KeyboardAutomata) GetID() uuid.UUID {
	kba.Lock()
	defer kba.Unlock()

	return kba.id
}
func (kba *KeyboardAutomata) GetType() place.ThingType {
	kba.Lock()
	defer kba.Unlock()

	return kba.Type
}

func (kba *KeyboardAutomata) GetPlace() *place.Place {
	kba.Lock()
	defer kba.Unlock()

	return kba.place
}

// This function is not to be used to add or remove a thing from a place
func (kba *KeyboardAutomata) SetPlace(newPlace *place.Place) {
	//kba.Lock()
	//defer kba.Unlock()

	kba.place = newPlace
}

func (kba *KeyboardAutomata) readInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite algo: ")
	for kba.alive {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		} else {
			newKeyboardInput := things.NewKeyboardInput(string(char))
			kba.place.AddThing(newKeyboardInput)
		}
	}
}

func (kba *KeyboardAutomata) readInputTest() {
	testString := string("\n\rKeyboard input test.")
	for kba.alive {
		for _, char := range testString {
			newKeyboardInput := things.NewKeyboardInput(string(char))
			kba.internalPlace.AddThing(newKeyboardInput)

			time.Sleep(time.Microsecond * 100)
		}
	}
}

func (kba *KeyboardAutomata) cleanInternalSpace() {
	for kba.alive {
		// Get one available time quantum to execute job.
		timeQuantum := kba.place.GetOneThingType(place.TimeQuantum)

		if timeQuantum != nil {

			// Get all keyboard inputs
			keyboardProduction := kba.internalPlace.LookThingsType(place.KeyboardInput)

			// Give back the time quantum
			kba.place.AddThing(timeQuantum)
		}

		time.Sleep(time.Millisecond)
	}
}

func (kba *KeyboardAutomata) Activate() {
	if kba.place == nil {
		log.Println("keyboard.KeyboardAutomata.Activate - No place to activate")
		return
	}
	kba.alive = true

	//go kba.readInput()
	go kba.readInputTest()
}
*/
