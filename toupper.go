package piscine

func ToUpper(s string) string {
	tableau := []rune(s)
	for i := 0; i < len(tableau); i++ {
		if tableau[i] >= 'a' && tableau[i] <= 'z' {
			tableau[i] = tableau[i] - 32
		}
	}
	return string(tableau)
}
