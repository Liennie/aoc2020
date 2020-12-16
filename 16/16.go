package main

import (
	"regexp"
	"strings"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type iRange struct {
	min, max int
}

func parseRange(s string) iRange {
	split := strings.SplitN(s, "-", 2)
	if len(split) != 2 {
		util.Panic("Invalid range: %q", s)
	}
	return iRange{
		min: util.Atoi(split[0]),
		max: util.Atoi(split[1]),
	}
}

func (r iRange) inRange(n int) bool {
	return r.min <= n && n <= r.max
}

type multiRange []iRange

func (r multiRange) inRange(n int) bool {
	for _, ir := range r {
		if ir.inRange(n) {
			return true
		}
	}
	return false
}

type rules map[string]multiRange

type ticket []int

func parse(filename string) (rules, ticket, []ticket) {
	ch := load.File(filename)

	ruleRe := regexp.MustCompile(`^([\w\s]+?): (\d+-\d+) or (\d+-\d+)$`)
	rs := rules{}
	for l := range ch {
		if len(l) == 0 {
			break
		}

		match := ruleRe.FindStringSubmatch(l)
		if len(match) != 4 {
			util.Panic("No rule match for %q", l)
		}

		rs[match[1]] = multiRange{
			parseRange(match[2]),
			parseRange(match[3]),
		}
	}

	expect := func(s string) {
		if l := <-ch; l != s {
			util.Panic("Expected %q, got %q", s, l)
		}
	}

	expect("your ticket:")
	mt := ticket{}
	for _, n := range strings.Split(<-ch, ",") {
		mt = append(mt, util.Atoi(n))
	}

	expect("")
	expect("nearby tickets:")
	nt := []ticket{}
	for l := range ch {
		t := ticket{}
		for _, n := range strings.Split(l, ",") {
			t = append(t, util.Atoi(n))
		}
		nt = append(nt, t)
	}

	return rs, mt, nt
}

func main() {
	defer util.Recover(log.Err)

	rs, _, nt := parse("input.txt")

	// Part 1
	sum := 0
	for _, t := range nt {
		for _, n := range t {
			valid := false
			for _, r := range rs {
				if r.inRange(n) {
					valid = true
					break
				}
			}
			if !valid {
				sum += n
			}
		}
	}
	log.Part1(sum)
}
