package piscine

func IsLower(s string) bool {
	tableau := []rune(s)
	for _, i := range tableau {
		if i < 96 || i > 123 {
			return false
		}
	}
	return true
}
