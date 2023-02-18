package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	a := '0'
	b := '1'
	c := '2'
	for a <= '7' {
		for b <= '8' {
			for c <= '9' {
				if a != '7' {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('7')
					z01.PrintRune('8')
					z01.PrintRune('9')
					z01.PrintRune('\n')
				}
				c++
			}
			b++
			c = b + 1
		}
		a++
		b = a + 1
		c = b + 1
	}
}
