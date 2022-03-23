package piscine

func Atoi(s string) int {
	o := 0
	c := 0
	a_s := []rune(s)
	if len(s) == 0 {
	} else {
		for _, word := range a_s {
			if word >= 48 && word <= 57 {
				for i := '0'; i < word; i++ {
					c++
				}
				o = o*10 + c
				c = 0
			} else {
				o = 0
				break
			}
		}
		if a_s[0] == 43 {
			for index, word := range a_s {
				if index == 0 {
				} else if word >= 48 && word <= 57 {
					for i := '0'; i < word; i++ {
						c++
					}
					o = o*10 + c
					c = 0
				} else {
					o = 0
					break
				}
			}
		}
		if a_s[0] == 45 {
			for index, word := range a_s {
				if index == 0 {
				} else if word >= 48 && word <= 57 {
					for i := '0'; i < word; i++ {
						c++
					}
					o = o*10 - c
					c = 0
				} else {
					o = 0
					break
				}
			}
		}
	}
	return o
}
