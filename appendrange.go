package piscine

func AppendRange(min, max int) []int {
	var tab []int
	if max <= min {
		return tab
	}
	for i := 0; i < max-min; i++ {
		tab = append(tab, min+i)
	}
	return tab
}
