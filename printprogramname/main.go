package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	a := []rune(s)
	c := 2
	for c < len(a) {
		c++
		z01.PrintRune(a[c-1])
	}
	z01.PrintRune('\n')
}

func main() {
	for i := 0; i < len(os.Args[:1]); i++ {
		PrintStr(os.Args[:1][i])
	}
}
