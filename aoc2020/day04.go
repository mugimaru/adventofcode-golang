package aoc2020

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registerFun("04", SolveDay04)
}

func SolveDay04(input string) (interface{}, interface{}) {
	passports := parsePassports(input)
	validPassportsP1 := 0
	validPassportsP2 := 0

	for _, passport := range passports {
		if validatePassportP1(passport) {
			validPassportsP1++
			if validatePassportP2(passport) {
				validPassportsP2++
			}
		}
	}

	return validPassportsP1, validPassportsP2
}

func parsePassports(puzzleInput string) []map[string]string {
	input := strings.Split(puzzleInput, "\n\n")
	passports := make([]map[string]string, len(input))

	for i, inp := range input {
		inp = strings.ReplaceAll(inp, "\n", " ")
		passport := make(map[string]string)

		for _, v := range strings.Split(inp, " ") {
			kw := strings.Split(v, ":")
			passport[kw[0]] = kw[1]
		}
		passports[i] = passport
	}

	return passports
}

func validatePassportP1(passport map[string]string) bool {
	hasAllRequiredFields := true
	for _, f := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := passport[f]; !ok {
			hasAllRequiredFields = false
		}

	}

	return hasAllRequiredFields
}

func validatePassportP2(p map[string]string) bool {
	byr, error := strconv.Atoi(p["byr"])
	if error != nil || byr < 1920 || byr > 2002 {
		return false
	}

	iyr, error := strconv.Atoi(p["iyr"])
	if error != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, error := strconv.Atoi(p["eyr"])
	if error != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	if valid := validateHeight(p["hgt"]); !valid {
		return false
	}

	if match, _ := regexp.MatchString("#(([0-9])|([a-f])){6}", p["hcl"]); !match {
		return false
	}

	if match, _ := regexp.MatchString("^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$", p["ecl"]); !match {
		return false
	}

	if match, _ := regexp.MatchString("^\\d{9}$", p["pid"]); !match {
		return false
	}

	return true
}

func validateHeight(hgt string) bool {
	value, err := strconv.Atoi(string([]rune(hgt)[0 : len(hgt)-2]))
	if err != nil {
		return false
	}

	switch string([]rune(hgt)[len(hgt)-2:]) {
	case "cm":
		return value >= 150 && value <= 193
	case "in":
		return value >= 59 && value <= 76
	default:
		return false
	}
}
