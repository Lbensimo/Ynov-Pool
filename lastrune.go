package piscine

func LastRune(s string) rune {
	c := []rune(s)
	return c[len(s)-1]
}
