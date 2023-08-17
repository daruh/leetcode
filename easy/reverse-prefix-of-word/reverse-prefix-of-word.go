package reverse_prefix_of_word

func reversePrefix(word string, ch byte) string {

	var newStr string
	j := -1

	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			j = i
			break
		}
	}

	if j != -1 {
		k := j
		for k >= 0 {
			newStr = newStr + string(word[k])
			k--
		}

		for l := j + 1; l < len(word); l++ {
			newStr = newStr + string(word[l])
		}
	} else {
		return word
	}

	return newStr
}
