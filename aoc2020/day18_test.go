package aoc2020

import (
	"testing"
)

func TestDay18CalcPart1(t *testing.T) {
	cases := map[string]int{
		"2 * 3 + (4 * 5)":                                 26,
		"5 + (8 * 3 + 9 + 3 * 4 * 3)":                     437,
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))":       12240,
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2": 13632,
	}

	for expr, expected := range cases {
		res := evalRPN(parseInfix(expr, day18OpaPart1))

		if res != expected {
			t.Errorf("`%v` expected=%v got=%v", expr, expected, res)
		}

	}
}

func TestDay18CalcPart2(t *testing.T) {
	cases := map[string]int{
		"1 + (2 * 3) + (4 * (5 + 6))":                     51,
		"2 * 3 + (4 * 5)":                                 46,
		"5 + (8 * 3 + 9 + 3 * 4 * 3)":                     1445,
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))":       669060,
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2": 23340,
	}

	for expr, expected := range cases {
		res := evalRPN(parseInfix(expr, day18OpaPart2))

		if res != expected {
			t.Errorf("`%v` expected=%v got=%v", expr, expected, res)
		}
	}
}
