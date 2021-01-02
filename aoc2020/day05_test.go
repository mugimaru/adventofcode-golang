package aoc2020

import (
	"testing"
)

func TestDay5DecodeTicket(t *testing.T) {
	ticket := decodeTicket("FBFBBFFRLR")
	seatID := ticket.seatID()

	if ticket.row != 44 || ticket.col != 5 || seatID != 357 {
		t.Errorf("expected row=44 col=5 seatID=357 got ticket=%v seatID=%v", ticket, seatID)
	}
}
