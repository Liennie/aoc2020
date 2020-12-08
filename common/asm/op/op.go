package op

import (
	"fmt"
	"strings"
)

type Code int

const (
	Nop Code = iota
	Acc
	Jmp
)

var codeStr = map[Code]string{
	Nop: "nop",
	Acc: "acc",
	Jmp: "jmp",
}

var strCode map[string]Code

func init() {
	strCode = map[string]Code{}
	for k, v := range codeStr {
		strCode[v] = k
	}
}

func (c Code) String() string {
	return codeStr[c]
}

func FromString(s string) Code {
	if c, ok := strCode[strings.ToLower(s)]; ok {
		return c
	}
	panic(fmt.Errorf("Invalid instruction %q", strings.ToLower(s)))
}
