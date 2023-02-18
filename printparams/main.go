package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for c < len(a) {
		c++
		z01.PrintRune(a[c-1])
	}
	z01.PrintRune('\n')
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		PrintStr(os.Args[i])
	}
}
