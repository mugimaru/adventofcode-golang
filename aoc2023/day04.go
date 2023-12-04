package aoc2023

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("04", SolveDay04)
}

type deck struct {
	Scratchcards []*scratchcard
}

func (d *deck) CalculateRewards() {
	for _, card := range d.Scratchcards {
		card.CalculateRewards(d)
	}
}

func (d *deck) DeckResults() (int, int) {
	var totalPoints, cardsTotal int
	for _, card := range d.Scratchcards {
		totalPoints += card.Points
		cardsTotal += card.ScratchcardsTotalWithRewards()
	}

	return totalPoints, cardsTotal
}

type scratchcard struct {
	ID             int
	WinningNumbers []int
	Numbers        []int
	Points         int
	Rewards        []*scratchcard
}

func (s *scratchcard) CalculateRewards(d *deck) {
	s.Points = 0
	var matchesCount int
	for i := 0; i < len(s.WinningNumbers); i++ {
		for j := 0; j < len(s.Numbers); j++ {
			if s.WinningNumbers[i] == s.Numbers[j] {
				matchesCount += 1
				if s.Points == 0 {
					s.Points = 1
				} else {
					s.Points *= 2
				}

				rewardID := s.ID + matchesCount
				if rewardID <= len(d.Scratchcards) {
					s.Rewards = append(s.Rewards, d.Scratchcards[rewardID-1])
				}
				break
			}
		}
	}
}

func (s *scratchcard) ScratchcardsTotalWithRewards() int {
	count := 1

	for _, card := range s.Rewards {
		count += card.ScratchcardsTotalWithRewards()
	}

	return count
}

func SolveDay04(input string) (interface{}, interface{}) {
	d := parseDay04Input(input)
	d.CalculateRewards()

	return d.DeckResults()
}

func parseDay04Input(input string) *deck {
	lines := strings.Split(input, "\n")
	d := &deck{Scratchcards: make([]*scratchcard, len(lines))}

	for i, line := range lines {
		d.Scratchcards[i] = parseScratchCard(i+1, line)
	}

	return d
}

func parseScratchCard(ID int, line string) *scratchcard {
	cards := strings.Split(strings.Split(line, ": ")[1], " | ")
	return &scratchcard{ID: ID, Numbers: parseNumbers(cards[1]), WinningNumbers: parseNumbers(cards[0])}
}

func parseNumbers(input string) []int {
	var nums []int
	for _, strNum := range strings.Split(input, " ") {
		strNum = strings.TrimLeft(strNum, " ")
		if strNum == "" {
			continue
		}
		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}
	return nums
}
