package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	i := nb - 1
	for i > 1 {
		nb *= i
		i--
	}
	return nb
}
