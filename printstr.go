package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for range s {
		c++
		z01.PrintRune(a[c-1])
	}
}
