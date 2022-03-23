package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	tableau := []rune(str)
	phrase := ""
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == ' ' || tableau[i] == '\n' || tableau[i] == '	' {
			if tableau[i+1] == ' ' || tableau[i+1] == '\n' || tableau[i+1] == '	' {
			} else {
				result = append(result, phrase)
				phrase = ""
			}
		} else {
			phrase = phrase + string(tableau[i])
		}
		if i+1 == len(tableau) {
			result = append(result, phrase)
		}
	}
	return result
}
