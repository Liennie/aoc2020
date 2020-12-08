package main

import (
	"log"

	"github.com/liennie/aoc2020/common/asm"
)

const (
	input = "input.txt"
)

func main() {
	program, err := asm.LoadFile(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	visited := map[int]bool{}
	debug := program.Next()
	for !visited[debug.Line] {
		visited[debug.Line] = true
		program.Step()
		debug = program.Next()
	}
	log.Printf("Part 1: %d", program.Reg().Acc)
}
