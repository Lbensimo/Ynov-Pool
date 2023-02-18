package piscine

func CheckBA(s string) bool {
	i := 0
	sign := 0
	for i < len(s) {
		if s[i] == '-' || s[i] == '+' {
			sign++
		} else if s[i] > 57 || s[i] < 48 {
			return false
		}
		if (s[i] == '-' || s[i] == '+') && i > 0 {
			return false
		}
		i++
	}
	if sign > 1 {
		return false
	}
	return true
}

func BasicAtoi(s string) int {
	if CheckBA(s) {
		return TrimAtoi(s)
	}
	return 0
}
