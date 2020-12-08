package asm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/liennie/aoc2020/common/asm/op"
)

func decode(line int, l string) (i *Instruction, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Line %d: %w", line, err)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	f := strings.Fields(l)
	if len(f) == 0 {
		return nil, nil
	}

	args := make([]int, len(f)-1)
	for i, arg := range f[1:] {
		var err error
		args[i], err = strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("Arg %d: %w", i, err)
		}
	}

	return newInstruction(line, op.FromString(f[0]), args...)
}

func Load(r io.Reader) (*Program, error) {
	p := &Program{}

	line := 1
	br := bufio.NewReader(r)
	for {
		l, err := br.ReadString('\n')
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			i, err := decode(line, l)
			if err != nil {
				return nil, err
			}
			if i != nil {
				p.instructions = append(p.instructions, *i)
			}
		}
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("ReadString: %w", err)
			}
			break
		}

		line++
	}

	return p, nil
}

func LoadFile(filename string) (*Program, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Load(file)
}
