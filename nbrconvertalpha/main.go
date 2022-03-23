package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	o := 0
	c := 0
	var up bool
	if len(os.Args) == 1 {
	} else {
		for i := 0; i < len(os.Args); i++ {
			if os.Args[1] == "--upper" && i == 0 {
				up = true
				i += 2
			} else if i == 0 {
				i++
			}
			for _, word := range os.Args[i] {
				if word >= 48 && word <= 57 {
					for j := '0'; j < word; j++ {
						c++
					}
					o = o*10 + c
					c = 0
				} else {
					o = 0
					break
				}
			}
			if o == 0 || o > 26 {
				z01.PrintRune(' ')
			} else if up == true {
				z01.PrintRune(rune(o + 64))
			} else {
				z01.PrintRune(rune(o + 96))
			}
			o = 0
		}
		z01.PrintRune('\n')
	}
}
