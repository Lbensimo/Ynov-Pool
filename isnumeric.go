package piscine

func IsNumeric(s string) bool {
	i := 0
	for i < len(s) {
		if s[i] > 57 || s[i] < 48 {
			return false
		}
		i++
	}
	return true
}
