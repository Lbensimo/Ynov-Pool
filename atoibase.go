package piscine

func FindBase(base string, r rune) int {
	for i := 0; i < len(base); i++ {
		if rune(base[i]) == r {
			return i
		}
	}
	return 0
}

func AtoiBase(s string, base string) int {
	if !CB(base) {
		return 0
	}
	i := 0
	result := 0
	r := []rune(s)
	for i < len(s) {
		result = result*len(base) + FindBase(base, r[i])
		i++
	}
	return result
}
