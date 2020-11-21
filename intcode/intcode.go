package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Memory []int

func (m Memory) Get(i int) int {
	return m[i]
}

func (m Memory) Set(i int, value int) {
	m[i] = value
}

type Mode = int8

const ModePosition = Mode(0)
const ModeImmediate = Mode(1)

type Op struct {
	Opcode int
	Mode1  Mode
	Mode2  Mode
	Mode3  Mode
}

func loadOp(v int) Op {
	mode3 := v / 10000
	mode2 := v % 10000 / 1000
	mode1 := v % 1000 / 100
	op := v % 100
	return Op{op, Mode(mode1), Mode(mode2), Mode(mode3)}
}

func Run(mem *Memory, chIn, chOut chan int) {
	var ip int = 0

	for {
		op := loadOp(mem.Get(ip))
		switch op.Opcode {
		case 1, 2:
			p1 := loadParam(mem, ip+1, op.Mode1)
			p2 := loadParam(mem, ip+2, op.Mode2)
			target := mem.Get(ip + 3)

			if op.Opcode == 1 {
				mem.Set(target, p1+p2)
			} else {
				mem.Set(target, p1*p2)
			}
			ip += 4
		case 3:
			addr := mem.Get(ip + 1)
			input := <-chIn
			mem.Set(addr, input)
			ip += 2
		case 4:
			chOut <- loadParam(mem, ip+1, op.Mode1)
			ip += 2
		case 99:
			close(chOut)
			return
		default:
			panic(fmt.Sprint("invalid opcode ", op.Opcode))
		}
	}
}

func loadParam(mem *Memory, ip int, mode Mode) int {
	value := mem.Get(ip)

	switch mode {
	case ModePosition:
		return mem.Get(value)
	case ModeImmediate:
		return value
	default:
		panic("Unknown mode")
	}
}

func LoadProgram(input string) Memory {
	program := []int{}

	for _, v := range strings.Split(input, ",") {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		program = append(program, value)
	}

	return Memory(program)
}

func CopyProgram(program Memory) Memory {
	cp := make(Memory, len(program))
	copy(cp, program)
	return cp
}
