package piscine

func NRune(s string, n int) rune {
	tableau := []rune(s)
	longueur := len(tableau)
	if n > longueur || n <= 0 {
		return 0
	} else {
		return (tableau[n-1])
	}
}
