package piscine

func IsNumeric(s string) bool {
	tableau := []rune(s)
	for _, i := range tableau {
		if i > 58 || i < 47 {
			return false
		}
	}
	return true
}
