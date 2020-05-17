package executor

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Run solves the puzzle input with given fun
func Run(input string, run func(string) (interface{}, interface{})) {
	startedAt := time.Now()
	part1, part2 := run(input)
	fmt.Printf("Part1 = %v\nPart2 = %v\n  solved in %s\n", part1, part2, time.Since(startedAt))
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
