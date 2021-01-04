package grid_test

import (
	"testing"

	"github.com/mugimaru73/adventofcode-golang/grid"
)

func TestFlipGrid(t *testing.T) {
	g := grid.ParseGrid(`.##
.#.
##.
`)

	expected := grid.ParseGrid(`##.
.#.
.##
`)

	flipped := grid.Flip(g)

	if !grid.Equal(expected, flipped) {
		t.Errorf("expected:\n%s\n\nflipped:\n%s", grid.Sprint(expected), grid.Sprint(flipped))
	}
}
