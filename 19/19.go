package main

import (
	"regexp"
	"strings"

	"github.com/liennie/aoc2020/common/load"
	"github.com/liennie/aoc2020/common/log"
	"github.com/liennie/aoc2020/common/util"
)

type rule interface {
	validate(msg string, rules map[int]rule) []int
}

type constRule struct {
	rule string
}

func (r *constRule) validate(msg string, rules map[int]rule) []int {
	if strings.HasPrefix(msg, string(r.rule)) {
		return []int{len(string(r.rule))}
	}
	return []int{}
}

type recRule struct {
	rules [][]int
}

func (r *recRule) validate(msg string, rules map[int]rule) []int {
	res := []int{}

	for _, rs := range r.rules {
		os := []int{0}

		for _, ri := range rs {
			nos := []int{}
			for _, o := range os {
				if rr := rules[ri]; rr != nil {
					for _, ros := range rr.validate(msg[o:], rules) {
						nos = append(nos, ros+o)
					}
				} else {
					util.Panic("Rule %d doesn't exist", ri)
				}
			}
			os = nos
		}

		res = append(res, os...)
	}

	return util.Uniq(res)
}

var (
	ruleRe = regexp.MustCompile(`^(\d+): (?:(".*?")|(\d+(?: \d+)*(?: \| \d+(?: \d+)*)*))$`)
)

func parseRule(l string) (int, rule) {
	match := ruleRe.FindStringSubmatch(l)
	if len(match) != 4 {
		util.Panic("No match for %q", l)
	}

	i := util.Atoi(match[1])

	if match[2] != "" {
		return i, &constRule{strings.Trim(match[2], "\"")}
	} else {
		r := &recRule{}
		for _, rss := range strings.Split(match[3], "|") {
			rs := []int{}
			for _, rr := range strings.Split(strings.TrimSpace(rss), " ") {
				rs = append(rs, util.Atoi(rr))
			}
			r.rules = append(r.rules, rs)
		}
		return i, r
	}
}

func parse(filename string) (map[int]rule, []string) {
	ch := load.File(filename)

	rules := map[int]rule{}
	for l := range ch {
		if l == "" {
			break
		}

		i, r := parseRule(l)
		if rules[i] != nil {
			util.Panic("Rule %d already exists", i)
		}
		rules[i] = r
	}

	messages := []string{}
	for l := range ch {
		messages = append(messages, l)
	}

	return rules, messages
}

func main() {
	defer util.Recover(log.Err)

	rules, messages := parse("input.txt")
	rule0 := rules[0]
	if rule0 == nil {
		util.Panic("Rule 0 doesn't exist")
	}

	// Part 1
	count := 0
	for _, msg := range messages {
		if util.Contains(rule0.validate(msg, rules), len(msg)) {
			count++
		}
	}
	log.Part1(count)

	// Part 2
	_, rules[8] = parseRule("8: 42 | 42 8")
	_, rules[11] = parseRule("11: 42 31 | 42 11 31")
	count = 0
	for _, msg := range messages {
		if util.Contains(rule0.validate(msg, rules), len(msg)) {
			count++
		}
	}
	log.Part2(count)
}
