package main

import (
	"log"

	"github.com/liennie/aoc2020/common/asm/op"

	"github.com/liennie/aoc2020/common/asm"
)

const (
	input = "input.txt"
)

type patch struct {
	visited map[int]bool
	patched bool
}

func (p *patch) Pre(i *asm.Instruction, r asm.Registers) error {
	if p.patched {
		return nil
	}

	if !p.visited[i.Line] && (i.Op == op.Nop || i.Op == op.Jmp) {
		p.visited[i.Line] = true
		p.patched = true

		switch i.Op {
		case op.Nop:
			i.Op = op.Jmp
		case op.Jmp:
			i.Op = op.Nop
		}
	}

	return nil
}

func (p *patch) Post(i *asm.Instruction, r asm.Registers) error {
	return nil
}

func main() {
	program, err := asm.LoadFile(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	log.Println(program.Run(asm.SimpleLoopDetection()))
	log.Printf("Part 1: %d", program.Reg().Acc)

	// Part 2
	program.Reset()
	p := &patch{
		visited: map[int]bool{},
	}
	for program.Run(asm.SimpleLoopDetection(), p) != nil {
		if !p.patched {
			log.Println("Nothing patched")
			break
		}

		p.patched = false
		program.Reset()
	}
	log.Printf("Part 2: %d", program.Reg().Acc)
}
