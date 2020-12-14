package util

import (
	"fmt"
	"strconv"
)

func Mod(a, b int) int {
	m := a % b
	if m < 0 {
		return m + b
	}
	return m
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return Abs(a*b) / GCD(a, b)
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		Panic("Atoi(%s): %w", s, err)
	}
	return i
}

func Panic(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}

func Recover(f func(err error)) {
	if e := recover(); e != nil {
		switch e.(type) {
		case error:
			f(e.(error))
		default:
			f(fmt.Errorf("%v", e))
		}
	}
}

func Perm(n int) [][]int {
	if n < 0 {
		Panic("Perm(%d)", n)
	}

	if n == 0 {
		return [][]int{{}}
	}

	res := Perm(n - 1)
	l := len(res)
	for i := 0; i < l; i++ {
		p := make([]int, len(res[i]))
		copy(p, res[i])
		p = append(p, n-1)
		res = append(res, p)
	}
	return res
}
