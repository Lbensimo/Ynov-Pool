package piscine

func Split(s, sep string) []string {
	var str []string
	var temp string
	var nil string
	i := 0
	for i < len(s) {
		if IsSep(s, sep, i) {
			if temp != nil {
				str = append(str, temp)
			}
			i += len(sep)
			temp = nil
		}
		if i < len(s) && !IsSep(s, sep, i) {
			temp += string(s[i])
		}
		i++
	}
	if i+len(sep)-1 >= len(s) && temp != nil {
		str = append(str, temp)
	}
	return str
}

func IsSep(s, sep string, i int) bool {
	i2 := 0
	for i < len(s) && i2 < len(sep) && s[i] == sep[i2] {
		i2++
		i++
	}
	return i2 == len(sep)
}
