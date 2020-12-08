package asm

import (
	"fmt"

	"github.com/liennie/aoc2020/common/asm/op"
)

type Instruction struct {
	Line int
	Op   op.Code
	Args []int
}

func newInstruction(line int, o op.Code, args ...int) (*Instruction, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Instruction %s needs 1 arg", o)
	}

	return &Instruction{
		Line: line,
		Op:   o,
		Args: args,
	}, nil
}

func (i Instruction) Exec(p *Program) {
	exec[i.Op](p, i.Args...)
}

var exec = map[op.Code]func(p *Program, args ...int){
	// Nop
	op.Nop: func(p *Program, args ...int) {},

	// Acc
	op.Acc: func(p *Program, args ...int) {
		p.reg.Acc += args[0]
	},

	// Jmp
	op.Jmp: func(p *Program, args ...int) {
		p.reg.Ip += args[0] - 1
	},
}
