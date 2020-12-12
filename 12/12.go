package main

import (
	"fmt"
	"strconv"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/recover"
)

const (
	input = "input.txt"
)

type nav struct {
	x, y, f, t int
}

func parse(filename string) []nav {
	res := []nav{}
	for l := range load.File(filename) {
		dir := nav{}
		switch l[0] {
		case 'N':
			dir = nav{0, 1, 0, 0}
		case 'E':
			dir = nav{1, 0, 0, 0}
		case 'S':
			dir = nav{0, -1, 0, 0}
		case 'W':
			dir = nav{-1, 0, 0, 0}
		case 'L':
			dir = nav{0, 0, 0, 1}
		case 'R':
			dir = nav{0, 0, 0, -1}
		case 'F':
			dir = nav{0, 0, 1, 0}
		default:
			panic(fmt.Errorf("Invalid direction: %q", l[:1]))
		}

		i, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(fmt.Errorf("Atoi: %w", err))
		}

		dir.x *= i
		dir.y *= i
		dir.f *= i
		dir.t *= i

		res = append(res, dir)
	}
	return res
}

func sin(a int) int {
	switch a {
	case 0:
		return 0
	case 90:
		return 1
	case 180:
		return 0
	case 270:
		return -1
	default:
		panic(fmt.Errorf("Sin: %d", a))
	}
}

func cos(a int) int {
	switch a {
	case 0:
		return 1
	case 90:
		return 0
	case 180:
		return -1
	case 270:
		return 0
	default:
		panic(fmt.Errorf("Cos: %d", a))
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	defer recover.Err(log.Err)

	navigation := parse(input)

	// Part 1
	ship := nav{0, 0, 0, 0}
	for _, dir := range navigation {
		if dir.t%90 != 0 {
			panic("Noooo")
		}
		ship.t += dir.t
		ship.t %= 360
		if ship.t < 0 {
			ship.t += 360
		}

		ship.x += dir.x + dir.f*cos(ship.t)
		ship.y += dir.y + dir.f*sin(ship.t)
	}
	log.Part1(abs(ship.x) + abs(ship.y))
}
