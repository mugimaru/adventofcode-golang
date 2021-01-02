package aoc2020

import (
	"sort"
	"strings"
)

func init() {
	registerFun("05", SolveDay05)
}

type Ticket struct {
	code string
	row  int
	col  int
}

func (t Ticket) seatID() int {
	return t.row*8 + t.col
}

func decodeTicket(code string) Ticket {
	rows := []int{0, 127}
	cols := []int{0, 7}

	for _, rune := range []rune(code) {
		switch string(rune) {
		case "F":
			rows[1] = rows[0] + (rows[1]-rows[0])/2
		case "B":
			rows[0] = rows[1] - (rows[1]-rows[0])/2
		case "L":
			cols[1] = cols[0] + (cols[1]-cols[0])/2
		case "R":
			cols[0] = cols[1] - (cols[1]-cols[0])/2
		}
	}

	return Ticket{code: code, row: rows[0], col: cols[0]}
}

func SolveDay05(input string) (interface{}, interface{}) {

	ticketCodes := strings.Split(input, "\n")
	seatIDS := make([]int, len(ticketCodes))

	highestSeatID := 0

	for i, code := range ticketCodes {
		ticket := decodeTicket(code)
		seatID := ticket.seatID()
		seatIDS[i] = seatID

		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	return highestSeatID, findMissingSeatID(seatIDS)
}

func findMissingSeatID(ids []int) int {
	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		if ids[i] == ids[i+1]-2 {
			return ids[i] + 1
		}
	}

	panic(":(")
}
