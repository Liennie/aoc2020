package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/recover"
	"github.com/liennie/aoc2020/common/util"
)

const (
	input = "input.txt"
)

func parse(filename string) (int, map[int]int) {
	ch := load.File(filename)
	defer func() {
		for range ch {
		}
	}()

	t, err := strconv.Atoi(<-ch)
	if err != nil {
		panic(fmt.Errorf("Atoi: %w", err))
	}

	res := map[int]int{}
	for i, s := range strings.Split(<-ch, ",") {
		if s == "x" {
			continue
		}

		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Errorf("Atoi: %w", err))
		}

		res[i] = n
	}

	return t, res
}

func main() {
	defer recover.Err(log.Err)

	t, ids := parse(input)

	// Part 1
	min := 999
	minId := 999
	for _, id := range ids {
		rem := t % id
		if rem > 0 {
			rem = id - rem
		}

		if rem < min {
			min = rem
			minId = id
		}
	}
	log.Part1(min * minId)

	// Part 2
	co, cm := 0, 1
	for o, m := range ids {
		if cm > m {
			co, cm, o, m = o, m, co, cm
		}

		mod := map[int]bool{}
		for i := 1; ; i++ {
			oo := (m * i) % cm

			if mod[oo] {
				panic(fmt.Errorf("Oh noes: %v", map[string]interface{}{"co": co, "cm": cm, "o": o, "m": m, "i": i, "mod": mod}))
			}
			mod[oo] = true

			if oo == util.Mod(o-co, cm) {
				cm = util.LCM(m, cm)
				co = cm - (m*i - o)
				break
			}
		}
	}
	log.Part2(cm - co)
}
