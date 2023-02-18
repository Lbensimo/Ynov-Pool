package piscine

func IterativePower(nb int, power int) int {
	lock := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for power > 1 {
		nb *= lock
		power--
	}
	return nb
}
