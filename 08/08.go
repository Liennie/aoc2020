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
	log.Println(program.Run(asm.SimpleLoopDetection()))
	log.Printf("Part 1: %d", program.Reg().Acc)
}
