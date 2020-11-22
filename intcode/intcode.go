package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Value int64
type Memory map[Value]Value

func (m Memory) Get(i Value) Value {
	return m[i]
}

func (m Memory) Set(i Value, value Value) {
	if i < 0 {
		panic(fmt.Sprint("Negative memory addr ", i))
	}

	m[i] = value
}

func (m Memory) setOp(i Value, value Value, mode Mode, relBase Value) {
	if mode == ModeRelative {
		i += relBase
	}

	m[i] = value
}

type Mode = int8

const ModePosition = Mode(0)
const ModeImmediate = Mode(1)
const ModeRelative = Mode(2)

type Op struct {
	Opcode int
	Mode1  Mode
	Mode2  Mode
	Mode3  Mode
}

func (op Op) loadP1(ip Value, mem *Memory, relativeBase Value) Value {
	return loadParam(mem, ip+1, op.Mode1, relativeBase)
}

func (op Op) loadP2(ip Value, mem *Memory, relativeBase Value) Value {
	return loadParam(mem, ip+2, op.Mode2, relativeBase)
}

func loadParam(mem *Memory, ip Value, mode Mode, relativeBase Value) Value {
	value := mem.Get(ip)

	switch mode {
	case ModePosition:
		return mem.Get(value)
	case ModeImmediate:
		return value
	case ModeRelative:
		return mem.Get(value + relativeBase)
	default:
		panic("Unknown mode")
	}
}

func loadOp(v Value) Op {
	mode3 := v / 10000
	mode2 := v % 10000 / 1000
	mode1 := v % 1000 / 100
	op := v % 100
	return Op{int(op), Mode(mode1), Mode(mode2), Mode(mode3)}
}

func RunAndRead(mem *Memory, input []int64) []int64 {
	chIn := make(chan int64, 100)
	chOut := make(chan int64, 100)
	chDone := make(chan int)

	for _, v := range input {
		chIn <- int64(v)
	}

	go Run(mem, chIn, chOut, chDone)
	<-chDone
	close(chIn)
	close(chOut)

	out := []int64{}
	for v := range chOut {
		out = append(out, v)
	}
	return out
}

func Run(mem *Memory, chIn chan int64, chOut chan int64, chDone chan int) {
	ip := Value(0)
	relativeBase := Value(0)

	for {
		op := loadOp(mem.Get(ip))
		switch op.Opcode {
		case 1, 2:
			p1 := op.loadP1(ip, mem, relativeBase)
			p2 := op.loadP2(ip, mem, relativeBase)
			target := mem.Get(ip + 3)

			if op.Opcode == 1 {
				mem.setOp(target, p1+p2, op.Mode3, relativeBase)
			} else {
				mem.setOp(target, p1*p2, op.Mode3, relativeBase)
			}
			ip += 4
		case 3:
			addr := mem.Get(ip + 1)
			input := <-chIn
			mem.setOp(addr, Value(input), op.Mode1, relativeBase)
			ip += 2
		case 4:
			chOut <- int64(op.loadP1(ip, mem, relativeBase))
			ip += 2
		case 5, 6:
			p1 := op.loadP1(ip, mem, relativeBase)

			if (op.Opcode == 5 && p1 != 0) || (op.Opcode == 6 && p1 == 0) {
				ip = op.loadP2(ip, mem, relativeBase)
			} else {
				ip += 3
			}
		case 7, 8:
			p1 := op.loadP1(ip, mem, relativeBase)
			p2 := op.loadP2(ip, mem, relativeBase)
			p3 := mem.Get(ip + 3)

			if (op.Opcode == 7 && p1 < p2) || (op.Opcode == 8 && p1 == p2) {
				mem.setOp(p3, 1, op.Mode3, relativeBase)
			} else {
				mem.setOp(p3, 0, op.Mode3, relativeBase)
			}
			ip += 4
		case 9:
			p1 := op.loadP1(ip, mem, relativeBase)
			relativeBase += p1
			ip += 2
		case 99:
			if chDone != nil {
				chDone <- 1
			}
			return
		default:
			panic(fmt.Sprint("invalid opcode ", op.Opcode))
		}
	}
}

func LoadProgram(input string) Memory {
	program := make(Memory)

	for i, v := range strings.Split(input, ",") {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}

		program.Set(Value(i), Value(value))
	}

	return program
}

func CopyProgram(program Memory) Memory {
	cp := make(Memory)
	for k, v := range program {
		cp[k] = v
	}

	return cp
}
