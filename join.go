package piscine

func Join(strs []string, sep string) string {
	i := 1
	str := strs[0]
	for i < len(strs) {
		str += sep
		str += strs[i]
		i++
	}
	return str
}
