package main

import (
	"sort"
	"strings"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

const (
	top = iota
	right
	bottom
	left
)

type food struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func parse(filename string) []food {
	res := []food{}
	for l := range load.File(filename) {
		split := strings.Split(l, " (contains ")
		if len(split) != 2 {
			util.Panic("Split len %d", len(split))
		}

		f := food{
			ingredients: map[string]bool{},
			allergens:   map[string]bool{},
		}
		for _, i := range strings.Split(split[0], " ") {
			f.ingredients[i] = true
		}
		for _, a := range strings.Split(strings.TrimRight(split[1], ")"), ", ") {
			f.allergens[a] = true
		}

		res = append(res, f)
	}

	return res
}

func main() {
	defer util.Recover(log.Err)

	foods := parse("input.txt")

	// Part 1
	ingredients := map[string]bool{}
	for _, f := range foods {
		for i := range f.ingredients {
			ingredients[i] = true
		}
	}

	options := map[string]map[string]bool{}
	for _, f := range foods {
		for a := range f.allergens {
			if options[a] == nil {
				options[a] = map[string]bool{}
				for i := range f.ingredients {
					options[a][i] = true
				}
			} else {
				for i := range ingredients {
					if !f.ingredients[i] {
						delete(options[a], i)
					}
				}
			}
		}
	}

	allergenIngredients := map[string]string{}
	for len(options) > 0 {
		for a, is := range options {
			if len(is) == 1 {
				for i := range is {
					allergenIngredients[a] = i
				}
				for _, is := range options {
					delete(is, allergenIngredients[a])
				}
				delete(options, a)
			}
		}
	}
	ingredientAllergens := map[string]string{}
	for a, i := range allergenIngredients {
		ingredientAllergens[i] = a
	}

	count := 0
	for _, f := range foods {
		for i := range f.ingredients {
			if _, ok := ingredientAllergens[i]; !ok {
				count++
			}
		}
	}
	log.Part1(count)

	// Part 2
	keys := []string{}
	for a := range allergenIngredients {
		keys = append(keys, a)
	}
	sort.Strings(keys)

	b := &strings.Builder{}
	for i, k := range keys {
		if i > 0 {
			b.WriteRune(',')
		}
		b.WriteString(allergenIngredients[k])
	}
	log.Part2(b.String())
}
