package piscine

import "github.com/01-edu/z01"

func CB(base string) bool {
	i := 0
	if len(base) < 2 {
		return false
	}
	for i < len(base) {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		i2 := i + 1
		for i2 < len(base) {
			if base[i] == base[i2] {
				return false
			}
			i2++
		}
		i++
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if !CB(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == -9223372036854775808 {
		PrintNbrBase(-922337203, base)
		PrintNbrBase(6854775808, base)
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}
	if nbr >= len(base) {
		PrintNbrBase(nbr/len(base), base)
		PrintNbrBase(nbr%len(base), base)
	} else {
		z01.PrintRune(rune(base[nbr]))
	}
}
