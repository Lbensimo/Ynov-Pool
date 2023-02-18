package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	i := n - 1
	if n == 1 {
		j := 0
		for j < 9 {
			z01.PrintRune(rune(j + 48))
			z01.PrintRune(',')
			z01.PrintRune(' ')
			j++
		}
		z01.PrintRune('9')
		z01.PrintRune('\n')
		return
	}
	var tab []int = TabInitN(n)
	for tab[0] < 10-n {
		for tab[n-1] <= 9 {
			PrintTab(n, tab)
			tab[n-1]++
		}
		if tab[i-1] == tab[i]-1 {
			i--
		} else {
			tab[i-1]++
			tab[i] = tab[i-1] + 1
			for i2 := 0; i2 < n-1; i2++ {
				tab[i] = tab[i-1] + 1
				if i != n-1 {
					i++
				}
			}
		}
	}
	PrintTab(n, tab)
}

func PrintTab(n int, tab []int) {
	i := 0
	if tab[0] == 10-n {
		for i < n {
			z01.PrintRune(rune(tab[i] + 48))
			i++
		}
		z01.PrintRune('\n')
		return
	} else {
		for i < n {
			z01.PrintRune(rune(tab[i] + 48))
			i++
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func TabInitN(n int) []int {
	i := 0
	var tab []int
	for i < n {
		tab = append(tab, i)
		i++
	}
	return tab
}
