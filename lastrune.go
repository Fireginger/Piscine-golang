package piscine

func LastRune(s string) rune {
	tableau := []rune(s)
	longueur := len(tableau)
	return (tableau[longueur-1])
}
