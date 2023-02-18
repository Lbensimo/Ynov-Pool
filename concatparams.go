package piscine

func BasicJoin2(elems []string) string {
	i := 1
	str := elems[0]
	for i < len(elems) {
		str += "\n"
		str += elems[i]
		i++
	}
	return str
}

func ConcatParams(args []string) string {
	return BasicJoin2(args)
}
