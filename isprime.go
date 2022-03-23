package piscine

func IsPrime(nb int) bool {
	result := 0
	sqrt := 0
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	for a := 1; result <= nb; a++ {
		result = a * a
		sqrt++
	}
	for i := 2; i <= sqrt; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
