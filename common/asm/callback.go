package asm

import (
	"fmt"
)

type Callback interface {
	Pre(*Instruction, Registers) error
	Post(*Instruction, Registers) error
}

type simpleLoopDetection struct {
	visited map[int]bool
}

func SimpleLoopDetection() Callback {
	return &simpleLoopDetection{
		visited: map[int]bool{},
	}
}

func (c *simpleLoopDetection) Pre(i *Instruction, reg Registers) error {
	if c.visited[reg.Ip] {
		return fmt.Errorf("Loop detected on line %d", i.Line)
	}
	c.visited[reg.Ip] = true
	return nil
}

func (c *simpleLoopDetection) Post(i *Instruction, reg Registers) error {
	return nil
}
