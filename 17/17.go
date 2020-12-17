package main

import (
	"fmt"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type point struct {
	x, y, z, w int
}

func (p point) min(o point) point {
	return point{
		util.Min(p.x, o.x),
		util.Min(p.y, o.y),
		util.Min(p.z, o.z),
		util.Min(p.w, o.w),
	}
}

func (p point) max(o point) point {
	return point{
		util.Max(p.x, o.x),
		util.Max(p.y, o.y),
		util.Max(p.z, o.z),
		util.Max(p.w, o.w),
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
				cells[point{x, y, 0, 0}] = true

				max = max.max(point{x, y, 0, 0})
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
				for w := p.w - 1; w <= p.w+1; w++ {
					if (x != p.x || y != p.y || z != p.z || w != p.w) && g.cells[point{x, y, z, w}] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (g *grid) step3() {
	min := point{}
	max := point{}
	cells := map[point]bool{}

	for x := g.min.x - 1; x <= g.max.x+1; x++ {
		for y := g.min.y - 1; y <= g.max.y+1; y++ {
			for z := g.min.z - 1; z <= g.max.z+1; z++ {
				if g.cells[point{x, y, z, 0}] {
					if n := g.neighbors(point{x, y, z, 0}); n == 2 || n == 3 {
						// Cell remains active
						cells[point{x, y, z, 0}] = true
						min = min.min(point{x, y, z, 0})
						max = max.max(point{x, y, z, 0})
					}
				} else {
					if n := g.neighbors(point{x, y, z, 0}); n == 3 {
						// Cell becomes active
						cells[point{x, y, z, 0}] = true
						min = min.min(point{x, y, z, 0})
						max = max.max(point{x, y, z, 0})
					}
				}
			}
		}
	}

	g.min = min
	g.max = max
	g.cells = cells
}

func (g *grid) step4() {
	min := point{}
	max := point{}
	cells := map[point]bool{}

	for x := g.min.x - 1; x <= g.max.x+1; x++ {
		for y := g.min.y - 1; y <= g.max.y+1; y++ {
			for z := g.min.z - 1; z <= g.max.z+1; z++ {
				for w := g.min.w - 1; w <= g.max.w+1; w++ {
					if g.cells[point{x, y, z, w}] {
						if n := g.neighbors(point{x, y, z, w}); n == 2 || n == 3 {
							// Cell remains active
							cells[point{x, y, z, w}] = true
							min = min.min(point{x, y, z, w})
							max = max.max(point{x, y, z, w})
						}
					} else {
						if n := g.neighbors(point{x, y, z, w}); n == 3 {
							// Cell becomes active
							cells[point{x, y, z, w}] = true
							min = min.min(point{x, y, z, w})
							max = max.max(point{x, y, z, w})
						}
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

func (g *grid) print3() {
	for z := g.min.z; z <= g.max.z; z++ {
		log.Printf("z = %d", z)
		for y := g.min.y; y <= g.max.y; y++ {
			for x := g.min.x; x <= g.max.x; x++ {
				if g.cells[point{x, y, z, 0}] {
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
		g.step3()
	}
	log.Part1(len(g.cells))

	// Part 2
	g = og.clone()
	for i := 0; i < 6; i++ {
		g.step4()
	}
	log.Part2(len(g.cells))
}
