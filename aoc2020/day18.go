package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("18", SolveDay18)
}

var day18OpaPart1 = map[string]int{
	"*": 0,
	"+": 0,
}

var day18OpaPart2 = map[string]int{
	"*": 0,
	"+": 1,
}

func SolveDay18(input string) (interface{}, interface{}) {
	sum1 := 0
	sum2 := 0
	for _, expr := range strings.Split(input, "\n") {
		sum1 += evalRPN(parseInfix(expr, day18OpaPart1))
		sum2 += evalRPN(parseInfix(expr, day18OpaPart2))
	}
	return sum1, sum2
}

// https://rosettacode.org/wiki/Parsing/Shunting-yard_algorithm#Go
func parseInfix(input string, opa map[string]int) (result []string) {
	var stack []string // holds operators and left parenthesis

	for _, tok := range strings.Split(strings.Replace(input, " ", "", -1), "") {
		switch tok {
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for len(stack) > 0 {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // Discard "("
				}

				result = append(result, op)
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					top := stack[len(stack)-1]
					if o2prec, isOp := opa[top]; !isOp || o1 > o2prec {
						break
					}

					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					result = append(result, top) // add it to result
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				result = append(result, tok) // add operand to result
			}
		}
	}

	// drain stack to result
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return
}

func evalRPN(tokens []string) int {
	var values []int

	for _, tok := range tokens {
		if num, err := strconv.Atoi(tok); err == nil {
			values = append(values, num)
		} else {
			ops := [2]int{}
			for i := 0; i < 2; i++ {
				ops[i], values = values[len(values)-1], values[:len(values)-1]
			}

			switch tok {
			case "*":
				values = append(values, ops[0]*ops[1])
			case "+":
				values = append(values, ops[0]+ops[1])
			}
		}
	}

	return values[0]
}
