package main

import (
	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type point struct {
	x, y int
}

func (p point) nw() point {
	return point{p.x, p.y - 1}
}

func (p point) ne() point {
	return point{p.x + 1, p.y - 1}
}

func (p point) e() point {
	return point{p.x + 1, p.y}
}

func (p point) se() point {
	return point{p.x, p.y + 1}
}

func (p point) sw() point {
	return point{p.x - 1, p.y + 1}
}

func (p point) w() point {
	return point{p.x - 1, p.y}
}

func parse(filename string) []point {
	res := []point{}
	for l := range load.File(filename) {
		p := point{}

		n := false
		s := false
		for _, c := range l {
			switch c {
			case 'e':
				if n {
					p = p.ne()
					n = false
				} else if s {
					p = p.se()
					s = false
				} else {
					p = p.e()
				}
			case 'w':
				if n {
					p = p.nw()
					n = false
				} else if s {
					p = p.sw()
					s = false
				} else {
					p = p.w()
				}
			case 'n':
				if n || s {
					util.Panic("Invalid n")
				}
				n = true
			case 's':
				if n || s {
					util.Panic("Invalid s")
				}
				s = true
			default:
				util.Panic("Invalid dir")
			}
		}

		res = append(res, p)
	}
	return res
}

func next(grid map[point]bool) map[point]bool {
	n := map[point]int{}
	for p := range grid {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x != y {
					n[point{p.x + x, p.y + y}]++
				}
			}
		}
	}

	next := map[point]bool{}
	for p, c := range n {
		if grid[p] {
			if c == 1 || c == 2 {
				next[p] = true
			}
		} else {
			if c == 2 {
				next[p] = true
			}
		}
	}

	return next
}

func main() {
	defer util.Recover(log.Err)

	ps := parse("input.txt")

	// Part 1
	grid := map[point]bool{}
	for _, p := range ps {
		if grid[p] {
			delete(grid, p)
		} else {
			grid[p] = true
		}
	}
	log.Part1(len(grid))

	// Part 2
	for i := 0; i < 100; i++ {
		grid = next(grid)
	}
	log.Part2(len(grid))
}
