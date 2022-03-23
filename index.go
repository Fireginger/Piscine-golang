package piscine

func Index(s string, toFind string) int {
	count := 0
	a := 0
	b := []rune(toFind)
	c := []rune(s)
	for q, y := range s {
		if count == 0 {
			if y == b[0] {
				count = q
				a++
				break
			} else if q == len(s)-1 && a == 0 {
				return -1
			}
		}
	}
	for i := 0; i < len(b); i++ {
		if c[count+i] == b[i] {
		} else {
			return -1
		}
	}
	return count
}
