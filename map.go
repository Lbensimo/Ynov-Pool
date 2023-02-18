package piscine

func Map(f func(int) bool, a []int) []bool {
	mab := make([]bool, len(a))
	i2 := 0
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			mab[i2] = true
		}
		i2++
	}
	return mab
}
