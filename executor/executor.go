package executor

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Run solves the puzzle input with given fun
func Run(input string, run func(string) (interface{}, interface{}, interface{})) {
	startedAt := time.Now()
	part1, part2, meta := run(input)

	fmt.Printf("Part1 = %v\nPart2 = %v\n  solved in %s", part1, part2, time.Since(startedAt))
	if meta != nil {
		fmt.Printf("  %v\n", meta)
	}
}

func ReadInput(defaultPath string) string {
	var sourceFile = defaultPath
	if len(os.Args) > 1 {
		sourceFile = os.Args[1]
	}

	data, err := ioutil.ReadFile(sourceFile)

	if err != nil {
		panic(err)
	}

	return string(data)
}
