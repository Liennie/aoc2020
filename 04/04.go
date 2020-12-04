package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	input = "input.txt"
)

type passport struct {
	birthYear      int `field:"byr"`
	issueYear      int `field:"iyr"`
	expirationYear int `field:"eyr"`
	height         int `field:"hgt"`
	heightUnit     string
	hairColor      string `field:"hcl"`
	eyeColor       string `field:"ecl"`
	passportID     string `field:"pid"`
	countryID      string `field:"cid"`
}

func (p *passport) set(key, value string) {
	ivalue := 0
	switch key {
	case "byr", "iyr", "eyr":
		var err error
		ivalue, err = strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
	case "hgt":
		var err error
		i := strings.IndexFunc(value, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		if i == -1 {
			ivalue, err = strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			value = ""
		} else {
			ivalue, err = strconv.Atoi(value[:i])
			if err != nil {
				panic(err)
			}
			value = value[i:]
		}
	}

	switch key {
	case "byr":
		p.birthYear = ivalue
	case "iyr":
		p.issueYear = ivalue
	case "eyr":
		p.expirationYear = ivalue
	case "hgt":
		p.height = ivalue
		p.heightUnit = value
	case "hcl":
		p.hairColor = value
	case "ecl":
		p.eyeColor = value
	case "pid":
		p.passportID = value
	case "cid":
		p.countryID = value
	}
}

func (p *passport) valid() bool {
	return p.birthYear != 0 &&
		p.issueYear != 0 &&
		p.expirationYear != 0 &&
		(p.height != 0 || p.heightUnit != "") &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportID != ""
}

func isHex(s string) bool {
	_, err := strconv.ParseUint(s, 16, 64)
	return err == nil
}

func isColor(s string) bool {
	switch s {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func isDigit(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

func (p *passport) validStrict() bool {
	return p.birthYear >= 1920 && p.birthYear <= 2002 &&
		p.issueYear >= 2010 && p.issueYear <= 2020 &&
		p.expirationYear >= 2020 && p.expirationYear <= 2030 &&
		((p.heightUnit == "cm" && p.height >= 150 && p.height <= 193) ||
			(p.heightUnit == "in" && p.height >= 59 && p.height <= 76)) &&
		(len(p.hairColor) == 7 && p.hairColor[0] == '#' && isHex(p.hairColor[1:])) &&
		isColor(p.eyeColor) &&
		len(p.passportID) == 9 && isDigit(p.passportID)
}

func load(filename string) ([]passport, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var res []passport
	var p passport
	var emptyP passport
	for {
		l, err := r.ReadString('\n')
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			fields := strings.Fields(l)
			for _, field := range fields {
				s := strings.SplitN(field, ":", 2)
				if len(s) != 2 {
					continue // ignore
				}

				p.set(s[0], s[1])
			}
		} else if p != emptyP {
			res = append(res, p)
			p = emptyP
		}
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("ReadString: %w", err)
			}
			break
		}
	}
	if p != emptyP {
		res = append(res, p)
		p = emptyP
	}

	return res, nil
}

func main() {
	passports, err := load(input)
	if err != nil {
		log.Printf("Load: %s", err)
	}

	// Part 1
	valid := 0
	for _, p := range passports {
		if p.valid() {
			valid++
		}
	}
	log.Printf("Part 1: %d", valid)

	// Part 2
	valid = 0
	for _, p := range passports {
		if p.validStrict() {
			valid++
		}
	}
	log.Printf("Part 2: %d", valid)
}
