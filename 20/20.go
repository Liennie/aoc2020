package main

import (
	"math/bits"
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

type border struct {
	id      uint
	flipped bool
}

type tile struct {
	id      int
	rawData [][]rune
	borders [4]border
}

func newBorder(id uint, len int) border {
	if fid := bits.Reverse(id) >> (bits.UintSize - len); fid < id {
		return border{
			id:      fid,
			flipped: true,
		}
	} else if fid == id {
		util.Panic("Fuck! Palindromic border id %d", id)
	}
	return border{
		id:      id,
		flipped: false,
	}
}

func tileBorders(rawData [][]rune) [4]border {
	res := [4]border{}

	// top
	id := uint(0)
	ln := len(rawData[0])
	for i, r := range rawData[0] {
		if r == '#' {
			id |= 1 << (ln - i - 1)
		} else if r != '.' {
			util.Panic("Invalid rune %q", string(r))
		}
	}
	res[top] = newBorder(id, ln)

	// right
	id = uint(0)
	ln = len(rawData)
	for i, l := range rawData {
		r := l[len(l)-1]
		if r == '#' {
			id |= 1 << (ln - i - 1)
		} else if r != '.' {
			util.Panic("Invalid rune %q", string(r))
		}
	}
	res[right] = newBorder(id, ln)

	// bottom
	id = uint(0)
	ln = len(rawData[len(rawData)-1])
	for i, r := range rawData[len(rawData)-1] {
		if r == '#' {
			id |= 1 << i
		} else if r != '.' {
			util.Panic("Invalid rune %q", string(r))
		}
	}
	res[bottom] = newBorder(id, ln)

	// left
	id = uint(0)
	ln = len(rawData)
	for i, l := range rawData {
		r := l[0]
		if r == '#' {
			id |= 1 << i
		} else if r != '.' {
			util.Panic("Invalid rune %q", string(r))
		}
	}
	res[left] = newBorder(id, ln)

	return res
}

func parse(filename string) (map[int]tile, map[uint][]int) {
	tiles := map[int]tile{}
	t := tile{}
	el := 0
	for l := range load.File(filename) {
		if l == "" {
			if t.id != 0 {
				if len(t.rawData) != el {
					util.Panic("Invalid tile height")
				}
				tiles[t.id] = t
				t = tile{}
			}
		} else if strings.HasPrefix(l, "Tile") {
			t.id = util.Atoi(strings.TrimPrefix(strings.TrimSuffix(l, ":"), "Tile "))
		} else {
			if el == 0 {
				el = len(l)
			} else if len(l) != el {
				util.Panic("Invalid tile width")
			}
			t.rawData = append(t.rawData, []rune(l))
		}
	}
	if t.id != 0 {
		if len(t.rawData) != el {
			util.Panic("Invalid tile height")
		}
		tiles[t.id] = t
	}

	borders := map[uint][]int{}
	for k, t := range tiles {
		t.borders = tileBorders(t.rawData)
		tiles[k] = t
		for _, b := range t.borders {
			borders[b.id] = append(borders[b.id], t.id)
		}
	}

	return tiles, borders
}

func main() {
	defer util.Recover(log.Err)

	tiles, borders := parse("input.txt")

	for bid, ts := range borders {
		if len(ts) > 2 {
			util.Panic("Fuck! %d tiles have border with id %d", len(ts), bid)
		}
	}

	// Part 1
	prod := 1
	c := 0
	for _, t := range tiles {
		e := 0
		for _, b := range t.borders {
			if len(borders[b.id]) == 1 {
				e++
			}
		}
		if e == 2 {
			prod *= t.id
			c++
		} else if e > 2 {
			util.Panic("WTF! Tile %d has %d edges", t.id, e)
		}
	}
	if c != 4 {
		util.Panic("WTF! Image has %d corners", c)
	}
	log.Part1(prod)
}
