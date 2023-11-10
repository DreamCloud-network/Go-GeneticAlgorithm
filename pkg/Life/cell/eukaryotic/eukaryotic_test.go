package eukaryotic

import (
	"math"
	"testing"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/place/things"
)

func TestCell(t *testing.T) {

	testPlace := place.NewPlace()

	for cont := 0; cont < 100; cont++ {
		newNutrient := things.NewNutrient()
		testPlace.AddThing(newNutrient)
	}

	t.Log("\n\rThing in place: ", testPlace.PrintThingsInPlace())

	// Create cells
	cells := make([]*EukaryoticCell, 0)
	for cont := 0; cont < 100; cont++ {
		cells = append(cells, NewEukaryoticCell())
		testPlace.AddThing(cells[cont])
		newNutrient := things.NewNutrient()
		cells[cont].cytoplasm.AddThing(newNutrient)
		cells[cont].Activate()
	}

	/*
		var wg sync.WaitGroup
		start := time.Now()
		for _, cell := range cells {
			wg.Add(1)
			go func(cell *EukaryoticCell) {
				// Decrement the counter when the goroutine completes.
				defer wg.Done()

				cell.Activate()

			}(cell)
		}

		wg.Wait()
		evaluateTime := time.Since(start)

		t.Logf("\n\rEvaluate took %s", evaluateTime)
		t.Log("\n\rThing in place: ", testPlace.PrintThingsInPlace())*/

	for oneAlive(cells) {
		t.Log("\n\rThing in place: ", testPlace.PrintThingsInPlace())
		time.Sleep(time.Second)
	}
}

func oneAlive(cells []*EukaryoticCell) bool {
	for _, cell := range cells {
		if cell.alive {
			return true
		}
	}
	return false
}

func sqrtNewton(x float64) float64 {
	const epsilon = 1e-9 // Precisão desejada
	z := x               // Suposição inicial

	for {
		// Aproxime a raiz quadrada usando o método de Newton-Raphson
		nextZ := z - (z*z-x)/(2*z)

		// Verifique se a diferença entre as iterações é menor que a precisão desejada
		if math.Abs(nextZ-z) < epsilon {
			return nextZ
		}

		z = nextZ
	}
}

func TestSqrt(t *testing.T) {
	x := 20.0 // Número para calcular a raiz quadrada

	// Use a função sqrtNewton para calcular a raiz quadrada
	result := sqrtNewton(x)

	// Compare o resultado com a função math.Sqrt para verificar a precisão
	actual := math.Sqrt(x)

	t.Logf("Raiz quadrada de %.2f calculada: %.6f\n", x, result)
	t.Logf("Raiz quadrada de %.2f real (math.Sqrt): %.6f\n", x, actual)
}
