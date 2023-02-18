package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(", y = ")
	PrintNbr(points.y)
	z01.PrintRune('\n')
}

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for range s {
		c++
		z01.PrintRune(a[c-1])
	}
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		PrintNbr(n / 10)
		z01.PrintRune('8')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	if n >= 10 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	} else {
		z01.PrintRune(rune(n) + 48)
	}
}
