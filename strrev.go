package piscine

func StrRev(s string) string {
	var rev string
	i := len(s) - 1
	for i >= 0 {
		rev += string(s[i])
		i--
	}
	return rev
}
