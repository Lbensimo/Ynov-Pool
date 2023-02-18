package piscine

func BasicAtoi2(s string) int {
	if !CheckBA(s) {
		return 0
	}
	i := 0
	for i <= len(s) {
		if (s[i] < 48 || s[i] > 57) && i != 0 {
			return 0
		}
	}
	return BasicAtoi(s)
}
