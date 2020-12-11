package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/recover"
)

const (
	input = "input.txt"
)

func parse(filename string) []int {
	res := []int{}
	for l := range load.File(filename) {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(fmt.Errorf("Atoi: %w", err))
		}

		res = append(res, i)
	}
	return res
}

func main() {
	defer recover.Err(log.Err)

	numbers := parse(input)

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

	// Part 2
	w := map[int]int{
		numbers[len(numbers)-1]: 1,
	}
	for i := len(numbers) - 2; i >= 0; i-- {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j]-numbers[i] > 3 {
				break
			}
			w[numbers[i]] += w[numbers[j]]
		}
	}
	log.Printf("Part 2: %d", w[0])
}
