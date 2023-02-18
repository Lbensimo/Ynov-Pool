package piscine

func BasicJoin(elems []string) string {
	i := 1
	str := elems[0]
	for i < len(elems) {
		str += elems[i]
		i++
	}
	return str
}
