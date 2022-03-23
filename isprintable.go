package piscine

func IsPrintable(s string) bool {
	tableau := []rune(s)
	for _, i := range tableau {
		if i > 126 || i < 31 {
			return false
		}
	}
	return true
}
