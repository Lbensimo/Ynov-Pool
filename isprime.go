package piscine

func IsPrime(nb int) bool {
	i := 2
	if nb < 2 {
		return false
	}
	for i < 46341 && i <= nb/2 {
		if nb%i == 0 {
			return false
		}
		i++
	}
	return true
}
