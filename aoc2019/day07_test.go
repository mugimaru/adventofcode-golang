package aoc2019

import (
	"testing"

	"github.com/mugimaru73/adventofcode-golang/intcode"
)

func TestTestSetting(t *testing.T) {
	expectedOut := 43210
	program := intcode.LoadProgram("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
	out := testSettings([5]int{4, 3, 2, 1, 0}, program)

	if out != expectedOut {
		t.Errorf("Expected out=%v, but out=%v", expectedOut, out)
	}
}

func TestPart1(t *testing.T) {
	expectedOut := 54321
	out, _ := SolveDay07("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")

	if out != expectedOut {
		t.Errorf("Expected out=%v, but out=%v", expectedOut, out)
	}
}

func TestPart2(t *testing.T) {
	expectedOut := 18216
	_, out := SolveDay07("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")

	if out != expectedOut {
		t.Errorf("Expected out=%v, but out=%v", expectedOut, out)
	}
}
