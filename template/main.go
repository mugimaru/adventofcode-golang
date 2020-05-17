package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var srcTemplate = `package main

import (
	"github.com/mugimaru73/adventofcode/executor"
)

func run(input string) (interface{}, interface{}) {
	return "Part1wip", "Part2wip"
}

func main() {
	executor.Run(executor.ReadInput("%s"), run)
}
`

func main() {
	if len(os.Args) != 3 {
		panic("Use `template 2019 02` to generate src and input files for 2nd puzzle of AOC 2019")
	}

	var basePath = path.Join(os.Args[1], "day"+os.Args[2])
	var srcFile = basePath + ".go"
	var inputFile = basePath + ".input.txt"

	path, err := filepath.Abs(srcFile)
	handleError(err)

	fmt.Printf("src=%v\ninput=%v\nabs=%v\n", srcFile, inputFile, path)

	if checkFileExists(srcFile) || checkFileExists(inputFile) {
		panic("file exists")
	}

	createFile(srcFile, fmt.Sprintf(srcTemplate, inputFile))
	createFile(inputFile, "")
}

func createFile(path string, content string) {
	path, err := filepath.Abs(path)
	handleError(err)

	file, err := os.Create(path)
	handleError(err)

	defer file.Close()

	file.WriteString(content)
}

func checkFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
