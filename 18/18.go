package main

import (
	"fmt"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type expr interface {
	eval() int
}

type constExpr int

func (e constExpr) eval() int {
	return int(e)
}

type opExpr struct {
	ex []expr
	op []rune
}

func (e opExpr) eval() int {
	if len(e.ex) != len(e.op)+1 {
		util.Panic("Lengths don't match: %d, %d", len(e.ex), len(e.op))
	}

	res := e.ex[0].eval()
	for i, op := range e.op {
		switch op {
		case '+':
			res += e.ex[i+1].eval()
		case '*':
			res *= e.ex[i+1].eval()
		}
	}

	return res
}

type tokenType int

const (
	ttNil tokenType = iota
	ttNum
	ttPlus
	ttMul
	ttLParen
	ttRParen
)

type token struct {
	tt  tokenType
	pos int
	val int
}

func (t token) String() string {
	switch t.tt {
	case ttNil:
		return "nil"
	case ttNum:
		return fmt.Sprintf("int(%d)", t.val)
	case ttPlus, ttMul, ttLParen, ttRParen:
		return string(t.val)
	}
	return "invalid"
}

func tokenizeExpr(e string) []token {
	if len(e) == 0 {
		return nil
	}

	tokens := []token{}
	n := 0
	ni := 0
	nn := false
	for i, c := range e {
		if c >= '0' && c <= '9' {
			if !nn {
				nn = true
				ni = i
				n = 0
			}
			n = n*10 + int(c-'0')
		} else {
			if nn {
				tokens = append(tokens, token{ttNum, ni + 1, n})
				nn = false
			}

			switch c {
			case '+':
				tokens = append(tokens, token{ttPlus, i + 1, '+'})
			case '*':
				tokens = append(tokens, token{ttMul, i + 1, '*'})
			case '(':
				tokens = append(tokens, token{ttLParen, i + 1, '('})
			case ')':
				tokens = append(tokens, token{ttRParen, i + 1, ')'})
			case ' ':
				// skip whitespace
			default:
				util.Panic("Invalid character %q, pos %d", string(c), i+1)
			}
		}
	}
	if nn {
		tokens = append(tokens, token{ttNum, ni + 1, n})
	}
	return tokens
}

func parseExpr(tokens []token) (expr, int) {
	ex := opExpr{}

	ee := true
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]

		if ee {
			if t.tt == ttNum {
				ex.ex = append(ex.ex, constExpr(t.val))
			} else if t.tt == ttLParen {
				e, ni := parseExpr(tokens[i+1:])
				if ni == -1 {
					util.Panic("Unexpected end of expression")
				}
				i += ni + 1
				ex.ex = append(ex.ex, e)
			} else {
				util.Panic("Invalid token %q, pos %d, expected expression", t, t.pos)
			}
			ee = false
		} else {
			if t.tt == ttRParen {
				return ex, i
			} else if t.tt == ttPlus || t.tt == ttMul {
				ex.op = append(ex.op, rune(t.val))
			} else {
				util.Panic("Invalid token %q, pos %d, expected operator", t, t.pos)
			}
			ee = true
		}
	}

	return ex, -1
}

func parse(filename string) []expr {
	res := []expr{}
	for l := range load.File(filename) {
		ex, i := parseExpr(tokenizeExpr(l))
		if i != -1 {
			util.Panic("Unexpected end of expression, token %d", i)
		}
		res = append(res, ex)
	}
	return res
}

func main() {
	defer util.Recover(log.Err)

	r := parse("input.txt")

	// Part 1
	sum := 0
	for _, e := range r {
		sum += e.eval()
	}
	log.Part1(sum)
}
