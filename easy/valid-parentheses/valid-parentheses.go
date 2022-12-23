package valid_parentheses

type (
	stack struct {
		top    *node
		length int
	}
	node struct {
		value rune
		prev  *node
	}
)

// Len Return the number of items in the stack
func (r *stack) Len() int {
	return r.length
}

// Pop the top item of the stack and return it
func (r *stack) Pop() rune {
	if r.length == 0 {
		return 0
	}

	n := r.top
	r.top = n.prev
	r.length--
	return n.value
}

// Push a value onto the top of the stack
func (r *stack) Push(value rune) {
	n := &node{value, r.top}
	r.top = n
	r.length++
}

func isValid(s string) bool {

	result := true

	var (
		i int
		l = len(s) - 1
	)
	var brackets = map[rune]rune{'{': '}', '[': ']', '(': ')'}

	stack := stack{top: nil, length: 0}

	if len(s) == 0 {
		return true
	}

	for {
		bracket := s[i]
		if val, ok := brackets[rune(bracket)]; ok && val > 0 {
			stack.Push(rune(bracket))
		} else {
			popped := stack.Pop()

			if val, ok := brackets[popped]; !ok || val != rune(bracket) {
				result = false
				break
			}
		}
		i++
		if i > l {
			break
		}
	}

	return result && stack.length == 0
}
