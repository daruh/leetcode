package plus_one

func plusOne(digits []int) []int {

	l := len(digits) - 1

	for i := l; i >= 0; i-- {
		val := digits[i]
		incVal := (val + 1) % 10
		digits[i] = incVal
		if incVal != 0 {
			break
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
