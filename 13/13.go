package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/recover"
)

const (
	input = "input.txt"
)

func parse(filename string) (int, []int) {
	ch := load.File(filename)
	defer func() {
		for range ch {
		}
	}()

	t, err := strconv.Atoi(<-ch)
	if err != nil {
		panic(fmt.Errorf("Atoi: %w", err))
	}

	res := []int{}
	for _, s := range strings.Split(<-ch, ",") {
		if s == "x" {
			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Errorf("Atoi: %w", err))
		}

		res = append(res, n)
	}

	return t, res
}

func main() {
	defer recover.Err(log.Err)

	t, ids := parse(input)

	// Part 1
	min := 999
	minId := 999
	for _, id := range ids {
		rem := t % id
		if rem > 0 {
			rem = id - rem
		}

		if rem < min {
			min = rem
			minId = id
		}
	}
	log.Part1(min * minId)
}
