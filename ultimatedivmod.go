package piscine

func UltimateDivMod(a *int, b *int) {
	swap := *a
	*a = *a / *b
	*b = swap % *b
}
