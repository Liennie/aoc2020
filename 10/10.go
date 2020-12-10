package main

import (
	"fmt"
	"log"
	"sort"
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

	sort.Ints(numbers)

	numbers = append([]int{0}, numbers...)
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	// Part 1
	d1 := 0
	d3 := 0
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff == 1 {
			d1++
		} else if diff == 3 {
			d3++
		} else if diff > 3 {
			panic("Vegeta, what does the scouter say about his power level?")
		}
	}
	log.Printf("Part 1: %d", d1*d3)
}
