package grid_test

import (
	"testing"

	"github.com/mugimaru73/adventofcode-golang/grid"
)

func TestAllOrientations(t *testing.T) {
	g := grid.ParseGrid(`.##
.#.
##.`)

	allOrient := grid.AllOrientations(g)

	if len(allOrient) != 8 {
		t.Logf("grid is expected to have 8 possible orientations got %d", len(allOrient))
		t.FailNow()
	}

	r1 := grid.Rotate(g)
	r2 := grid.Rotate(r1)
	r3 := grid.Rotate(r2)
	f1 := grid.Flip(g)
	f2 := grid.Flip(r1)
	f3 := grid.Flip(r2)
	f4 := grid.Flip(r3)

	expected := [][][]rune{g, r1, r2, r3, f1, f2, f3, f4}

	for i, expGrid := range expected {
		if !grid.Equal(expGrid, allOrient[i]) {
			t.Errorf("expected gird at %d:\n%s\n\ngot:\n%s", i, grid.Sprint(expGrid), grid.Sprint(allOrient[i]))
		}
	}
}
