package piscine

func IsNum(s rune) bool {
	i := 0
	if s > 57 || s < 48 {
		return false
	}
	i++
	return true
}

func TrimAtoi(s string) int {
	i := 0
	result := 0
	r := []rune(s)
	neg := 1
	for i < len(s) && !IsNum(r[i]) {
		if s[i] == '-' {
			neg *= -1
		}
		i++
	}
	for i < len(s) {
		if IsNum(r[i]) {
			result = result*10 + int(s[i]-'0')
		}
		i++
	}
	return result * neg
}
