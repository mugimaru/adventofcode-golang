package aoc2023

import (
	"sort"
	"strconv"
	"strings"
)

func init() {
	registerFun("07", SolveDay07)
}

var camelCardsValues = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

var camelCardsValuesP2 = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

type camelCardsPlayer struct {
	Hand    []rune
	Bid     int
	score   int
	scoreP2 int
}

func newPlayer(hand []rune, bid int) *camelCardsPlayer {
	return &camelCardsPlayer{Hand: hand, Bid: bid}
}

func SolveDay07(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	players := make([]*camelCardsPlayer, len(lines))

	for i, line := range lines {
		l := strings.Split(line, " ")
		bid, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}

		players[i] = newPlayer([]rune(l[0]), bid)
		players[i].score = calculateScoreP1(players[i].Hand)
		players[i].scoreP2 = calculateScoreP2(players[i].Hand)
	}

	sort.Slice(players, func(i, j int) bool {
		p := players[j]
		p2 := players[i]

		return compareHands(p.Hand, p2.Hand, p.score, p2.score, camelCardsValues)
	})
	totalP1 := calculateTotalScore(players)

	sort.Slice(players, func(i, j int) bool {
		p := players[j]
		p2 := players[i]

		return compareHands(p.Hand, p2.Hand, p.scoreP2, p2.scoreP2, camelCardsValuesP2)
	})
	totalP2 := calculateTotalScore(players)

	return totalP1, totalP2
}

func compareHands(h1 []rune, h2 []rune, s1 int, s2 int, cardValues map[rune]int) bool {
	if s1 > s2 {
		return true
	}

	if s1 == s2 {
		for i := 0; i < len(h1); i++ {
			if cardValues[h1[i]] == cardValues[h2[i]] {
				continue
			}
			return cardValues[h1[i]] > cardValues[h2[i]]
		}
		panic("hands are equal")
	}

	return false
}

func calculateTotalScore(players []*camelCardsPlayer) int {
	total := 0
	for i := 0; i < len(players); i++ {
		total += (i + 1) * players[i].Bid
	}

	return total
}

func calculateScoreP1(hand []rune) int {
	cntMap := make(map[rune]int)
	for _, r := range hand {
		cntMap[r] += 1
	}
	cntMapKeys := make([]rune, len(cntMap))
	i := 0
	for k := range cntMap {
		cntMapKeys[i] = k
		i++
	}
	sort.Slice(cntMapKeys, func(i, j int) bool {
		return cntMap[cntMapKeys[i]] > cntMap[cntMapKeys[j]]
	})

	if len(cntMap) == 1 {
		return 6
	}

	if len(cntMap) == 2 && cntMap[cntMapKeys[0]] == 4 {
		return 5
	}

	if len(cntMap) == 2 && cntMap[cntMapKeys[0]] == 3 {
		return 4
	}

	if cntMap[cntMapKeys[0]] == 3 {
		return 3
	}

	if cntMap[cntMapKeys[0]] == 2 && cntMap[cntMapKeys[1]] == 2 {
		return 2
	}

	if cntMap[cntMapKeys[0]] == 2 {
		return 1
	}

	return 0
}

func calculateScoreP2(hand []rune) int {
	cntMap := make(map[rune]int)
	var cntJ int
	for _, r := range hand {
		if r == 'J' {
			cntJ++
		} else {
			cntMap[r] += 1
		}
	}
	cntMapKeys := make([]rune, len(cntMap))
	var i int
	for k := range cntMap {
		cntMapKeys[i] = k
		i++
	}
	sort.Slice(cntMapKeys, func(i, j int) bool {
		return cntMap[cntMapKeys[i]] > cntMap[cntMapKeys[j]]
	})

	if len(cntMap) <= 1 {
		return 6
	}

	if cntMap[cntMapKeys[0]]+cntJ == 4 {
		return 5
	}

	if cntMap[cntMapKeys[0]]+cntJ == 3 && cntMap[cntMapKeys[1]] == 2 {
		return 4
	}

	if cntMap[cntMapKeys[0]]+cntJ == 3 {
		return 3
	}

	if cntMap[cntMapKeys[0]]+cntJ == 2 && cntMap[cntMapKeys[1]] == 2 {
		return 2
	}

	if cntMap[cntMapKeys[0]]+cntJ == 2 {
		return 1
	}

	return 0
}
