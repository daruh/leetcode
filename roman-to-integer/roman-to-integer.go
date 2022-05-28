package roman_to_integer

type char uint8

type roman struct {
	rmap map[char]int
}

func NewRoman() *roman {
	r := roman{
		rmap: map[char]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000},
	}
	return &r
}

func (r *roman) val(s char) int {
	return r.rmap[s]
}

func romanToInt(s string) int {

	var (
		result    int
		l         = len(s) - 1
		i         int
		romanConv = NewRoman()
	)
	for {
		currentValue := romanConv.val(char(s[i]))

		if i == l {
			result += currentValue
			break
		}

		forwardValue := romanConv.val(char(s[i+1]))
		if currentValue >= forwardValue {
			result += currentValue
			i++
			continue
		}

		if currentValue < forwardValue {
			result += forwardValue - currentValue
			i += 2
		}
		if i > l {
			break
		}
	}

	return result
}
