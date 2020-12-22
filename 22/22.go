package main

import (
	"fmt"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

func parse(filename string) ([]int, []int) {
	ch := load.File(filename)

	expect := func(s string) {
		if l := <-ch; l != s {
			util.Panic("Expected %q, got %q", s, l)
		}
	}

	expect("Player 1:")
	p1 := []int{}
	for l := range ch {
		if l == "" {
			break
		}
		p1 = append(p1, util.Atoi(l))
	}

	expect("Player 2:")
	p2 := []int{}
	for l := range ch {
		p2 = append(p2, util.Atoi(l))
	}

	return p1, p2
}

func combat(op1, op2 []int) ([]int, []int) {
	p1 := make([]int, len(op1))
	p2 := make([]int, len(op2))
	copy(p1, op1)
	copy(p2, op2)

	for len(p1) > 0 && len(p2) > 0 {
		if p1[0] > p2[0] {
			p1 = append(p1[1:], p1[0], p2[0])
			p2 = p2[1:]
		} else if p1[0] < p2[0] {
			p2 = append(p2[1:], p2[0], p1[0])
			p1 = p1[1:]
		} else {
			util.Panic("Cards are equal")
		}
	}

	return p1, p2
}

func recursiveCombat(op1, op2 []int) ([]int, []int) {
	p1 := make([]int, len(op1))
	p2 := make([]int, len(op2))
	copy(p1, op1)
	copy(p2, op2)

	prev := map[string]bool{}
	prevKey := func() string {
		return fmt.Sprint(p1, p2)
	}

	for len(p1) > 0 && len(p2) > 0 {
		if prev[prevKey()] {
			return append(p1, p2...), []int{}
		}
		prev[prevKey()] = true

		p1c := p1[0]
		p2c := p2[0]
		p1 = p1[1:]
		p2 = p2[1:]

		if p1c <= len(p1) && p2c <= len(p2) {
			p1r, p2r := recursiveCombat(p1[:p1c], p2[:p2c])
			if len(p1r) > len(p2r) {
				p1 = append(p1, p1c, p2c)
			} else {
				p2 = append(p2, p2c, p1c)
			}
		} else if p1c > p2c {
			p1 = append(p1, p1c, p2c)
		} else if p2c > p1c {
			p2 = append(p2, p2c, p1c)
		} else {
			util.Panic("Cards are equal")
		}
	}

	return p1, p2
}

func score(w []int) int {
	score := 0
	l := len(w)
	for i, c := range w {
		score += (l - i) * c
	}
	return score
}

func main() {
	defer util.Recover(log.Err)

	op1, op2 := parse("input.txt")

	// Part 1
	p1, p2 := combat(op1, op2)
	if len(p1) > len(p2) {
		log.Part1(score(p1))
	} else {
		log.Part1(score(p2))
	}

	// Part 2
	p1, p2 = recursiveCombat(op1, op2)
	if len(p1) > len(p2) {
		log.Part2(score(p1))
	} else {
		log.Part2(score(p2))
	}
}
