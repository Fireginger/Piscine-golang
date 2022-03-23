package piscine

func BasicAtoi2(s string) int {
	o_number := 0
	c := 0
	as := []rune(s)
	for _, word := range as {
		if word >= 48 && word <= 57 {
			for i := '0'; i < word; i++ {
				c++
			}
		} else {
			o_number = 0
			break
		}
		o_number = o_number*10 + c
		c = 0
	}
	return o_number
}
