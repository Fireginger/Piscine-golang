package piscine

func ToLower(s string) string {
	tableau := []rune(s)
	for i := 0; i < len(tableau); i++ {
		if tableau[i] >= 'A' && tableau[i] <= 'Z' {
			tableau[i] = tableau[i] + 32
		}
	}
	return string(tableau)
}
