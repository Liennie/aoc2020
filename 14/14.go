package main

import (
	"math"
	"regexp"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

const (
	input  = "input.txt"
	mask36 = 0xfffffffff
)

func parseMask(mask string) (uint64, uint64, uint64) {
	maskAnd := uint64(math.MaxUint64)
	maskOr := uint64(0)
	maskFloat := uint64(0)

	for i, c := range mask {
		switch c {
		case '1':
			maskOr |= 1 << (len(mask) - i - 1)
		case '0':
			maskAnd &= ^(1 << (len(mask) - i - 1))
		case 'X':
			maskFloat |= 1 << (len(mask) - i - 1)
		default:
			util.Panic("Invalid mask %q", mask)
		}
	}

	return maskAnd & mask36, maskOr & mask36, maskFloat & mask36
}

var (
	maskRe = regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRe  = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

func parse1(filename string) map[int]uint64 {
	maskAnd, maskOr, _ := parseMask("")
	mem := map[int]uint64{}
	for l := range load.File(filename) {
		if match := maskRe.FindStringSubmatch(l); len(match) == 2 {
			maskAnd, maskOr, _ = parseMask(match[1])
		} else if match = memRe.FindStringSubmatch(l); len(match) == 3 {
			mem[util.Atoi(match[1])] = (uint64(util.Atoi(match[2])) & maskAnd) | maskOr
		} else {
			log.Println(len(match))
			util.Panic("No match for %q", l)
		}
	}
	return mem
}

func floatMasks(float uint64) []uint64 {
	bits := []int{}
	for i := 0; i < 36; i++ {
		if (float>>i)&1 == 1 {
			bits = append(bits, i)
		}
	}

	res := []uint64{}
	for _, p := range util.Perm(len(bits)) {
		mask := uint64(0)
		for _, i := range p {
			mask |= 1 << bits[i]
		}
		res = append(res, mask)
	}

	return res
}

func parse2(filename string) map[int]uint64 {
	_, maskOr, maskFloat := parseMask("")
	mem := map[int]uint64{}
	for l := range load.File(filename) {
		if match := maskRe.FindStringSubmatch(l); len(match) == 2 {
			_, maskOr, maskFloat = parseMask(match[1])
		} else if match = memRe.FindStringSubmatch(l); len(match) == 3 {
			for _, mask := range floatMasks(maskFloat) {
				mem[int((uint64(util.Atoi(match[1]))&(^maskFloat))|mask|maskOr)] = uint64(util.Atoi(match[2])) & mask36
			}
		} else {
			log.Println(len(match))
			util.Panic("No match for %q", l)
		}
	}
	return mem
}

func main() {
	defer util.Recover(log.Err)

	// Part 1
	sum := uint64(0)
	for _, n := range parse1(input) {
		sum += n
	}
	log.Part1(int(sum))

	// Part 2
	sum = 0
	for _, n := range parse2(input) {
		sum += n
	}
	log.Part2(int(sum))
}
