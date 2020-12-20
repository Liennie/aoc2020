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

func normalizeDir(dir int) int {
	return util.Mod(dir, 4)
}

type border struct {
	id      uint
	flipped bool
}

type tile struct {
	id      int
	rawData [][]rune
	borders [4]border
}

func copy2d(dst [][]rune, src [][]rune) {
	for y := 0; y < util.Min(len(dst), len(src)); y++ {
		copy(dst[y], src[y])
	}
}

func clone2d(d [][]rune) [][]rune {
	res := make([][]rune, len(d))
	for y, row := range d {
		res[y] = make([]rune, len(row))
	}

	copy2d(res, d)

	return res
}

func rotate(d [][]rune, a int) [][]rune {
	a = normalizeDir(a)
	if a > 0 {
		ly := len(d) - 1
		for y := 0; y < len(d)/2; y++ {
			lx := len(d[y]) - 1
			for x := 0; x < (len(d[y])+1)/2; x++ {
				switch a {
				case 1:
					d[y][x], d[x][lx-y], d[ly-y][lx-x], d[ly-x][y] = d[x][lx-y], d[ly-y][lx-x], d[ly-x][y], d[y][x]
				case 2:
					d[y][x], d[x][lx-y], d[ly-y][lx-x], d[ly-x][y] = d[ly-y][lx-x], d[ly-x][y], d[y][x], d[x][lx-y]
				case 3:
					d[y][x], d[x][lx-y], d[ly-y][lx-x], d[ly-x][y] = d[ly-x][y], d[y][x], d[x][lx-y], d[ly-y][lx-x]
				}
			}
		}
	}
	return d
}

func flip(d [][]rune, fd int) [][]rune {
	if fd%2 == 0 {
		for y := 0; y < len(d); y++ {
			lx := len(d[y]) - 1
			for x := 0; x < len(d[y])/2; x++ {
				d[y][x], d[y][lx-x] = d[y][lx-x], d[y][x]
			}
		}
	} else {
		ly := len(d) - 1
		for y := 0; y < len(d)/2; y++ {
			for x := 0; x < len(d[y]); x++ {
				d[y][x], d[ly-y][x] = d[ly-y][x], d[y][x]
			}
		}
	}
	return d
}

func (t tile) rotate(a int) tile {
	res := tile{
		id:      t.id,
		rawData: rotate(clone2d(t.rawData), a),
		borders: t.borders,
	}

	switch normalizeDir(a) {
	case 1:
		res.borders[top], res.borders[right], res.borders[bottom], res.borders[left] = res.borders[right], res.borders[bottom], res.borders[left], res.borders[top]
	case 2:
		res.borders[top], res.borders[right], res.borders[bottom], res.borders[left] = res.borders[bottom], res.borders[left], res.borders[top], res.borders[right]
	case 3:
		res.borders[top], res.borders[right], res.borders[bottom], res.borders[left] = res.borders[left], res.borders[top], res.borders[right], res.borders[bottom]
	}

	return res
}

func (t tile) flip(fd int) tile {
	res := tile{
		id:      t.id,
		rawData: flip(clone2d(t.rawData), fd),
		borders: t.borders,
	}

	if fd%2 == 0 {
		res.borders[left], res.borders[right] = res.borders[right], res.borders[left]
	} else {
		res.borders[top], res.borders[bottom] = res.borders[bottom], res.borders[top]
	}

	for d := range res.borders {
		res.borders[d].flipped = !res.borders[d].flipped
	}

	return res
}

