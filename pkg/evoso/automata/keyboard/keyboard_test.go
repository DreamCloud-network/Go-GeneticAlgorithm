package keyboard

import (
	"testing"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/automata/keyboardechoterminal"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/automata/terminal"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/things"
)

func TestKeyboardAutoma(t *testing.T) {
	t.Log("TestKeyboardAutoma")

	testEnvironment := environment.NewEnvironment()
	for cont := 0; cont < 10; cont++ {
		newTimeQuantum := things.NewTimeQuantum()
		testEnvironment.AddThing(&newTimeQuantum.Thing)
	}

	newKeyboardAutomata := NewKeyboardAutomata()
	testEnvironment.AddThing(&newKeyboardAutomata.Thing)
	newKeyboardAutomata.Activate()

	newKeyboardEcho := keyboardechoterminal.NewKeyboardEchoTerminalAutomata()
	testEnvironment.AddThing(&newKeyboardEcho.Thing)
	newKeyboardEcho.Activate()

	newTerminal := terminal.NewTerminalAutomata()
	testEnvironment.AddThing(&newTerminal.Thing)
	newTerminal.Activate()

	for {
		time.Sleep(time.Second)
	}
}
