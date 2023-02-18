package piscine

func WhatChar(c rune) int {
	if c <= 57 && c >= 48 {
		return 0
	} else if c <= 122 && c >= 97 {
		return 1
	} else if c <= 90 && c >= 65 {
		return 2
	} else {
		return 3
	}
}

func Capitalize(s string) string {
	r := []rune(s)
	i := 0
	wrd := false
	for i < len(s) {
		// C'est pas encore un mot, on avance jusqu'à ce que c'en soit un
		for !wrd && i < len(s) {
			if WhatChar(r[i]) == 0 || WhatChar(r[i]) == 2 {
				wrd = true
			} else if WhatChar(r[i]) == 1 {
				wrd = true
				r[i] -= 32
			} else {
				i++
			}
		}
		i++
		for wrd && i < len(s) {
			if WhatChar(r[i]) == 3 {
				wrd = false
			} else if WhatChar(r[i]) == 2 {
				r[i] += 32
			} else {
				i++
			}
		}
	}
	str := string(r)
	return str
}
