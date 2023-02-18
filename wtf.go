package piscine

import "github.com/01-edu/z01"

func Foutredieu(tab [8][8]int) {
	i := 0
	i2 := 0
	for i2 < 8 {
		for tab[i2][i] != 1 {
			if i > 7 {
				i = 0
				i2++
			}
			i++
		}
		PrintNbr(i + 1)
		i2++
		i = 0
	}
	z01.PrintRune('\n')
}
