package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func data(s string) [][]rune {
	res := [][]rune{}
	for _, l := range strings.Split(s, ";") {
		res = append(res, []rune(l))
	}
	return res
}

func TestRotate(t *testing.T) {
	tests := []struct {
		data string
		a    int
		want string
	}{
		{"ab;cd", 0, "ab;cd"},
		{"ab;cd", 1, "bd;ac"},
		{"ab;cd", 2, "dc;ba"},
		{"ab;cd", 3, "ca;db"},
		{`abc;def;ghi`, 0, `abc;def;ghi`},
		{`abc;def;ghi`, 1, `cfi;beh;adg`},
		{`abc;def;ghi`, 2, `ihg;fed;cba`},
		{`abc;def;ghi`, 3, `gda;heb;ifc`},
		{`abcd;efgh;ijkl;mnop`, 0, `abcd;efgh;ijkl;mnop`},
		{`abcd;efgh;ijkl;mnop`, 1, `dhlp;cgko;bfjn;aeim`},
		{`abcd;efgh;ijkl;mnop`, 2, `ponm;lkji;hgfe;dcba`},
		{`abcd;efgh;ijkl;mnop`, 3, `miea;njfb;okgc;plhd`},
		{`abcde;fghij;klmno;pqrst;uvwxy`, 0, `abcde;fghij;klmno;pqrst;uvwxy`},
		{`abcde;fghij;klmno;pqrst;uvwxy`, 1, `ejoty;dinsx;chmrw;bglqv;afkpu`},
		{`abcde;fghij;klmno;pqrst;uvwxy`, 2, `yxwvu;tsrqp;onmlk;jihgf;edcba`},
		{`abcde;fghij;klmno;pqrst;uvwxy`, 3, `upkfa;vqlgb;wrmhc;xsnid;ytoje`},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/%d", test.data, test.a), func(t *testing.T) {
			a := assertions.New(t)
			a.So(rotate(data(test.data), test.a), should.Resemble, data(test.want))
		})
	}
}

func TestFlip(t *testing.T) {
	tests := []struct {
		data string
		fd   int
		want string
	}{
		{"ab;cd", top, "ba;dc"},
		{"ab;cd", left, "cd;ab"},
		{`abc;def;ghi`, top, `cba;fed;ihg`},
		{`abc;def;ghi`, left, `ghi;def;abc`},
		{`abcd;efgh;ijkl;mnop`, top, `dcba;hgfe;lkji;ponm`},
		{`abcd;efgh;ijkl;mnop`, left, `mnop;ijkl;efgh;abcd`},
		{`abcde;fghij;klmno;pqrst;uvwxy`, top, `edcba;jihgf;onmlk;tsrqp;yxwvu`},
		{`abcde;fghij;klmno;pqrst;uvwxy`, left, `uvwxy;pqrst;klmno;fghij;abcde`},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/%d", test.data, test.fd), func(t *testing.T) {
			a := assertions.New(t)
			a.So(flip(data(test.data), test.fd), should.Resemble, data(test.want))
		})
	}
}
