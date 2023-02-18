package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	i := 0
	for i < len(a)-1 {
		if f(a[i], a[i+1]) > 0 {
			swap := a[i]
			a[i] = a[i+1]
			a[i+1] = swap
			i = 0
		} else {
			i++
		}
	}
}
