package piscine

func IsAlpha(s string) bool {
	tableau := []rune(s)
	for _, i := range tableau {
		if i > 123 || i < 47 {
			return false
		}
	}
	return true
}
