package piscine

import (
	"github.com/01-edu/z01"
)

func TabInit(n int) []int {
	var tab []int
	for n > 0 {
		tab = append(tab, n%10)
		n /= 10
	}
	return tab
}

func PrintN(n int) {
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

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	i := 0
	tab := TabInit(n)
	for i < len(tab)-1 {
		if tab[i] > tab[i+1] {
			swap := tab[i]
			tab[i] = tab[i+1]
			tab[i+1] = swap
			i = 0
		} else {
			i++
		}
	}
	i = 0
	for i <= len(tab)-1 {
		PrintN(tab[i])
		i++
	}
}
