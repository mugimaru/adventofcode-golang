package aoc2023

import (
	"github.com/mugimaru73/adventofcode-golang/utils"
	"log"
	"regexp"
	"sort"
	"strings"
)

func init() {
	registerFun("05", SolveDay05)
}

type AlmanacRange struct {
	SrcStart  int
	DestStart int
	Length    int
}

func (ar *AlmanacRange) withinSrcRange(value int) bool {
	return ar.SrcStart <= value && ar.SrcStart+ar.Length > value
}

func (ar *AlmanacRange) MapValue(value int) int {
	return ar.DestStart + value - ar.SrcStart
}

type AlmanacMap struct {
	Src    string
	Dest   string
	Ranges []*AlmanacRange
}

type Almanac struct {
	Maps map[string]*AlmanacMap
}

func (a *Almanac) PerformMapping(value int, from string, to string) int {
	m := a.Maps[from]
	if m == nil {
		log.Fatalf("unable to map %s->%s", from, to)
	}
	mappedValue := value
	for _, r := range m.Ranges {
		if r.withinSrcRange(value) {
			mappedValue = r.MapValue(value)
		}

		if r.SrcStart+r.Length > value {
			break
		}
	}

	if to == m.Dest {
		return mappedValue
	}

	return a.PerformMapping(mappedValue, m.Dest, to)
}

func SolveDay05(input string) (interface{}, interface{}) {
	a, seeds := parseDay05Input(input)

	seedToLocationCache := map[int]int{}
	mapWithCache := func(cache map[int]int, seed int) int {
		if res, ok := cache[seed]; ok {
			return res
		}
		cache[seed] = a.PerformMapping(seed, "seed", "location")
		return cache[seed]
	}

	var minLocation, minLocationP2 int
	for i, seed := range seeds {
		value := mapWithCache(seedToLocationCache, seed)
		if i == 0 || value < minLocation {
			minLocation = value
		}
	}

	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			value := mapWithCache(seedToLocationCache, seed)
			if i == 0 && seeds[i] == seed || value < minLocationP2 {
				minLocationP2 = value
			}
		}
	}
	return minLocation, minLocationP2
}

var almanacTitleRegexp *regexp.Regexp = regexp.MustCompile("(?P<from>.+)-to-(?P<to>.+) map")

func parseDay05Input(input string) (*Almanac, []int) {
	blocks := strings.Split(input, "\n\n")
	seeds := utils.ParseNumsRow(blocks[0][7:len(blocks[0])])
	almanac := &Almanac{Maps: make(map[string]*AlmanacMap)}

	for _, block := range blocks[1:] {
		am := &AlmanacMap{}
		lines := strings.Split(block, "\n")
		mapDesc := almanacTitleRegexp.FindAllStringSubmatch(lines[0], -1)[0]

		am.Dest = mapDesc[2]
		am.Src = mapDesc[1]

		for _, line := range lines[1:] {
			l := utils.ParseNumsRow(line)
			am.Ranges = append(am.Ranges, &AlmanacRange{SrcStart: l[1], DestStart: l[0], Length: l[2]})
		}

		sort.Slice(am.Ranges, func(i, j int) bool {
			return am.Ranges[i].SrcStart < am.Ranges[j].SrcStart
		})
		almanac.Maps[am.Src] = am
	}

	return almanac, seeds
}
