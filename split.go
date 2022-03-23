package piscine

func Split(s, sep string) []string {
	var result []string
	tableau := []rune(s)
	tableau2 := []rune(sep)
	phrase := ""
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == tableau2[0] && tableau[i+1] == tableau2[1] {
			result = append(result, phrase)
			phrase = ""
			i += len(sep) - 1
		} else {
			phrase = phrase + string(tableau[i])
		}
		if i+1 == len(tableau) {
			result = append(result, phrase)
		}
	}
	return result
}
