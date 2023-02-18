package piscine

func AlphaCount(s string) int {
	count := 0
	i := 0
	for i < len(s) {
		if (s[i] <= 90 && s[i] >= 65) || (s[i] <= 122 && s[i] >= 97) {
			count++
		}
		i++
	}
	return count
}
