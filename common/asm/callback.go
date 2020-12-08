package asm

import "fmt"

type Callback interface {
	Pre(DebugInfo, Registers) error
	Post(DebugInfo, Registers) error
}

type simpleLoopDetection struct {
	visited map[int]bool
}

func SimpleLoopDetection() Callback {
	return &simpleLoopDetection{
		visited: map[int]bool{},
	}
}

func (c *simpleLoopDetection) Pre(debug DebugInfo, reg Registers) error {
	if c.visited[debug.Line] {
		return fmt.Errorf("Loop detected on line %d", debug.Line)
	}
	c.visited[debug.Line] = true
	return nil
}

func (c *simpleLoopDetection) Post(debug DebugInfo, reg Registers) error {
	return nil
}
