package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tamere []int
	o := n
	c := 0
	e := 1
	r := 0
	if n < 0 {
	} else {
		for o > 9 {
			o = o / 10
			c++
		}
		o = n
		for j := 0; j <= c; j++ {
			if o > 9 {
				for o > 9 {
					o = o / 10
					e = e * 10
				}
				tamere = append(tamere, o)
				r = r + o*e
				o = n - r
				if o < e/10 {
					tamere = append(tamere, 0)
				}
				e = 1
			} else {
				tamere = append(tamere, o)
			}
			if tamere[j] == 0 {
				j++
			}
		}
		for k := 0; k <= 9; k++ {
			for l := 0; l <= len(tamere)-1; l++ {
				if tamere[l] == k {
					z01.PrintRune(rune(tamere[l] + 48))
				}
			}
		}
	}
}
