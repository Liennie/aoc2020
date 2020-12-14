package main

import (
	"math"
	"regexp"

	"github.com/liennie/aoc2020/common/util"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/recover"
)

const (
	input  = "input.txt"
	mask36 = 0xfffffffff
)

func parseMask(mask string) (uint64, uint64) {
	maskAnd := uint64(math.MaxUint64)
	maskOr := uint64(0)

	for i, c := range mask {
		switch c {
		case '1':
			maskOr |= 1 << (len(mask) - i - 1)
		case '0':
			maskAnd &= ^(1 << (len(mask) - i - 1))
		case 'X':
			//
		default:
			util.Panic("Invalid mask %q", mask)
		}
	}

	return maskAnd & mask36, maskOr & mask36
}

func parse(filename string) map[int]uint64 {
	maskRe := regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

	maskAnd, maskOr := parseMask("")
	mem := map[int]uint64{}
	for l := range load.File(filename) {
		if match := maskRe.FindStringSubmatch(l); len(match) == 2 {
			maskAnd, maskOr = parseMask(match[1])
		} else if match = memRe.FindStringSubmatch(l); len(match) == 3 {
			mem[util.Atoi(match[1])] = (uint64(util.Atoi(match[2])) & maskAnd) | maskOr
		} else {
			log.Println(len(match))
			util.Panic("No match for %q", l)
		}
	}
	return mem
}

func main() {
	defer recover.Err(log.Err)

	mem := parse(input)

	// Part 1
	sum := uint64(0)
	for _, n := range mem {
		sum += n
	}
	log.Part1(int(sum))
}
