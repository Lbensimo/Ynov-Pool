package piscine

func Sqrt(nb int) int {
	i := 2
	if nb == 1 {
		return 1
	}
	if nb < 4 {
		return 0
	}
	for i < 46341 && i <= nb/2 {
		if nb == i*i {
			return i
		}
		i++
	}
	return 0
}
