package keyboard

import (
	"sync"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type KeyboardInput struct {
	sync.Mutex

	environment.Thing

	timereceived time.Time
	kbInput      []rune
}

func NewKeyboardInput(input []rune) *KeyboardInput {
	newKBI := &KeyboardInput{
		Mutex: sync.Mutex{},

		Thing: *environment.NewThing(environment.KeyboardInput, nil),

		timereceived: time.Now(),
		kbInput:      input,
	}

	newKBI.SetObject(newKBI)

	return newKBI
}

func (kbi *KeyboardInput) AddValue(input rune) {
	kbi.Lock()
	defer kbi.Unlock()

	kbi.kbInput = append(kbi.kbInput, input)
}

func (kbi *KeyboardInput) SetValue(input []rune) {
	kbi.Lock()
	defer kbi.Unlock()

	kbi.kbInput = input
}

func (kbi *KeyboardInput) GetValue() []rune {
	kbi.Lock()
	defer kbi.Unlock()

	return kbi.kbInput
}

func (kbi *KeyboardInput) SetTime(t time.Time) {
	kbi.Lock()
	defer kbi.Unlock()

	kbi.timereceived = t
}

func (kbi *KeyboardInput) GetTime() time.Time {
	kbi.Lock()
	defer kbi.Unlock()

	return kbi.timereceived
}
