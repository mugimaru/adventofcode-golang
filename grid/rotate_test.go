package grid_test

import (
	"testing"

	"github.com/mugimaru73/adventofcode-golang/grid"
)

func TestRotateGrid(t *testing.T) {
	g := grid.ParseGrid(`.##
.#.
##.`)

	expected := grid.ParseGrid(`#..
###
..#`)

	rotated := grid.Rotate(g)

	if !grid.Equal(expected, rotated) {
		t.Errorf("expected:\n%s\n\nrotated:\n%s", grid.Sprint(expected), grid.Sprint(rotated))
	}
}
