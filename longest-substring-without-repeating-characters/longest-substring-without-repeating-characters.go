package longest_substring_without_repeating_characters

type char uint8
type runeSet map[char]struct{}

func (s runeSet) add(r char) {
	s[r] = struct{}{}
}

func (s runeSet) remove(r char) {
	delete(s, r)
}

func (s runeSet) has(r char) bool {
	_, ok := s[r]
	return ok
}

func lengthOfLongestSubstring(s string) int {

	runeSet := runeSet{}
	var (
		i             int
		j             int
		maxResult     int
		prevMaxResult int
		rangeLen      = len(s)
	)

	for {
		if len(s) == 0 {
			return maxResult
		}

		if i >= rangeLen || j >= rangeLen {
			break
		}
		if has := runeSet.has(char(s[j])); has == false {
			runeSet.add(char(s[j]))
			maxResult++
			j++
		} else {
			runeSet.remove(char(s[i]))
			i++
			maxResult--
		}
		if maxResult > prevMaxResult {
			prevMaxResult = maxResult
		}
	}

	return prevMaxResult
}
