package main

import (
	"github.com/01-edu/z01"
)

func main() {
	lt := 'z'
	for lt >= 'a' {
		z01.PrintRune(lt)
		lt--
	}
	z01.PrintRune('\n')
}
