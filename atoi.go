package piscine

func Atoi(s string) int {
	if !CheckBA(s) {
		return 0
	}
	return TrimAtoi(s)
}
