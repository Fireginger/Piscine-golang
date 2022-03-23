package piscine

func MakeRange(min, max int) []int {
	if min > max || min == max {
		return nil
	}
	answer := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		answer[i] = i + min
	}
	return answer
}
