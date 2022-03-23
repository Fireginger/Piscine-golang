package piscine

func BasicAtoi(s string) int {
	o_number := 0
	for _, c := range s {
		o_number = o_number*10 + int(c-48)
	}
	return o_number
}
