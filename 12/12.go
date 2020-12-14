package main

import (
	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
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
			util.Panic("Invalid direction: %q", l[:1])
		}

		i := util.Atoi(l[1:])

		dir.x *= i
		dir.y *= i
		dir.f *= i
		dir.t *= i

		res = append(res, dir)
	}
	return res
}

func sin(a int) int {
	if a < 0 {
		a += 360
	}
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
		util.Panic("Sin: %d", a)
	}
	return 0
}

func cos(a int) int {
	if a < 0 {
		a += 360
	}
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
		util.Panic("Cos: %d", a)
	}
	return 0
}

func rot(x, y, a int) (int, int) {
	if a < 0 {
		a += 360
	}
	switch a {
	case 0:
		return x, y
	case 90:
		return -y, x
	case 180:
		return -x, -y
	case 270:
		return y, -x
	default:
		util.Panic("Rot: %d", a)
	}
	return 0, 0
}

func main() {
	defer util.Recover(log.Err)

	navigation := parse(input)

	// Part 1
	ship := nav{0, 0, 0, 0}
	for _, dir := range navigation {
		if dir.t%90 != 0 {
			panic("Noooo")
		}

		ship.t += dir.t
		ship.t %= 360
		ship.x += dir.x + dir.f*cos(ship.t)
		ship.y += dir.y + dir.f*sin(ship.t)
	}
	log.Part1(util.Abs(ship.x) + util.Abs(ship.y))

	// Part2
	ship = nav{0, 0, 0, 0}
	way := nav{10, 1, 0, 0}
	for _, dir := range navigation {
		if dir.t%90 != 0 {
			panic("Noooo")
		}

		way.x += dir.x
		way.y += dir.y
		way.x, way.y = rot(way.x, way.y, dir.t)
		ship.x += dir.f * way.x
		ship.y += dir.f * way.y
	}
	log.Part2(util.Abs(ship.x) + util.Abs(ship.y))
}
