package piscine

func Compare(a, b string) int {
	i := 0
	for i < len(a) && i < len(b) {
		if a[i] == b[i] {
			i++
		} else if a[i] > b[i] {
			return 1
		} else {
			return -1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(b) < len(a) {
		return 1
	} else {
		return 0
	}
}
