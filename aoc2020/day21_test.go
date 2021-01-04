package aoc2020

import (
	"testing"
)

func TestDay21(t *testing.T) {
	input := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

	p1, p2 := SolveDay21(input)
	if p1 != 5 || p2 != "mxmxvkd,sqjhc,fvjkl" {
		t.Errorf("expected p1=5 p2=`mxmxvkd,sqjhc,fvjkl` p1=%v p2=`%v`", p1, p2)
	}
}
