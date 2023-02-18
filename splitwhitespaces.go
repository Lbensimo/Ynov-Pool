package piscine

func FindWord(s string, i int) string {
	str := make([]rune, len(s))
	i2 := 0
	for i < len(s) && s[i] != ' ' {
		str[i2] = rune(s[i])
		i++
		i2++
	}
	return string(str)
}

func CountSpace(s string) int {
	i := 0
	count := 1
	if len(s) < 1 {
		return 0
	}
	if s[0] == ' ' {
		i++
	}
	for i < len(s) {
		if s[i] == ' ' {
			count++
			for i < len(s) && s[i] == ' ' {
				i++
			}
		}
		i++
	}
	if len(s) >= 2 {
		if s[i-2] == ' ' {
			count--
		}
	}
	return count
}

func SplitWhiteSpaces(s string) []string {
	i := 0
	i2 := 0
	if CountSpace(s) == 0 {
		var merde []string
		return merde
	}
	for s[i] == ' ' && i < len(s) {
		i++
	}
	str := make([]string, CountSpace(s))
	for i < len(s) {
		str[i2] += string(s[i])
		i++
		if i < len(s) && s[i] == ' ' {
			i2++
			i++
			for i < len(s) && s[i] == ' ' {
				i++
			}
		}
	}
	return str
}
