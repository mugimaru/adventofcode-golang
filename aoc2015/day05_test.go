package aoc2015

import (
	"testing"
)

func TestDay5isStringNice(t *testing.T) {
	cases := map[string]bool{
		"ugknbfddgicrmopn": true,
		"aaa":              true,
		"jchzalrnumimnmhp": false,
		"haegwjzuvuyypxyu": false,
		"dvszwmarrgswjxmb": false,
	}
	for str, exp := range cases {
		if !isStringNice(str) == exp {
			t.Errorf("`%s` is expected to be nice=%v", str, exp)
		}
	}
}

func TestDay5isStringNice2(t *testing.T) {
	cases := map[string]bool{
		"qjhvhtzxzqqjkmpb": true,
		"xxyxx":            true,
		"uurcxstgmygtbstg": false,
		"ieodomkazucvgmuy": false,
	}
	for str, exp := range cases {
		if !isStringNice2(str) == exp {
			t.Errorf("`%s` is expected to be nice2=%v", str, exp)
		}
	}
}
