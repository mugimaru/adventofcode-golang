package aoc2020

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registerFun("07", SolveDay07)
}

func SolveDay07(input string) (interface{}, interface{}) {
	rules := make(map[string]map[string]int)

	for _, line := range strings.Split(input, "\n") {
		inp := strings.Split(line, " bags contain ")
		bag := inp[0]
		re := regexp.MustCompile(`( bags?)|(\.)`)
		contains := make(map[string]int)

		for _, b := range strings.Split(re.ReplaceAllString(inp[1], ""), ", ") {
			if b == "no other" {
				continue
			}

			rb := []rune(b)
			amount, err := strconv.Atoi(string(rb[0]))
			if err != nil {
				panic(err)
			}

			contains[string(rb[2:])] = amount
		}

		rules[bag] = contains
	}

	cache := make(map[string]bool)
	cache["shiny gold"] = true

	p1 := -1
	for bag := range rules {
		if containsShinyGoldenBag(bag, &rules, &cache) {
			p1++
		}
	}

	return p1, amountOfBagsInside("shiny gold", &rules) - 1
}

func amountOfBagsInside(bag string, rules *map[string]map[string]int) int {
	sum := 1
	for b, count := range (*rules)[bag] {
		sum += count * amountOfBagsInside(b, rules)
	}

	return sum
}

func containsShinyGoldenBag(bag string, rules *map[string]map[string]int, cache *map[string]bool) bool {
	if (*cache)[bag] {
		return true
	}

	for b := range (*rules)[bag] {
		if containsShinyGoldenBag(b, rules, cache) {
			(*cache)[b] = true
			return true
		}
	}

	return false
}
