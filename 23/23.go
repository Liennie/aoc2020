package main

import (
	"container/ring"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type cups struct {
	cur *ring.Ring
	dir map[int]*ring.Ring
	max int
	len int
}

func (c *cups) clone() *cups {
	res := &cups{
		cur: ring.New(c.len),
		dir: map[int]*ring.Ring{},
		max: c.max,
		len: c.len,
	}
	c.cur.Do(func(v interface{}) {
		res.cur.Value = v
		res.dir[v.(int)] = res.cur
		res.cur = res.cur.Next()
	})
	return res
}

func (c *cups) expand(to int) *cups {
	x := ring.New(to - c.len)
	for i := 0; i < to-c.len; i++ {
		x.Value = c.max + i + 1
		c.dir[c.max+i+1] = x
		x = x.Next()
	}
	c.cur.Prev().Link(x)
	c.max += to - c.len
	c.len = to
	return c
}

func (c *cups) move() {
	rem := c.cur.Unlink(3)

	ex := map[int]bool{}
	rem.Do(func(v interface{}) {
		ex[v.(int)] = true
	})

	var dst *ring.Ring
	m := c.cur.Value.(int) - 1
	for i := 1; i < c.max; i++ {
		if !ex[m] {
			if d, ok := c.dir[m]; ok {
				dst = d
				break
			}
		}
		m--
		if m <= 0 {
			m = c.max
		}
	}

	if dst == nil {
		util.Panic("Destination cup not found")
	}

	dst.Link(rem)
	c.cur = c.cur.Next()
}

func (c *cups) one() *ring.Ring { // to rule them all
	if one, ok := c.dir[1]; ok {
		return one
	}
	util.Panic("1 not found")
	return nil
}

func (c *cups) order() int {
	r := c.one().Next()
	o := 0
	for i := 1; i < c.len; i++ {
		o *= 10
		o += r.Value.(int)
		r = r.Next()
	}
	return o
}

func parse(filename string) *cups {
	res := &cups{}
	ns := []int{}
	for l := range load.File(filename) {
		c := util.Atoi(l)
		if c > res.max {
			res.max = c
		}
		if c <= 0 {
			util.Panic("No")
		}
		ns = append(ns, c)
	}
	res.len = len(ns)
	res.dir = map[int]*ring.Ring{}
	res.cur = ring.New(res.len)
	for i := 0; i < res.len; i++ {
		res.cur.Value = ns[i]
		res.dir[ns[i]] = res.cur
		res.cur = res.cur.Next()
	}
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

	// Part 2
	c = oc.clone().expand(1000000)
	for i := 0; i < 10000000; i++ {
		c.move()
	}
	one := c.one()
	log.Part2(one.Next().Value.(int) * one.Next().Next().Value.(int))
}
