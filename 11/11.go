package main

import (
	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
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
	for y := 0; y < util.Min(len(dst), len(src)); y++ {
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

func occupied(seats [][]rune, x, y int) int {
	count := 0
	for i := util.Max(y-1, 0); i < util.Min(y+2, len(seats)); i++ {
		for j := util.Max(x-1, 0); j < util.Min(x+2, len(seats[i])); j++ {
			if (i != y || j != x) && seats[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func occupied2(seats [][]rune, x, y int) int {
	count := 0
	for yd := -1; yd <= 1; yd++ {
		for xd := -1; xd <= 1; xd++ {
			if yd != 0 || xd != 0 {
				for i, j := y+yd, x+xd; i >= 0 && i < len(seats) && j >= 0 && j < len(seats[i]); i, j = i+yd, j+xd {
					if seats[i][j] == '#' {
						count++
						break
					} else if seats[i][j] == 'L' {
						break
					}
				}
			}
		}
	}
	return count
}

func count(seats [][]rune, s rune) int {
	c := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == s {
				c++
			}
		}
	}
	return c
}

func foreach(seats [][]rune, f func(r rune, x, y int)) {
	for y, row := range seats {
		for x, seat := range row {
			f(seat, x, y)
		}
	}
}

func main() {
	defer util.Recover(log.Err)

	start := parse(input)

	// Part 1
	seats := clone(start)
	next := clone(start)
	modified := true
	for modified {
		modified = false
		foreach(seats, func(r rune, x, y int) {
			if r == 'L' && occupied(seats, x, y) == 0 {
				next[y][x] = '#'
				modified = true
			} else if r == '#' && occupied(seats, x, y) >= 4 {
				next[y][x] = 'L'
				modified = true
			}
		})
		copy2d(seats, next)
	}
	log.Part1(count(seats, '#'))

	// Part 2
	copy2d(seats, start)
	copy2d(next, start)
	modified = true
	for modified {
		modified = false
		foreach(seats, func(r rune, x, y int) {
			if r == 'L' && occupied2(seats, x, y) == 0 {
				next[y][x] = '#'
				modified = true
			} else if r == '#' && occupied2(seats, x, y) >= 5 {
				next[y][x] = 'L'
				modified = true
			}
		})
		copy2d(seats, next)
	}
	log.Part2(count(seats, '#'))
}
