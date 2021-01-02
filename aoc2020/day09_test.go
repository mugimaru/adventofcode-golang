package aoc2020

import (
	"testing"
)

func TestDay9(t *testing.T) {
	var input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	numbers := parseDay09Input(input)
	invNum := findFirstInvalidNumber(numbers, 5)
	encWeakness := findEncryptionWeakness(numbers, invNum)
	if invNum != 127 || encWeakness != 62 {
		t.Errorf("expected invNum=127 encWeakness=62 got invNum=%v encWeakness=%v", invNum, encWeakness)
	}
}
