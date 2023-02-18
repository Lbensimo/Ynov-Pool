package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for range s {
		c++
		z01.PrintRune(a[c-1])
	}
	z01.PrintRune('\n')
}

func StrLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}

func Index(s string, ToFind string) bool {
	i := 0
	i2 := 0
	for i < StrLen(s) {
		for s[i+i2] == ToFind[i2] {
			i2++
			if i2 == StrLen(ToFind) {
				return true
			}
		}
		i2 = 0
		i++
	}
	return false
}

func Sort(s []rune) {
	for i := 0; i < len(s)-1; {
		if s[i] > s[i+1] {
			swap := s[i]
			s[i] = s[i+1]
			s[i+1] = swap
			i = 0
		} else {
			i++
		}
	}
	PrintStr(string(s))
}

func main() {
	a := os.Args
	var tabr []rune
	var str []rune
	if len(a) <= 1 || a[1] == "-h" || a[1] == "--help" {
		fmt.Print("--insert\n", "  -i\n", "\t This flag inserts the string into the string passed as argument.\n")
		fmt.Print("--order\n", "  -o\n", "\t This flag will behave like a boolean, if it is called it will order the argument.\n")
		return
	}
	var order bool
	for i := 1; i < 2; i++ {
		for i2 := 1; i2 < len(a); i2++ {
			if a[i2] == "--order" || a[i2] == "-o" {
				order = true
			}
		}
		for i2 := 1; i2 < len(a); i2++ {
			if Index(a[i2], "--insert") {
				for z := 9; z < len(a[i2]); z++ {
					tabr = append(tabr, rune(a[i2][z]))
				}
			} else if Index(a[i2], "-i") {
				for z := 3; z < len(a[i2]); z++ {
					tabr = append(tabr, rune(a[i2][z]))
				}
			}
		}
		for i3 := 1; i3 < len(a); i3++ {
			if a[i3] != "-o" && a[i3] != "--order" &&
				!Index(a[i3], "-i") && !Index(a[i3], "--insert") {
				str = []rune(a[i3])
			}
		}
	}
	strr := string(str)
	strt := string(tabr)
	strr += strt
	str = []rune(strr)
	if order {
		Sort(str)
		return
	} else {
		PrintStr(strr)
	}
}
