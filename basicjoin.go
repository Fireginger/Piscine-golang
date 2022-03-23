package piscine

func BasicJoin(elems []string) string {
	var z string
	for i := 0; i < len(elems); i++ {
		z = z + elems[i]
	}
	return z
}
