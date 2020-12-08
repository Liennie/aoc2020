package asm

import (
	"fmt"

	"github.com/liennie/aoc2020/common/asm/op"
)

type Instruction interface {
	Debug() DebugInfo
	Exec(*Program)
}

func newInstruction(line int, o op.Code, args ...int) (Instruction, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Instruction %s needs 1 arg", o)
	}

	d := DebugInfo{
		Line: line,
		Op:   o,
		Args: args,
	}

	switch o {
	case op.Nop:
		return nop{d}, nil
	case op.Acc:
		return acc{d}, nil
	case op.Jmp:
		return jmp{d}, nil
	}

	return nil, fmt.Errorf("Invalid instruction %s", o)
}

type DebugInfo struct {
	Line int
	Op   op.Code
	Args []int
}

func (d DebugInfo) Debug() DebugInfo {
	return d
}

type nop struct {
	DebugInfo
}

func (i nop) Exec(p *Program) {}

type acc struct {
	DebugInfo
}

func (i acc) Exec(p *Program) {
	p.reg.Acc += i.Args[0]
}

type jmp struct {
	DebugInfo
}

func (i jmp) Exec(p *Program) {
	p.reg.Ip += i.Args[0] - 1
}
