package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	a := Atoi(os.Args[1])
	b := os.Args[2]
	c := Atoi(os.Args[3])
	if CheckBA(os.Args[1]) && CheckBA(os.Args[3]) && (b == "+" || b == "-" ||
		b == "/" || b == "*" || b == "%") && len(os.Args[1]) < 19 {
		if b == "+" {
			PrintNbr(a + c)
			z01.PrintRune('\n')
			return
		} else if b == "-" {
			PrintNbr(a - c)
			z01.PrintRune('\n')
			return
		} else if b == "*" {
			PrintNbr(a * c)
			z01.PrintRune('\n')
			return
		} else if b == "/" {
			if a == 0 || c == 0 {
				PrintStr("No division by 0\n")
				return
			}
			PrintNbr(a / c)
			z01.PrintRune('\n')
			return
		} else if b == "%" {
			if a == 0 || c == 0 {
				PrintStr("No modulo by 0\n")
				return
			}
			PrintNbr(a % c)
			z01.PrintRune('\n')
			return
		} else {
			return
		}
	}
	return
}

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for range s {
		c++
		z01.PrintRune(a[c-1])
	}
}

func CheckBA(s string) bool {
	i := 0
	sign := 0
	for i < len(s) {
		if s[i] == '-' || s[i] == '+' {
			sign++
		} else if s[i] > 57 || s[i] < 48 {
			return false
		}
		if (s[i] == '-' || s[i] == '+') && i > 0 {
			return false
		}
		i++
	}
	if sign > 1 {
		return false
	}
	return true
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

func Atoi(str string) int {
	neg := 1
	i := 0
	res := 0
	if str[0] == '-' {
		neg = -1
		i++
	}
	for i < len(str) {
		res = res*10 + int(str[i]-'0')
		i++
	}
	return res * neg
}
