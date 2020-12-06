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

func load(filename string) ([]map[rune]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var res []map[rune]bool
	m := map[rune]bool{}
	for {
		l, err := r.ReadString('\n')
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			for _, c := range l {
				m[c] = true
			}
		} else if len(m) > 0 {
			res = append(res, m)
			m = map[rune]bool{}
		}
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("ReadString: %w", err)
			}
			break
		}
	}
	if len(m) > 0 {
		res = append(res, m)
	}

	return res, nil
}

func main() {
	answers, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	sum := 0
	for _, a := range answers {
		sum += len(a)
	}
	log.Printf("Part 1: %d", sum)
}
