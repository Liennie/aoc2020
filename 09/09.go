package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/liennie/aoc2020/common/load"
)

const (
	input = "input.txt"
)

func parse(filename string) ([]int, error) {
	res := []int{}
	for l := range load.File(filename) {
		if l.Err != nil {
			return nil, l.Err
		}

		i, err := strconv.Atoi(l.Line)
		if err != nil {
			return nil, fmt.Errorf("Atoi: %w", err)
		}

		res = append(res, i)
	}
	return res, nil
}

func main() {
	numbers, err := parse(input)
	if err != nil {
		log.Printf("Parse: %s", err)
	}

	// Part 1
	const preamble = 25
	n := -1
loop1:
	for i := preamble; i < len(numbers); i++ {
		for j := i - preamble; j < i; j++ {
			for k := i - preamble; k < i; k++ {
				if numbers[j] != numbers[k] && numbers[j]+numbers[k] == numbers[i] {
					continue loop1
				}
			}
		}
		n = numbers[i]
		break
	}
	log.Printf("Part 1: %d", n)
}
