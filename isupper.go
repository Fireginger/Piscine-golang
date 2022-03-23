package piscine

func IsUpper(s string) bool {
	tableau := []rune(s)
	for _, i := range tableau {
		if i < 64 || i > 91 {
			return false
		}
	}
	return true
}
