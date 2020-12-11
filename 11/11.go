package main

import (
	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/recover"
)

const (
	input = "input.txt"
)

func parse(filename string) [][]rune {
	res := [][]rune{}
	for l := range load.File(filename) {
		res = append(res, []rune(l))
	}
	return res
}

func copy2d(dst [][]rune, src [][]rune) {
	for y := 0; y < min(len(dst), len(src)); y++ {
		copy(dst[y], src[y])
	}
}

func clone(seats [][]rune) [][]rune {
	res := make([][]rune, len(seats))
	for y, row := range seats {
		res[y] = make([]rune, len(row))
	}

	copy2d(res, seats)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func occupied(seats [][]rune, x, y int) int {
	count := 0
	for i := max(y-1, 0); i < min(y+2, len(seats)); i++ {
		for j := max(x-1, 0); j < min(x+2, len(seats[i])); j++ {
			if (i != y || j != x) && seats[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	defer recover.Err(log.Err)

	seats := parse(input)

	// Part 1
	next := clone(seats)
	modified := true
	for modified {
		modified = false
		for y, row := range seats {
			for x, seat := range row {
				if seat == 'L' && occupied(seats, x, y) == 0 {
					next[y][x] = '#'
					modified = true
				} else if seat == '#' && occupied(seats, x, y) >= 4 {
					next[y][x] = 'L'
					modified = true
				}
			}
		}
		copy2d(seats, next)
	}

	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}

	log.Part1(count)
}
