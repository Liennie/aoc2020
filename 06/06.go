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

func load(filename string) ([][]map[rune]struct{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var res [][]map[rune]struct{}
	var group []map[rune]struct{}
	for {
		l, err := r.ReadString('\n')
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			m := map[rune]struct{}{}
			for _, r := range l {
				m[r] = struct{}{}
			}
			group = append(group, m)
		} else if len(group) > 0 {
			res = append(res, group)
			group = nil
		}
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("ReadString: %w", err)
			}
			break
		}
	}
	if len(group) > 0 {
		res = append(res, group)
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
	answers, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	sum := 0
	for _, group := range answers {
		sum += len(any(group))
	}
	log.Printf("Part 1: %d", sum)

	// Part 2
	sum = 0
	for _, group := range answers {
		sum += len(all(group))
	}
	log.Printf("Part 2: %d", sum)
}
