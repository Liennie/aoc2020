package asm

import (
	"fmt"

	"github.com/liennie/aoc2020/common/asm/op"
)

const (
	argc = 1
)

type Instruction struct {
	Line int
	Op   op.Code
	Args [argc]int
}

func newInstruction(line int, o op.Code, args ...int) (*Instruction, error) {
	if len(args) != argc {
		return nil, fmt.Errorf("Instruction %s needs %d args", o, argc)
	}

	i := &Instruction{
		Line: line,
		Op:   o,
	}

	copy(i.Args[:], args[:argc])

	return i, nil
}

func (p *Program) Exec(i Instruction) {
	switch i.Op {
	case op.Nop:
		// Do nothing

	case op.Acc:
		p.reg.Acc += i.Args[0]

	case op.Jmp:
		p.reg.Ip += i.Args[0] - 1

	default:
		panic(fmt.Errorf("Invalid instruction %s", i.Op))
	}
}
