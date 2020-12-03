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

func main() {
	b, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1

	x, y := 0, 0
	trees := 0

	for y < len(b) {
		if b[y][x] == '#' {
			trees++
		}

		x = (x + 3) % len(b[y]) // the width is the same for all lines, this is fine
		y++
	}

	log.Printf("Part 1: %d", trees)
}
