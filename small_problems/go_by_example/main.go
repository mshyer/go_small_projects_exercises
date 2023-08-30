package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Replace:   ", s.Replace("fooo00oooo aa ooo", "o", "0", -1))
	p("Replace:   ", s.Replace("foossooo", "o", "0", 2))
	p("Split:     ", s.Split("a-b-c-d-e", "-")[1])
}