package timequantumregulator

import (
	"testing"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/evoso/environment"
)

func TestGenome(t *testing.T) {
	t.Log("TestGenome")

	newGenome := NewBasicTimeQuantumRegulatorGenome()

	t.Log("Genome:", newGenome.String())
}

func TestTimeQuantum(t *testing.T) {
	t.Log("TestTimeQuantum")

	newEnvironment := environment.NewEnvironment()

	newTimeQuantumRegulator := NewQuantumRegulator()

	newEnvironment.AddThing(&newTimeQuantumRegulator.Thing)

	t.Log("Tings in place:", newEnvironment.PrintThingsInPlace())

	newTimeQuantumRegulator.Activate()

	for cont := 0; cont < 10; cont++ {
		t.Log("Tings in place:", newEnvironment.PrintThingsInPlace())
		time.Sleep(1 * time.Second)
	}
}
