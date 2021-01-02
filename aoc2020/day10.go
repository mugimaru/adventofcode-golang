package aoc2020

import (
	"sort"
)

func init() {
	registerFun("10", SolveDay10)
}

func SolveDay10(input string) (interface{}, interface{}) {
	adapters := parseDay09Input(input)
	sort.Ints(adapters)
	var prevAdapter int
	var diffs [3]int

	for _, adapter := range adapters {
		diffs[adapter-prevAdapter-1]++

		prevAdapter = adapter
	}
	diffs[2]++

	adapters = append(adapters, adapters[len(adapters)-1]+3)

	cache := make(map[int]int)
	adaptersMap := make(map[int]bool)
	for _, ad := range adapters {
		adaptersMap[ad] = true
	}

	return diffs[0] * diffs[2], countPaths(0, &cache, &adaptersMap)
}

func countPaths(start int, cache *map[int]int, adaptersMap *map[int]bool) int {
	if count, ok := (*cache)[start]; ok {
		return count
	}

	singleCandidate := true
	count := 0
	for i := 1; i <= 3; i++ {
		candidate := start + i
		if _, ok := (*adaptersMap)[candidate]; ok {
			count += countPaths(candidate, cache, adaptersMap)
			singleCandidate = false
		}
	}

	if singleCandidate {
		count++
	}

	(*cache)[start] = count

	return count
}
