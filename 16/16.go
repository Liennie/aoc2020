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

func (r rules) valid(n int) bool {
	for _, mr := range r {
		if mr.inRange(n) {
			return true
		}
	}
	return false
}

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
		if len(t) != len(mt) {
			util.Panic("Invalid ticket length: %d; should be %d", len(t), len(mt))
		}
		nt = append(nt, t)
	}

	return rs, mt, nt
}

func main() {
	defer util.Recover(log.Err)

	rs, mt, nt := parse("input.txt")

	// Part 1
	sum := 0
	for _, t := range nt {
		for _, n := range t {
			if !rs.valid(n) {
				sum += n
			}
		}
	}
	log.Part1(sum)

	// Part 2
	nnt := []ticket{}
	for _, t := range nt {
		valid := true
		for _, n := range t {
			if !rs.valid(n) {
				valid = false
				break
			}
		}
		if valid {
			nnt = append(nnt, t)
		}
	}
	nt = nnt

	fs := map[string]map[int]bool{}
	for r := range rs {
		fs[r] = map[int]bool{}
		for i := range mt {
			fs[r][i] = true
		}
	}

	for _, t := range append(nt, mt) {
		for i, n := range t {
			for r, mr := range rs {
				if !mr.inRange(n) {
					delete(fs[r], i)
				}
			}
		}
	}

	res := map[string]int{}
	for len(fs) > 0 {
		for f, s := range fs {
			if len(s) == 1 {
				for k := range s {
					res[f] = k
				}
				for _, s := range fs {
					delete(s, res[f])
				}
				delete(fs, f)
			}
		}
	}

	prod := 1
	for f, i := range res {
		if strings.HasPrefix(f, "departure") {
			prod *= mt[i]
		}
	}
	log.Part2(prod)
}
