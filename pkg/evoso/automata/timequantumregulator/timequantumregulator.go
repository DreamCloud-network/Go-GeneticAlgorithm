package timequantumregulator

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/automata/basicmachinary"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

type TimeQuantumRegulator struct {
	basicmachinary.BasicMachinary
}

func NewTimeQuantumRegulator() *TimeQuantumRegulator {
	newQR := &TimeQuantumRegulator{
		BasicUnit: *basicunit.NewBasicUnit(environment.TimeQuantumRegulator),
	}

	newQR.SetObject(newQR)

	return newQR
}

func (tqr *TimeQuantumRegulator) Activate() {
	if tqr.GetEnvironment() == nil {
		log.Println("TimeQuantumRegulator.Activate - environment not set")
	}

	tqr.BasicUnit.Activate()

	NewTimeQuantumProducer(&tqr.BasicUnit)

}
