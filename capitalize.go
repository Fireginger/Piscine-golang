package piscine

func Capitalize(s string) string {
	tableau := []rune(s)
	if s[0] >= 'a' && s[0] <= 'z' {
		tableau[0] = tableau[0] - 32
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] >= 'A' && s[i-1] <= 'Z' || s[i-1] >= 'a' && s[i-1] <= 'z' || s[i-1] >= '0' && s[i-1] <= '9' {
			if s[i] >= 'A' && s[i] <= 'Z' {
				tableau[i] = tableau[i] + 32
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			tableau[i] = tableau[i] - 32
		}
	}
	return string(tableau)
}
