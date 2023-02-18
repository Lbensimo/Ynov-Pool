package piscine

func SortIntegerTable(table []int) {
	i := 0
	for i < len(table)-1 {
		if table[i] > table[i+1] {
			swap := table[i]
			table[i] = table[i+1]
			table[i+1] = swap
			i = 0
		} else {
			i++
		}
	}
}
