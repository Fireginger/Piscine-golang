package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, r := range os.Args[0] {
		if i == 0 || i == 1 && (os.Args[0][i]) == 46 || (os.Args[0][i]) == 47 {
			continue
		}
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
