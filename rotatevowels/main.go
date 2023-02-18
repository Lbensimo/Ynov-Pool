package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicJoin(elems []string) string {
	i := 1
	str := elems[0]
	for i < len(elems) {
		str += " "
		str += elems[i]
		i++
	}
	return str
}

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for range s {
		c++
		z01.PrintRune(a[c-1])
	}
}

func Vowels(char rune) bool {
	vowels := "aeiouAEIOU"
	for i := 0; i < len(vowels); i++ {
		if char == rune(vowels[i]) {
			return true
		}
	}
	return false
}

func main() {
	i2 := 0
	if len(os.Args) <= 1 {
		z01.PrintRune('\n')
		return
	}
	str := []rune(BasicJoin(os.Args[1:]))
	for i := 0; i < len(str)/2 && i2 < len(str)/2; {
		if Vowels(str[i]) {
			for !Vowels(str[len(str)-i2-1]) && i2 < len(str)/2 {
				i2++
			}
			if Vowels(str[len(str)-i2-1]) {
				swap := str[i]
				str[i] = str[len(str)-i2-1]
				str[len(str)-i2-1] = swap
				i2++
				i++
			}
		} else if !Vowels(str[i]) && !Vowels(str[len(str)-i2-1]) {
			i++
		}
		if Vowels(str[len(str)-i2-1]) {
			for !Vowels(str[i]) && i < len(str)/2 {
				i++
			}
			if Vowels(str[i]) {
				swap := str[i]
				str[i] = str[len(str)-i2-1]
				str[len(str)-i2-1] = swap
				i++
				i2++
			}
		} else if !Vowels(str[i]) && !Vowels(str[len(str)-i2-1]) {
			i2++
		}
	}
	PrintStr(string(str))
	z01.PrintRune(' ')
	z01.PrintRune('\n')
}
