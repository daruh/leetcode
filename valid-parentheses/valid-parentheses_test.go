package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	/*Input: s = "()"
	Output: true
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		s := "()"
		//act
		var result = isValid(s)
		//assert
		assert.Equal(t, true, result)
	})

	/*Input: s = "()[]{}"
	Output: true
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		s := "()[]{}"
		//act
		var result = isValid(s)
		//assert
		assert.Equal(t, true, result)
	})

	/*Input: s = "(]"
	Output: false
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		s := "(]"
		//act
		var result = isValid(s)
		//assert
		assert.Equal(t, false, result)
	})
}
