package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	per := 0
	i := 1
	if len(os.Args) <= 1 {
		return
	}
	if os.Args[i] == "--upper" {
		per = 32
		i++
	} else {
		per = 0
	}
	for i < len(os.Args) {
		if len(os.Args[i]) > 1 && len(os.Args[i]) < 3 {
			if os.Args[i][0] == '1' && (os.Args[i][1] <= 57 && os.Args[i][1] >= 48) {
				z01.PrintRune(('j') + rune(os.Args[i][1]-48) - rune(per))
			} else {
				if (os.Args[i][0] == '2') &&
					(os.Args[i][1] < 54 && os.Args[i][1] >= 48) {
					z01.PrintRune(('t') + rune(os.Args[i][1]-48) - rune(per))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			if os.Args[i][0] <= 57 && os.Args[i][0] > 48 {
				z01.PrintRune(('a' - 1) + rune(os.Args[i][0]-48) - rune(per))
			} else {
				z01.PrintRune(' ')
			}
		}
		i++
	}
	z01.PrintRune('\n')
}
