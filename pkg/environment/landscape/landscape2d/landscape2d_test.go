package landscape2d

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
)

/*
func TestLandscape2d(t *testing.T) {
	t.Log("TestLandscape2d")

	land := New(10, 10)

	for pos, position := range land.Positions[0][0].nextPositions {
		if position != nil {
			t.Log(Directions(pos).String() + ": " + position.String())
		}
	}
}*/

func TestUUIDComparison(t *testing.T) {
	uuid1 := uuid.New()

	uuid2 := uuid.New()

	start := time.Now()
	equal := uuid1 == uuid2
	elapsed := time.Since(start)
	log.Println("Equal: ", equal)
	log.Println("Comparison time: ", elapsed)
}

func TestMap(t *testing.T) {
	t.Log("TestMap")

	land := New(10, 10)

	for cont, posWidth := range land.Positions {
		for cont2, posHeight := range posWidth {
			t.Log(cont, cont2, posHeight.String())
		}
	}
}
