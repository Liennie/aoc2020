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
	log.Part1(n)

	// Part 2
	min, max := -1, -1
loop2:
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := 0
			mn, mx := numbers[i], numbers[i]
			for k := i; k <= j; k++ {
				nk := numbers[k]

				sum += nk

				if nk < mn {
					mn = nk
				} else if nk > mx {
					mx = nk
				}
			}
			if sum == n {
				min = mn
				max = mx
				break loop2
			}
		}
	}
	log.Part2(min + max)
}
