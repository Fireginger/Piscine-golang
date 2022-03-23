package piscine

var c *int

func UltimateDivMod(a *int, b *int) {
	c := *a
	*a = *a / *b
	*b = c % *b
}
