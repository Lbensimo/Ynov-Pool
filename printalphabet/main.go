package main

import (
	"github.com/01-edu/z01"
)

func main() {
	lt := 'a'
	for lt <= 'z' {
		z01.PrintRune(lt)
		lt++
	}
	z01.PrintRune('\n')
}
