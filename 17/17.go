package main

import (
	"fmt"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type point struct {
	x, y, z int
}

func (p point) min(o point) point {
	return point{
		util.Min(p.x, o.x),
		util.Min(p.y, o.y),
		util.Min(p.z, o.z),
	}
}

func (p point) max(o point) point {
	return point{
		util.Max(p.x, o.x),
		util.Max(p.y, o.y),
		util.Max(p.z, o.z),
	}
}

type grid struct {
	min, max point
	cells    map[point]bool
}

func parse(filename string) grid {
	min := point{}
	max := point{}
	cells := map[point]bool{}

	y := 0
	for l := range load.File(filename) {
		for x, c := range l {
			if c == '#' {
				cells[point{x, y, 0}] = true

				max = max.max(point{x, y, 0})
			}
		}
		y++
	}
	return grid{
		min:   min,
		max:   max,
		cells: cells,
	}
}

func (g *grid) neighbors(p point) int {
	count := 0
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				if (x != p.x || y != p.y || z != p.z) && g.cells[point{x, y, z}] {
					count++
				}
			}
		}
	}
	return count
}

func (g *grid) step() {
	min := point{}
	max := point{}
	cells := map[point]bool{}

	for x := g.min.x - 1; x <= g.max.x+1; x++ {
		for y := g.min.y - 1; y <= g.max.y+1; y++ {
			for z := g.min.z - 1; z <= g.max.z+1; z++ {
				if g.cells[point{x, y, z}] {
					if n := g.neighbors(point{x, y, z}); n == 2 || n == 3 {
						// Cell remains active
						cells[point{x, y, z}] = true
						min = min.min(point{x, y, z})
						max = max.max(point{x, y, z})
					}
				} else {
					if n := g.neighbors(point{x, y, z}); n == 3 {
						// Cell becomes active
						cells[point{x, y, z}] = true
						min = min.min(point{x, y, z})
						max = max.max(point{x, y, z})
					}
				}
			}
		}
	}

	g.min = min
	g.max = max
	g.cells = cells
}

func (g *grid) clone() grid {
	new := grid{
		min:   g.min,
		max:   g.max,
		cells: map[point]bool{},
	}
	for k, v := range g.cells {
		new.cells[k] = v
	}
	return new
}

func (g *grid) print() {
	for z := g.min.z; z <= g.max.z; z++ {
		log.Printf("z = %d", z)
		for y := g.min.y; y <= g.max.y; y++ {
			for x := g.min.x; x <= g.max.x; x++ {
				if g.cells[point{x, y, z}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	log.Println(len(g.cells))
}

func main() {
	defer util.Recover(log.Err)

	og := parse("input.txt")

	// Part 1
	g := og.clone()
	for i := 0; i < 6; i++ {
		g.step()
	}
	log.Part1(len(g.cells))
}
