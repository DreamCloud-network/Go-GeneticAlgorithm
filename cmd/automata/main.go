package main

import (
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/automata/environment"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/automata/keyboard"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/automata/keyboardecho"
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
