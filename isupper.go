package piscine

func IsUpper(s string) bool {
	i := 0
	for i < len(s) {
		if s[i] > 90 || s[i] < 65 {
			return false
		}
		i++
	}
	return true
}
