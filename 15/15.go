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

func solve(numbers []int, m int) int {
	i := 0

	prev := map[int]int{}
	for ; i < len(numbers)-1; i++ {
		prev[numbers[i]] = i
	}

	cur := numbers[i]

	for ; i < m-1; i++ {
		if p, ok := prev[cur]; ok {
			prev[cur], cur = i, i-p
		} else {
			prev[cur], cur = i, 0
		}
	}
	return cur
}

func main() {
	defer util.Recover(log.Err)

	numbers := parse(input)

	// Part 1
	log.Part1(solve(numbers, 2020))

	// Part 2
	log.Part2(solve(numbers, 30000000))

}
