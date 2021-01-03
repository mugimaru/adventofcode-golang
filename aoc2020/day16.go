package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("16", SolveDay16)
}

type ticketRule struct {
	Min int
	Max int
}

func SolveDay16(input string) (interface{}, interface{}) {
	fields, myTicket, nearbyTickets := parseDay16Input(input)
	var validTickets [][]int

	possibleFieldNames := make(map[string][]bool, len(fields))
	for field := range fields {
		possibleFieldNames[field] = make([]bool, len(myTicket))
		for i := 0; i < len(myTicket); i++ {
			possibleFieldNames[field][i] = true
		}
	}

	errorRate := 0
	for _, ticket := range append(nearbyTickets, myTicket) {
		isTicketValid := true
		for index, value := range ticket {
			invalidFieldsForIndex := []string{}
			isIndexValid := false
			for fieldName, rules := range fields {
				isFieldValid := false
				for _, rule := range rules {
					if value >= rule.Min && value <= rule.Max {
						isFieldValid = true
						isIndexValid = true
					}
				}

				if !isFieldValid {
					invalidFieldsForIndex = append(invalidFieldsForIndex, fieldName)
				}
			}

			if isIndexValid {
				for _, fieldName := range invalidFieldsForIndex {
					possibleFieldNames[fieldName][index] = false
				}
			} else {
				errorRate += value
				isTicketValid = false
			}
		}

		if isTicketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	fieldIndexToName := make([]string, len(fields))
	unresolved := len(fields)
	for unresolved != 0 {
		for i := 0; i < len(fieldIndexToName); i++ {
			if fieldIndexToName[i] != "" {
				continue
			}

			for name, possibilities := range possibleFieldNames {
				var possibleIndexes []int
				for i, possible := range possibilities {
					if possible && fieldIndexToName[i] == "" {
						possibleIndexes = append(possibleIndexes, i)
					}
				}

				if len(possibleIndexes) == 1 {
					index := possibleIndexes[0]
					unresolved--
					fieldIndexToName[index] = name
					for j := 0; j < len(myTicket); j++ {
						if j != index {
							possibleFieldNames[name][j] = false
						}
					}
				}
			}
		}
	}

	p2 := 1
	for i, name := range fieldIndexToName {
		if strings.HasPrefix(name, "departure") {
			p2 *= myTicket[i]
		}
	}

	return errorRate, p2
}

func parseDay16Input(input string) (map[string][]ticketRule, []int, [][]int) {
	inp := strings.Split(input, "\n\n")

	fields := make(map[string][]ticketRule)

	for _, line := range strings.Split(inp[0], "\n") {
		r := strings.Split(line, ": ")
		rs := strings.Split(r[1], " or ")
		rules := make([]ticketRule, len(rs))

		for i, rule := range rs {
			splittedRule := strings.Split(rule, "-")
			min, _ := strconv.Atoi(splittedRule[0])
			max, _ := strconv.Atoi(splittedRule[1])
			rules[i] = ticketRule{min, max}
		}

		fields[r[0]] = rules
	}

	mtf := strings.Split(strings.Split(inp[1], "\n")[1], ",")
	myTicket := make([]int, len(mtf))

	for i, field := range mtf {
		myTicket[i], _ = strconv.Atoi(field)
	}

	ntf := strings.Split(inp[2], "\n")[1:]
	nearbyTickets := make([][]int, len(ntf))
	for i, t := range ntf {
		flds := strings.Split(t, ",")
		ticket := make([]int, len(flds))

		for j, field := range flds {
			ticket[j], _ = strconv.Atoi(field)
		}

		nearbyTickets[i] = ticket
	}

	return fields, myTicket, nearbyTickets
}
