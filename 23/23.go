package main

import (
	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type cups struct {
	cups []int
	cur  int
	max  int
	len  int
}

func (c cups) clone() cups {
	res := cups{
		cups: make([]int, len(c.cups)),
		cur:  c.cur,
		max:  c.max,
		len:  c.len,
	}
	copy(res.cups, c.cups)
	return res
}

func (c *cups) move() {
	rem := [3]int{}
	for i := range rem {
		rem[i] = c.cups[(c.cur+i+1)%c.len]
	}

	dst := 0
	min := c.max
	for i := 4; i < c.len; i++ {
		d := (c.cur + i) % c.len
		m := util.Mod(c.cups[c.cur]-c.cups[d], c.max)
		if m < min {
			dst = d
			min = m
		}
	}

	cnt := util.Mod(dst-(c.cur+3), c.len)
	for i := 0; i < cnt; i++ {
		c.cups[(c.cur+i+1)%c.len] = c.cups[(c.cur+i+4)%c.len]
	}

	for i := range rem {
		c.cups[(c.cur+cnt+i+1)%c.len] = rem[i]
	}

	c.cur++
	c.cur %= c.len
}

func (c *cups) order() int {
	s := 0
	for ; s < c.len; s++ {
		if c.cups[s] == 1 {
			break
		}
	}

	o := 0
	for i := 1; i < c.len; i++ {
		o *= 10
		o += c.cups[(s+i)%c.len]
	}

	return o
}

func parse(filename string) cups {
	res := cups{}
	for l := range load.File(filename) {
		c := util.Atoi(l)
		if c > res.max {
			res.max = c
		}
		if c <= 0 {
			util.Panic("No")
		}
		res.cups = append(res.cups, c)
	}
	res.len = len(res.cups)
	return res
}

func main() {
	defer util.Recover(log.Err)

	oc := parse("input.txt")

	// Part 1
	c := oc.clone()
	for i := 0; i < 100; i++ {
		c.move()
	}
	log.Part1(c.order())
}
