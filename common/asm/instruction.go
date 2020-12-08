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

func (i Instruction) Exec(p *Program) {
	exec[i.Op](p, i.Args)
}

var exec = map[op.Code]func(p *Program, args [argc]int){
	// Nop
	op.Nop: func(p *Program, args [argc]int) {},

	// Acc
	op.Acc: func(p *Program, args [argc]int) {
		p.reg.Acc += args[0]
	},

	// Jmp
	op.Jmp: func(p *Program, args [argc]int) {
		p.reg.Ip += args[0] - 1
	},
}
