package aoc2020

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	registerFun("14", SolveDay14)
}

func SolveDay14(input string) (interface{}, interface{}) {
	mem := make(map[string]int64)
	mem2 := make(map[string]int)
	mask := ""

	for _, line := range strings.Split(input, "\n") {
		fields := strings.Split(line, " = ")

		if strings.HasPrefix(line, "mask") {
			mask = fields[1]
		} else {
			value, _ := strconv.Atoi(fields[1])
			addr := fields[0][4 : len(fields[0])-1]
			numAddr, _ := strconv.Atoi(addr)

			addMaskedValue(mask, value, addr, &mem)
			addValueToMaskedAdresses(mask, numAddr, value, &mem2)
		}
	}

	sum := 0
	for _, value := range mem {
		sum += int(value)
	}

	sum2 := 0
	for _, value := range mem2 {
		sum2 += value
	}

	return sum, sum2
}

func addValueToMaskedAdresses(mask string, numAddr int, value int, mem *map[string]int) {
	address := lpad(fmt.Sprintf("%b", numAddr), 36, "0")
	maskedAddress := ""
	floatingBits := 0
	for i, bit := range mask {
		switch bit {
		case '1':
			maskedAddress += "1"
		case '0':
			maskedAddress += string(address[i])
		case 'X':
			maskedAddress += "X"
			floatingBits++
		}
	}

	for i := 0; i < int(math.Pow(2, float64(floatingBits))); i++ {
		comb := lpad(fmt.Sprintf("%b", i), floatingBits, "0")

		cpMaskedAddress := maskedAddress
		for _, bit := range comb {
			cpMaskedAddress = strings.Replace(cpMaskedAddress, "X", string(bit), 1)
		}

		(*mem)[cpMaskedAddress] = value
	}
}

func addMaskedValue(mask string, num int, addr string, mem *map[string]int64) {
	value := lpad(fmt.Sprintf("%b", num), 36, "0")
	maskedValue := ""
	for i, bit := range mask {
		if bit == 'X' {
			maskedValue += string(value[i])
		} else {
			maskedValue += string(bit)
		}
	}

	res, _ := strconv.ParseInt(maskedValue, 2, 64)
	(*mem)[addr] = res
}

func lpad(value string, n int, p string) string {
	for len(value) < n {
		value = p + value
	}

	return value
}
