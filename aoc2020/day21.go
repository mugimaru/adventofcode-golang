package aoc2020

import (
	"sort"
	"strings"
)

func init() {
	registerFun("21", SolveDay21)
}

func SolveDay21(input string) (interface{}, interface{}) {
	ingredients := make(map[string]int)
	allergens := make(map[string][]string)

	for _, line := range strings.Split(input, "\n") {
		spl := strings.Split(line, " (contains ")
		ingr := strings.Fields(spl[0])
		for _, ingr := range ingr {
			ingredients[ingr]++
		}

		for _, allerg := range strings.Split(spl[1][:len(spl[1])-1], ", ") {
			if len(allergens[allerg]) == 0 {
				allergens[allerg] = ingr
			} else {
				allergens[allerg] = intersectStringSlices(allergens[allerg], ingr)
			}
		}
	}

	for {
		done := true
		for allergen, ings := range allergens {
			if len(ings) != 1 {
				done = false
			} else {
				for otherAllergen := range allergens {
					if otherAllergen != allergen {
						allergens[otherAllergen] = removeFromStringSlice(allergens[otherAllergen], ings[0])
					}
				}
			}
		}

		if done {
			break
		}
	}

	for _, ings := range allergens {
		for _, ing := range ings {
			delete(ingredients, ing)
		}
	}

	sum := 0
	for _, c := range ingredients {
		sum += c
	}

	var names []string
	for k := range allergens {
		names = append(names, k)
	}
	sort.Strings(names)

	var canonicalDIL []string
	for _, n := range names {
		canonicalDIL = append(canonicalDIL, allergens[n][0])
	}

	return sum, strings.Join(canonicalDIL, ",")
}

func removeFromStringSlice(s []string, v string) (res []string) {
	for _, item := range s {
		if item != v {
			res = append(res, item)
		}
	}
	return
}

func intersectStringSlices(s1 []string, s2 []string) (res []string) {
	m := make(map[string]int)

	for _, v := range s1 {
		m[v]++
	}

	for _, v := range s2 {
		m[v]++
	}

	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}

	return
}
