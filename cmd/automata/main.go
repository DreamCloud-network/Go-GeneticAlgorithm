package main

import (
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/automata/environment"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/automata/keyboard"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/automata/keyboardecho"
)

func main() {

	testEnvironment := environment.NewEnvironment()
	testEnvironment.SetTimeQuantum(10)

	newKeyboardAutomata := keyboard.NewKeyboardAutomata()
	testEnvironment.AddThing(newKeyboardAutomata)

	newKeyboardEcho := keyboardecho.NewKeyboardEcho()
	testEnvironment.AddThing(newKeyboardEcho)

	newKeyboardAutomata.Activate()
	newKeyboardEcho.Activate()

	for {
		time.Sleep(time.Second)
	}
}
