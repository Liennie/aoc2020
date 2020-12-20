package log

import (
	"log"
)

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
}

func Err(err error) {
	log.Printf("Error: %s", err)
}

func Part1(n int) {
	log.Printf("Part 1: %d", n)
}

func Part2(n int) {
	log.Printf("Part 2: %d", n)
}

func Print(v ...interface{}) {
	log.Print(v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}
