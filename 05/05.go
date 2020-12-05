package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	input = "input.txt"
)

type seat struct {
	row, col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func decode(s string) seat {
	row := 0
	for i, c := range s[:7] {
		if c == 'B' {
			row += 1 << (6 - i)
		} else if c != 'F' {
			panic(fmt.Errorf("Invalid character: %q", c))
		}
	}

	col := 0
	for i, c := range s[7:10] {
		if c == 'R' {
			col += 1 << (2 - i)
		} else if c != 'L' {
			panic(fmt.Errorf("Invalid character: %q", c))
		}
	}

	return seat{
		row: row,
		col: col,
	}
}

func load(filename string) ([]seat, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var res []seat
	for {
		l, err := r.ReadString('\n')
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			res = append(res, decode(l))
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

func main() {
	seats, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	max := 0
	for i, s := range seats {
		if s.id() > max || i == 0 {
			max = s.id()
		}
	}
	log.Printf("Part 1: %d", max)
}
