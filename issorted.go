package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	sorted := false
	detros := false
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			detros = true
		} else if f(a[i], a[i+1]) < 0 {
			sorted = true
		}
	}
	if sorted && detros {
		return false
	}
	return true
}
