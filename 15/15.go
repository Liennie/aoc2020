package main

import (
	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

const (
	input = "input.txt"
)

func parse(filename string) []int {
	res := []int{}
	for l := range load.File(filename) {
		res = append(res, util.Atoi(l))
	}
	return res
}

func main() {
	defer util.Recover(log.Err)

	numbers := parse(input)

	// Part 1
	for i := len(numbers) - 1; i < 2019; i++ {
		n := numbers[i]
		next := 0
		for j := i - 1; j >= 0; j-- {
			if numbers[j] == n {
				next = i - j
				break
			}
		}
		numbers = append(numbers, next)
	}
	log.Part1(numbers[2019])
}
