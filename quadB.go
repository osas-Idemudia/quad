package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('/')

		if y >= 1 {
			for a := 0; a < (x - 2); a++ {
				z01.PrintRune('*')
			}
			if x >= 2 {
				z01.PrintRune(92)
			}
			z01.PrintRune('\n')

			for b := 0; b < (y - 2); b++ {
				z01.PrintRune('*')
				for a := 0; a < (x - 2); a++ {
					z01.PrintRune(' ')
				}
				if x >= 2 {
					z01.PrintRune('*')
				}
				z01.PrintRune('\n')
			}

			if y >= 2 {
				z01.PrintRune(92)
				for a := 0; a < (x - 2); a++ {
					z01.PrintRune('*')
				}
				if x >= 2 {
					z01.PrintRune('/')
				}
				z01.PrintRune('\n')
			}
		}
	}
}
