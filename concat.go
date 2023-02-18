package piscine

func Concat(str1 string, str2 string) string {
	str1 += str2
	concat := []rune(str1)
	return string(concat)
}
