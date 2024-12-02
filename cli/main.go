package main

import (
	"flag"
	"fmt"
	"github.com/mugimaru73/adventofcode-golang/aoc2023"
	"github.com/mugimaru73/adventofcode-golang/aoc2024"
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
		if err := aoc2015.Run(day, readInput()); err != nil {
			panic(err)
		}
	case "2019":
		err := aoc2019.Run(day, readInput())
		if err != nil {
			panic(err)
		}
	case "2020":
		err := aoc2020.Run(day, readInput())
		if err != nil {
			panic(err)
		}
	case "2023":
		err := aoc2023.Run(day, readInput())
		if err != nil {
			panic(err)
		}
	case "2024":
		err := aoc2024.Run(day, readInput())
		if err != nil {
			panic(err)
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
