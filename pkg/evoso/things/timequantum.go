package things

import (
	"sync"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type TimeQuantum struct {
	sync.Mutex

	environment.Thing
}

func NewTimeQuantum() *TimeQuantum {
	return &TimeQuantum{
		Mutex: sync.Mutex{},

		Thing: *environment.NewThing(environment.TimeQuantum, nil),
	}
}
