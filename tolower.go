package piscine

func ToLower(s string) string {
	c := []rune(s)
	i := 0
	for i < len(s) {
		if s[i] <= 90 && s[i] >= 65 {
			c[i] += 32
		}
		i++
	}
	return string(c)
}
