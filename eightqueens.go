package piscine

import (
	"github.com/01-edu/z01"
)

var tab [8][8]int

func EightQueens() {
	Recur(0)
}

func Recur(x int) {
	for y := 0; y < 8; y++ {
		if Verif(tab, x, y) {
			tab[y][x] = 1
			if x >= 7 {
				Aff(tab)
			}
			Recur(x + 1)
			tab[y][x] = 0
		}
	}
}

func Aff(tab [8][8]int) {
	i := 0
	i2 := 0
	for i2 < 8 {
		for tab[i][i2] != 1 {
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

func Fou(tab [8][8]int, x int, y int) bool {
	a := x - 1
	o := y - 1
	for a >= 0 && o >= 0 {
		if tab[o][a] == 1 {
			return false
		}
		a--
		o--
	}
	a = x + 1
	o = y + 1
	for a < 8 && o < 8 {
		if tab[o][a] == 1 {
			return false
		}
		a++
		o++
	}
	a = x + 1
	o = y - 1
	for a < 8 && o >= 0 {
		if tab[o][a] == 1 {
			return false
		}
		a++
		o--
	}
	a = x - 1
	o = y + 1
	for a >= 0 && o < 8 {
		if tab[o][a] == 1 {
			return false
		}
		a--
		o++
	}
	return true
}

func Verif(tab [8][8]int, x int, y int) bool {
	i := 0
	for i < 8 {
		if tab[y][i] == 1 {
			return false
		}
		i++
	}
	return Fou(tab, x, y)
}