func (t tile) borderDir(bid uint) (int, bool) {
	for d, b := range t.borders {
		if b.id == bid {
			return d, b.flipped
		}
	}
	util.Panic("Border %d not found in tile %d", bid, t.id)
	return 0, false
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
	var corner tile
	cm := 0
	for _, t := range tiles {
		e := 0
		bm := 0
		for d, b := range t.borders {
			if len(borders[b.id]) == 1 {
				e++
				bm |= 1 << d
			}
		}
		if e == 2 {
			prod *= t.id
			c++
			corner = t
			cm = bm
		} else if e > 2 {
			util.Panic("WTF! Tile %d has %d edges", t.id, e)
		}
	}
	if c != 4 {
		util.Panic("WTF! Image has %d corners", c)
	}
	log.Part1(prod)

	// Part 2
	tileImg := [][]tile{}
	switch cm {
	case 0b0011:
		tileImg = append(tileImg, []tile{corner.rotate(1)})
	case 0b0110:
		tileImg = append(tileImg, []tile{corner.rotate(2)})
	case 0b1100:
		tileImg = append(tileImg, []tile{corner.rotate(3)})
	case 0b1001:
		tileImg = append(tileImg, []tile{corner})
	default:
		util.Panic("Tile is not a corner")
	}
	for y := 0; len(borders[tileImg[y][0].borders[bottom].id]) == 2; y++ {
		t := tileImg[y][0]
		b := t.borders[bottom]
		for _, tid := range borders[b.id] {
			if tid != t.id {
				at := tiles[tid]
				d, f := at.borderDir(b.id)
				at = at.rotate(d)
				if f == b.flipped {
					at = at.flip(top)
				}
				tileImg = append(tileImg, []tile{at})
				break
			}
		}
		if y == len(tileImg)-1 {
			util.Panic("Didn't find bottom adjacent tile for tile %d, border %d", t.id, b.id)
		}
	}
	for y := range tileImg {
		for x := 0; len(borders[tileImg[y][x].borders[right].id]) == 2; x++ {
			t := tileImg[y][x]
			b := t.borders[right]
			for _, tid := range borders[b.id] {
				if tid != t.id {
					at := tiles[tid]
					d, f := at.borderDir(b.id)
					at = at.rotate(d + 1)
					if f == b.flipped {
						at = at.flip(left)
					}
					tileImg[y] = append(tileImg[y], at)
					break
				}
			}
			if x == len(tileImg[y])-1 {
				util.Panic("Didn't find right adjacent tile for tile %d, border %d", t.id, b.id)
			}
		}
	}

	for y := range tileImg {
		if len(tileImg[y]) != len(tileImg[0]) {
			util.Panic("Invalid image width")
		}
	}
	if len(tileImg) != len(tileImg[0]) {
		util.Panic("Invalid image height")
	}

	img := [][]rune{}
	y := -1
	for _, r := range tileImg {
		for ty := 1; ty < len(r[0].rawData)-1; ty++ {
			img = append(img, []rune{})
			y++
			for _, t := range r {
				img[y] = append(img[y], t.rawData[ty][1:len(t.rawData[ty])-1]...)
			}
		}
	}

	monster := [][]rune{
		[]rune("                  # "),
		[]rune("#    ##    ##    ###"),
		[]rune(" #  #  #  #  #  #   "),
	}

	type point struct {
		x, y int
	}

	allXY := map[point]bool{}
	for f := 0; f < 2 && len(allXY) == 0; f++ {
		for r := 0; r < 4 && len(allXY) == 0; r++ {
			for y := 0; y < len(img)-len(monster); y++ {
				for x := 0; x < len(img[y])-len(monster[0]); x++ {
					found := true
					foundXY := map[point]bool{}
					for my := range monster {
						for mx, mr := range monster[my] {
							if mr != ' ' && mr != img[y+my][x+mx] {
								found = false
								break
							} else if mr == '#' {
								foundXY[point{y + my, x + mx}] = true
							}
						}
						if !found {
							break
						}
					}
					if found {
						for p := range foundXY {
							allXY[p] = true
						}
					}
				}
			}
			rotate(img, 1)
		}
		flip(img, top)
	}

	count := 0
	for y := range img {
		for _, r := range img[y] {
			if r == '#' {
				count++
			}
		}
	}
	log.Part2(count - len(allXY))
}
