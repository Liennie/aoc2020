package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	input = "input.txt"
)

var bagRe = regexp.MustCompile(`^([\w ]+?) bags contain (\d [\w ]+? bags?(?:, \d [\w ]+? bags?)*|no other bags).$`)
var childRe = regexp.MustCompile(`^(\d) ([\w ]+?) bags?$`)

type bag struct {
	color    string
	parents  map[*bag]int
	children map[*bag]int
}

func (b *bag) childrenCount() int {
	res := 0
	for child, count := range b.children {
		res += child.childrenCount()*count + count
	}
	return res
}

type rules map[string]*bag

func (r rules) get(color string) *bag {
	if _, ok := r[color]; !ok {
		r[color] = &bag{
			color:    color,
			parents:  map[*bag]int{},
			children: map[*bag]int{},
		}
	}
	return r[color]
}

func (r rules) addChild(parent string, count int, child string) {
	parentBag := r.get(parent)
	childBag := r.get(child)

	parentBag.children[childBag] = count
	childBag.parents[parentBag] = count
}

func load(filename string) (map[string]*bag, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	res := rules{}
	for {
		l, err := r.ReadString('\n')
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			match := bagRe.FindStringSubmatch(l)
			if len(match) != 3 {
				return nil, fmt.Errorf("No match for %q, %v", l, match)
			}

			color := match[1]

			if match[2] == "no other bags" {
				res.get(color)
			} else {
				for _, child := range strings.Split(match[2], ", ") {
					match := childRe.FindStringSubmatch(child)
					if len(match) != 3 {
						return nil, fmt.Errorf("No match for %q, %v", child, match)
					}

					count, err := strconv.Atoi(match[1])
					if err != nil {
						return nil, fmt.Errorf("Atoi: %w", err)
					}

					res.addChild(color, count, match[2])
				}
			}
		}
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("ReadString: %w", err)
			}
			break
		}
	}

	return res, nil
}

func any(group []map[rune]struct{}) map[rune]struct{} {
	res := map[rune]struct{}{}
	for _, m := range group {
		for r := range m {
			res[r] = struct{}{}
		}
	}
	return res
}

func all(group []map[rune]struct{}) map[rune]struct{} {
	res := map[rune]struct{}{}
	for r := 'a'; r <= 'z'; r++ {
		res[r] = struct{}{}
	}
	for _, m := range group {
		for r := range res {
			if _, ok := m[r]; !ok {
				delete(res, r)
			}
		}
	}
	return res
}

func main() {
	rules, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	res := map[string]bool{}
	toSearch := []*bag{rules["shiny gold"]}
	for len(toSearch) > 0 {
		bag := toSearch[0]
		toSearch = toSearch[1:]

		res[bag.color] = true
		for parent := range bag.parents {
			toSearch = append(toSearch, parent)
		}
	}
	log.Printf("Part 1: %d", len(res)-1)

	// Part 2
	log.Printf("Part 2: %d", rules["shiny gold"].childrenCount())
}
