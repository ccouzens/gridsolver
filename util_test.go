package grid

import (
	"testing"
)

func TestSquareAt(t *testing.T) {
	grid := Grid{
		rowClues: make([]Clue, 2),
		squares: []Square{
			White, White,
			Black, White,
		},
	}

	if grid.SquareAt(0, 1) != Black {
		t.Error("Expected 0,1 to be Black")
	}

	grid.squares[2] = White
	if grid.SquareAt(0, 1) != White {
		t.Error("Expected 0,1 to be White")
	}
}
