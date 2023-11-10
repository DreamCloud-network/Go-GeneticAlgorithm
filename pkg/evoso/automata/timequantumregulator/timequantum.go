package timequantumregulator

import "github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"

type TimeQuantum struct {
	environment.Thing
}

func NewTimeQuantum() *TimeQuantum {
	newTQ := &TimeQuantum{
		Thing: *environment.NewThing(environment.TimeQuantum, nil),
	}

	return newTQ
}