package piscine

import (
	"github.com/01-edu/z01"
)

func Print(a int, b int) {
	if a == 98 {
		z01.PrintRune('9')
		z01.PrintRune('8')
		z01.PrintRune(' ')
		z01.PrintRune('9')
		z01.PrintRune('9')
		return
	}
	if a < 10 {
		z01.PrintRune('0')
		z01.PrintRune(rune(a) + 48)
	} else {
		z01.PrintRune(rune(a/10) + 48)
		z01.PrintRune(rune(a%10) + 48)
	}
	z01.PrintRune(' ')
	if b < 10 {
		z01.PrintRune('0')
		z01.PrintRune(rune(b) + 48)
		z01.PrintRune(',')
		z01.PrintRune(' ')
	} else {
		z01.PrintRune(rune(b/10) + 48)
		z01.PrintRune(rune(b%10) + 48)
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func PrintComb2() {
	a := 0
	b := 1
	for a <= 98 {
		for b <= 99 {
			Print(a, b)
			b++
		}
		a++
		b = a + 1
	}
	z01.PrintRune('\n')
}
