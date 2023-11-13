package things

import (
	"sync"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type TerminalOutput struct {
	sync.Mutex

	environment.Thing

	timereceived time.Time
	val          string
}

func NewTerminalOutput(input string) *TerminalOutput {
	newTerminalOutput := &TerminalOutput{
		Mutex: sync.Mutex{},

		Thing: *environment.NewThing(environment.TerminalOutput, nil),

		timereceived: time.Now(),
		val:          input,
	}

	newTerminalOutput.SetObject(newTerminalOutput)

	return newTerminalOutput
}

func (to *TerminalOutput) GetValue() string {
	to.Lock()
	defer to.Unlock()

	return to.val
}

func (to *TerminalOutput) SetValue(input string) {
	to.Lock()
	defer to.Unlock()

	to.val = input
}

func (to *TerminalOutput) GetTime() time.Time {
	to.Lock()
	defer to.Unlock()

	return to.timereceived
}

func (to *TerminalOutput) SetTime(t time.Time) {
	to.Lock()
	defer to.Unlock()

	to.timereceived = t
}
