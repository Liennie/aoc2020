package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	input = "input.txt"
)

func load(filename string) ([][]byte, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var res [][]byte
	for {
		b, err := r.ReadBytes('\n')
		b = bytes.TrimSpace(b)
		if len(b) > 0 {
			res = append(res, b)
		}
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("ReadBytes: %w", err)
			}
			break
		}
	}

	return res, nil
}

func trees(b [][]byte, sx, sy int) int {
	x, y := 0, 0
	trees := 0

	for y < len(b) {
		if b[y][x] == '#' {
			trees++
		}

		x = (x + sx) % len(b[y]) // the width is the same for all lines, this is fine
		y += sy
	}

	return trees
}

func main() {
	b, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	log.Printf("Part 1: %d", trees(b, 3, 1))

	// Part 2
	slopes := []struct{ x, y int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	p := 1
	for _, slope := range slopes {
		p *= trees(b, slope.x, slope.y)
	}

	log.Printf("Part 2: %d", p)
}
