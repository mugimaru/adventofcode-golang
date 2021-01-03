package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("19", SolveDay19)
}

type d19Rule struct {
	char  rune
	rules [][]int
}

func SolveDay19(input string) (interface{}, interface{}) {
	inputParts := strings.Split(input, "\n\n")
	rules := parseDay19Rules(inputParts[0])

	messages := strings.Split(inputParts[1], "\n")
	p1 := countMessages(rules, messages)
	rules[8] = d19Rule{rules: [][]int{{42}, {42, 8}}}
	rules[11] = d19Rule{rules: [][]int{{42, 31}, {42, 11, 31}}}
	return p1, countMessages(rules, messages)
}

func parseDay19Rules(input string) map[int]d19Rule {
	lines := strings.Split(input, "\n")
	rules := make(map[int]d19Rule)

	for _, line := range lines {
		var char rune
		options := [][]int{}

		parts := strings.Split(line, ": ")
		id, _ := strconv.Atoi(parts[0])

		if strings.HasPrefix(parts[1], "\"") {
			char = []rune(parts[1])[1]
		} else {
			for _, option := range strings.Split(parts[1], "|") {
				branch := []int{}
				for _, v := range strings.Fields(option) {
					item, _ := strconv.Atoi(v)
					branch = append(branch, item)
				}

				options = append(options, branch)
			}
		}

		rules[id] = d19Rule{char, options}
	}

	return rules
}

func matchRule(word []rune, rules map[int]d19Rule, stack []int) bool {
	if len(stack) > len(word) {
		return false
	}
	if len(stack) == 0 || len(word) == 0 {
		return len(stack) == 0 && len(word) == 0
	}

	stack, c := stack[:len(stack)-1], stack[len(stack)-1]
	r := rules[c]
	if len(r.rules) == 0 {
		if word[0] == r.char {
			stackCp := make([]int, len(stack))
			copy(stackCp, stack)
			return matchRule(word[1:], rules, stackCp)
		}
	} else {
		for _, rule := range r.rules {
			newStack := make([]int, len(stack))
			copy(newStack, stack)
			for i := len(rule) - 1; i >= 0; i-- {
				newStack = append(newStack, rule[i])
			}

			if matchRule(word, rules, newStack) {
				return true
			}

		}
	}

	return false
}

func countMessages(rules map[int]d19Rule, messages []string) (total int) {
	for _, message := range messages {
		r := rules[0].rules[0]
		stack := make([]int, len(r))
		for i, v := range r {
			stack[len(r)-i-1] = v
		}

		if matchRule([]rune(message), rules, stack) {
			total++
		}
	}

	return
}
