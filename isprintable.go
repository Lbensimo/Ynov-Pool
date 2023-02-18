package piscine

func IsPrintable(s string) bool {
	i := 0
	for i < len(s) {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
		i++
	}
	return true
}
