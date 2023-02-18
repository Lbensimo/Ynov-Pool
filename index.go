package piscine

func Index(s string, ToFind string) int {
	i := 0
	i2 := 0
	for i < StrLen(s) {
		for s[i+i2] == ToFind[i2] {
			i2++
			if i2 == StrLen(ToFind) {
				return i
			}
		}
		i2 = 0
		i++
	}
	return -1
}
