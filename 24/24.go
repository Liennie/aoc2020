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
}
