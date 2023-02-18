package piscine

func MakeRange(min, max int) []int {
	var tabr []int
	if min >= max {
		return tabr
	}
	tab := make([]int, max-min)
	i2 := 0
	for i := min; i < max; i++ {
		tab[i2] = i
		i2++
	}
	return tab
}
