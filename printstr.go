package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	sString := s
	for i := 0; i < len(sString); i++ {
		z01.PrintRune(rune(sString[i]))
	}
}
