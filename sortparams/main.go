package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Compare(a, b string) int {
	i := 0
	for i < len(a)-1 && i < len(b)-1 {
		if a[i] == b[i] {
			i++
		} else if a[i] > b[i] {
			return 1
		} else {
			return -1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(b) < len(a) {
		return 1
	} else {
		return 0
	}
}

func PrintStr(s string) {
	a := []rune(s)
	c := 0
	for c < len(a) {
		c++
		z01.PrintRune(a[c-1])
	}
	z01.PrintRune('\n')
}

func main() {
	var stab []string
	for i := 1; i < len(os.Args); i++ {
		stab = append(stab, os.Args[i])
	}
	for i := 0; i < len(stab)-1; {
		if stab[i] > stab[i+1] {
			swap := stab[i]
			stab[i] = stab[i+1]
			stab[i+1] = swap
			i = 0
		} else {
			i++
		}
	}
	for i := 0; i < len(stab); i++ {
		PrintStr(stab[i])
	}
}
