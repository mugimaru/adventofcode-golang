package aoc2015

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("07", solveDay07)
}

func solveDay07(input string) (interface{}, interface{}) {
	circuit := make(circuit)

	for _, line := range strings.Split(input, "\n") {
		var op *operation
		fields := strings.Fields(line)

		switch len(fields) {
		case 3:
			op = &operation{l: fields[0]}
		case 4:
			op = &operation{gate: fields[0], l: fields[1]}
		default:
			op = &operation{l: fields[0], gate: fields[1], r: fields[2]}
		}

		circuit[fields[len(fields)-1]] = op
	}

	a := circuit.getValue("a")
	circuit.reset()
	circuit.setValue("b", a)

	return a, circuit.getValue("a")
}

type operation struct {
	gate     string
	l, r     string
	hasValue bool
	value    uint16
}

type circuit map[string]*operation

func (c *circuit) reset() {
	for _, operation := range *c {
		operation.hasValue = false
	}
}

func (c *circuit) setValue(wire string, value uint16) {
	(*c)[wire].value = value
	(*c)[wire].hasValue = true
}

func (c *circuit) getValue(wire string) uint16 {
	operation, ok := (*c)[wire]
	if !ok {
		v, _ := strconv.Atoi(wire)
		return uint16(v)
	}

	if operation.hasValue {
		return operation.value
	}

	var value uint16
	switch operation.gate {
	case "":
		value = c.getValue(operation.l)
	case "NOT":
		value = ^c.getValue(operation.l)
	case "AND":
		value = c.getValue(operation.l) & c.getValue(operation.r)
	case "OR":
		value = c.getValue(operation.l) | c.getValue(operation.r)
	case "LSHIFT":
		value = c.getValue(operation.l) << c.getValue(operation.r)
	case "RSHIFT":
		value = c.getValue(operation.l) >> c.getValue(operation.r)
	}

	c.setValue(wire, value)
	return value
}
