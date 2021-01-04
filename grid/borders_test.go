package grid_test

import (
	"testing"

	"github.com/mugimaru73/adventofcode-golang/grid"
)

func TestGetRemoveBorders(t *testing.T) {
	g := grid.ParseGrid(".##\n.#.\n##.")
	ng := grid.RemoveBorders(g)

	if grid.Sprint(ng) != "\n#" {
		t.Errorf("expected\n%s\n got %s\n", "#", grid.Sprint(ng))
	}
}

func TestGetRowCol(t *testing.T) {
	g := grid.ParseGrid(".##\n.#.\n##.")

	if r := grid.GetRow(g, 1); !grid.BordersEqual(r, []rune(".#.")) {
		t.Errorf("GetRow(1) expected=`%s` got=`%s`", ".#.", string(r))
	}

	if r := grid.GetRow(g, 2); !grid.BordersEqual(r, []rune("##.")) {
		t.Errorf("GetRow(2) expected=`%s` got=`%s`", "##.", string(r))
	}

	if c := grid.GetCol(g, 0); !grid.BordersEqual(c, []rune(`..#`)) {
		t.Errorf("GetCol(0) expected=`%s` got=`%s`", "..#", string(c))
	}

	if c := grid.GetCol(g, 2); !grid.BordersEqual(c, []rune(`#..`)) {
		t.Errorf("GetCol(2) expected=`%s` got=`%s`", "#..", string(c))
	}
}
