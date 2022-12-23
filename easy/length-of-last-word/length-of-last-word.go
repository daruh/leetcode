package length_of_last_word

func lengthOfLastWord(s string) int {

	l := len(s) - 1
	var lenLast int
	shouldBreak := false

	for i := l; i >= 0; i-- {

		if rune(s[i]) == (' ') {
			if shouldBreak {
				break
			}
			continue
		}

		if rune(s[i]) != (' ') {
			lenLast++
			shouldBreak = true
		}
	}

	return lenLast
}
