package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	input = "input.txt"
)

type passport struct {
	birthYear      int    `field:"byr"`
	issueYear      int    `field:"iyr"`
	expirationYear int    `field:"eyr"`
	height         string `field:"hgt"`
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
	}

	switch key {
	case "byr":
		p.birthYear = ivalue
	case "iyr":
		p.issueYear = ivalue
	case "eyr":
		p.expirationYear = ivalue
	case "hgt":
		p.height = value
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
		p.height != "" &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportID != ""
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
}
