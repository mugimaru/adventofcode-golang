package aoc2023

import (
	"github.com/mugimaru73/adventofcode-golang/utils"
	"regexp"
	"strings"
)

func init() {
	registerFun("08", SolveDay08)
}

var nodeIdRegexp = regexp.MustCompile("[A-Z0-9]{3}")

type mapNode struct {
	Left  string
	Right string
}

func SolveDay08(input string) (interface{}, interface{}) {
	nodes, directions := parseDay08Input(input)

	p1 := stepsToTarget("AAA", nodes, directions, isEndOfPath)
	return p1, solveDay08Part2(nodes, directions)
}

func isEndOfPath(s string) bool {
	return s == "ZZZ"
}

func isEndOfPathP2(s string) bool {
	return s[2] == 'Z'
}

func parseDay08Input(input string) (map[string]*mapNode, []rune) {
	l := strings.Split(input, "\n\n")

	nodes := make(map[string]*mapNode)
	for _, n := range strings.Split(l[1], "\n") {
		ids := nodeIdRegexp.FindAllString(n, -1)
		nodes[ids[0]] = &mapNode{Left: ids[1], Right: ids[2]}
	}

	return nodes, []rune(l[0])
}

func solveDay08Part2(m map[string]*mapNode, directions []rune) int {
	var dist []int
	for node := range m {
		if node[2] == 'A' {
			dist = append(dist, stepsToTarget(node, m, directions, isEndOfPathP2))
		}
	}

	return utils.LCM(dist[0], dist[1], dist[2:]...)
}

func stepsToTarget(currentNodeId string, m map[string]*mapNode, directions []rune, cond func(string) bool) int {
	var steps, iDir int

	for {
		if iDir == len(directions) {
			iDir = 0
		}

		if cond(currentNodeId) {
			return steps
		}

		if directions[iDir] == 'R' {
			currentNodeId = m[currentNodeId].Right
		} else {
			currentNodeId = m[currentNodeId].Left
		}
		steps++
		iDir++
	}
}
