package mitochondria

import (
	"sync"
	"testing"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/place"
	"github.com/DreamCloud-networkrkrk/Go-GeneticAlgorithm/pkg/place/things"
)

func TestMitochondria(t *testing.T) {
	t.Log("TestMitochondria")

	testPlace := place.NewPlace()

	for cont := 0; cont < 100000; cont++ {
		newNutrient := things.NewNutrient()
		testPlace.AddThing(newNutrient)
	}

	t.Log("\n\rThing in place: ", testPlace.PrintThingsInPlace())

	// Create mitochondria
	mitochondrias := make([]*MitochondriaCell, 0)
	for cont := 0; cont < 10; cont++ {
		mitochondrias = append(mitochondrias, NewMitochondriaCell())
		testPlace.AddThing(mitochondrias[cont])
	}

	var wg sync.WaitGroup
	start := time.Now()
	for _, mitochondria := range mitochondrias {
		wg.Add(1)
		go func(mitochondria *MitochondriaCell) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			mitochondria.Activate()

		}(mitochondria)
	}

	wg.Wait()
	evaluateTime := time.Since(start)

	t.Logf("\n\rEvaluate took %s", evaluateTime)
	t.Log("\n\rThing in place: ", testPlace.PrintThingsInPlace())
}

func TestPlaceAndThings(t *testing.T) {
	t.Log("TestPlaceAndThings")

	newPlace := place.NewPlace()

	start := time.Now()
	for cont := 0; cont < 100000; cont++ {
		newNutrient := things.NewNutrient()
		newPlace.AddThing(newNutrient)
	}
	evaluateTime := time.Since(start)
	t.Logf("\n\rAdd things took %s", evaluateTime)
	t.Log("\n\rThings in place: ", newPlace.PrintThingsInPlace())

	var wg sync.WaitGroup
	start = time.Now()
	for cont := 0; cont < 100; cont++ {
		wg.Add(1)
		go func(newPlace *place.Place, jobNum int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			cont := 0

			for {
				thing := newPlace.GetOneThingType(place.Nutrient)
				if thing == nil {
					break
				}
				cont++
			}

			t.Log("\n\rJob", jobNum, " - Num things: ", cont)
		}(&newPlace, cont)
	}
	wg.Wait()
	evaluateTime = time.Since(start)
	t.Logf("\n\rGet one thing took %s", evaluateTime)
}
