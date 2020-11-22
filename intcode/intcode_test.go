package intcode

import (
	"strconv"
	"testing"
)

func TestDay05Features1(t *testing.T) {
	input := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	testCases := map[int]int{7: 999, 8: 1000, 9: 10001}

	for inp, expectedOut := range testCases {
		p := LoadProgram(input)
		out := RunAndRead(&p, []int64{int64(inp)})
		if len(out) != 1 && out[0] != 999 {
			t.Errorf("Expected out is %v, got %v", expectedOut, out)
		}
	}
}

func TestDay05FeaturesPosMode(t *testing.T) {
	input := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	testCases := map[int]int{0: 0, 1: 1, 5: 1}

	for inp, expectedOut := range testCases {
		p := LoadProgram(input)
		out := RunAndRead(&p, []int64{int64(inp)})
		if len(out) != 1 && out[0] != 999 {
			t.Errorf("Expected out is %v, got %v", expectedOut, out)
		}
	}
}

func TestDay05FeaturesImmMode(t *testing.T) {
	input := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	testCases := map[int]int{0: 0, 1: 1, 5: 1}

	for inp, expectedOut := range testCases {
		p := LoadProgram(input)
		out := RunAndRead(&p, []int64{int64(inp)})
		if len(out) != 1 && out[0] != 999 {
			t.Errorf("Expected out is %v, got %v", expectedOut, out)
		}
	}
}

func TestDay09Features1(t *testing.T) {
	mem := LoadProgram("104,1125899906842624,99")
	out := RunAndRead(&mem, []int64{})

	if int(out[0]) != 1125899906842624 {
		t.Errorf("Expected out=%v, but out=%v", 1125899906842624, out[0])
	}
}

func TestDay09Features2(t *testing.T) {
	mem := LoadProgram("1102,34915192,34915192,7,4,7,99,0")
	out := RunAndRead(&mem, []int64{})

	if len(strconv.FormatInt(out[0], 10)) != 16 {
		t.Errorf("Expected 16-digits number, got %v", out[0])
	}
}

func TestDay09Features3(t *testing.T) {
	mem := LoadProgram("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	out := RunAndRead(&mem, []int64{})
	expectedOut := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	for i := 0; i < len(out); i++ {
		if out[i] != expectedOut[i] {
			t.Errorf("Expected %v, got %v", expectedOut, out)
		}
	}
}
