package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/mugimaru73/adventofcode-golang/aoc2015"
	"github.com/mugimaru73/adventofcode-golang/aoc2019"
	"github.com/mugimaru73/adventofcode-golang/aoc2020"
)

var year string
var day string
var inputFilePath string

func init() {
	var dayInt int
	flag.IntVar(&dayInt, "day", 1, "day")
	flag.StringVar(&year, "year", "2020", "year")
	flag.StringVar(&inputFilePath, "input", "", "path to input file")

	flag.Parse()
	day = fmt.Sprintf("%02d", dayInt)
}

func main() {
	switch year {
	case "2015":
		if error := aoc2015.Run(day, readInput()); error != nil {
			panic(error)
		}
	case "2019":
		error := aoc2019.Run(day, readInput())
		if error != nil {
			panic(error)
		}
	case "2020":
		error := aoc2020.Run(day, readInput())
		if error != nil {
			panic(error)
		}
	default:
		panic("not implemented")
	}
}

func readInput() string {
	if inputFilePath == "" {
		inputFilePath = fmt.Sprintf("inputs/%s/day%s.txt", year, day)
	}

	content, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	return string(content)
}
