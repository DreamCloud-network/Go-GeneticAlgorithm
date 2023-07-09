package landscape2d

import (
	"errors"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/landscape"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/environment/thingstype"
)

var (
	ErrorInvalidPosition = errors.New("invalid position")
	ErrorOutOfBoundry    = errors.New("out of boundry")
)

type Landscape2D struct {
	Positions [][]*Position2D
}

// New creates a new square landscape with a flat surface (one portion by position).
func New(width, height int) Landscape2D {
	var NewSquare Landscape2D

	NewSquare.Positions = make([][]*Position2D, width)

	for w := 0; w < width; w++ {
		NewSquare.Positions[w] = make([]*Position2D, height)
		for h := 0; h < height; h++ {
			NewSquare.Positions[w][h] = NewPosition2D(w, h)
		}
	}

	// Link all positions
	for width, widthPosition := range NewSquare.Positions {
		for height, position := range widthPosition {
			position.nextPositions[Center] = position

			if (height + 1) < len(NewSquare.Positions) {
				position.nextPositions[East] = NewSquare.Positions[width][height+1]
			}

			if (height) > 0 {
				position.nextPositions[West] = NewSquare.Positions[width][height-1]
			}

			if (width + 1) < len(widthPosition) {
				position.nextPositions[South] = NewSquare.Positions[width+1][height]
				if (height + 1) < len(NewSquare.Positions) {
					position.nextPositions[SouthEast] = NewSquare.Positions[width+1][height+1]
				}

				if (height) > 0 {
					position.nextPositions[SouthWest] = NewSquare.Positions[width+1][height-1]
				}
			}

			if (width) > 0 {
				position.nextPositions[North] = NewSquare.Positions[width-1][height]

				if (height + 1) < len(NewSquare.Positions) {
					position.nextPositions[NorthEast] = NewSquare.Positions[width-1][height+1]
				}

				if (height) > 0 {
					position.nextPositions[NorthWest] = NewSquare.Positions[width-1][height-1]
				}
			}
		}
	}
	return NewSquare
}

func (landscape *Landscape2D) GetThings() []landscape.Thing {
	/*	var allThings []landscape.Thing

		for _, positions := range landscape.Positions {
			for _, position := range positions {
				allThings = append(allThings, position.GetThings()...)
			}
		}
		return allThings*/
	return nil
}

func (landscape *Landscape2D) CleanAllThings() {
	for _, positions := range landscape.Positions {
		for _, position := range positions {
			position.CleanAllThings()
		}
	}
}

// GetPosition returns the position at the given coordinates or return ou of boundry error.
func (landscape *Landscape2D) GetPosition(position *Position2D) (*Position2D, error) {
	if position.X < 0 || position.X >= len(landscape.Positions) {
		return nil, ErrorOutOfBoundry
	}

	if position.Y < 0 || position.Y >= len(landscape.Positions[0]) {
		return nil, ErrorOutOfBoundry
	}
	return landscape.Positions[position.X][position.Y], nil
}

func (environment *Landscape2D) TotalThingByType(thingType thingstype.ThingType) int {
	total := 0

	for _, positions := range environment.Positions {
		for _, position := range positions {
			total += len(position.GetThingsByType(thingType))
		}
	}

	return total
}

func (environment *Landscape2D) GetAllThingByType(thingType thingstype.ThingType) []landscape.Thing {
	things := make([]landscape.Thing, 0)

	for _, positions := range environment.Positions {
		for _, position := range positions {
			things = append(things, position.GetThingsByType(thingType)...)
		}
	}

	return things
}

func (environment *Landscape2D) DrawsASCII() string {
	landASCII := "\n\r"

	landASCII += " "
	for range environment.Positions {
		landASCII += "---"
	}

	landASCII += "\n\r"

	for _, positions := range environment.Positions {
		landASCII += "|"
		for _, position := range positions {
			things := position.GetThings()
			robots := position.GetThingsByType(thingstype.Robot)

			if len(robots) > 0 {
				landASCII += " X "
			} else if len(things) > 0 {
				landASCII += " * "
			} else {
				landASCII += "   "
			}
		}
		landASCII += "|\n\r"
	}

	landASCII += " "

	for range environment.Positions {
		landASCII += "---"
	}

	landASCII += "\n\r"
	return landASCII
}
