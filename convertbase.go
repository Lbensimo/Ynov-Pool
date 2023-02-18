package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return ItoaBase(AtoiBase(nbr, baseFrom), baseTo)
}

func ItoaBase(n int, base string) string {
	var dest []rune
	var temp []rune
	size := 0
	nbr := n
	for nbr >= len(base) {
		nbr /= len(base)
		size++
	}
	for i := 0; i <= size; i++ {
		temp = append(temp, rune(base[n%len(base)]))
		n /= len(base)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		dest = append(dest, temp[i])
	}
	return string(dest)
}
