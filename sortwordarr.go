package piscine

func SortWordArr(a []string) {
	i := 0
	for i < len(a)-1 {
		if Compare(a[i], a[i+1]) > 0 {
			swap := a[i]
			a[i] = a[i+1]
			a[i+1] = swap
			i = 0
		} else {
			i++
		}
	}
}
