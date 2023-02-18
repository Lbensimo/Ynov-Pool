package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	if !isEven(len(os.Args)) {
		PrintStr(EvenMsg)
	} else {
		PrintStr(OddMsg)
	}
}

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)
