package main

import (
	"github.com/01-edu/z01"
)

func main() {
	lt := '0'
	for lt <= '9' {
		z01.PrintRune(lt)
		lt++
	}
	z01.PrintRune('\n')
}
