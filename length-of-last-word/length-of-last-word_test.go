package length_of_last_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {

	/*
		Input: s = "Hello World"
		Output: 5
		Explanation: The last word is "World" with length 5.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		s := "Hello World"
		//act
		result := lengthOfLastWord(s)
		//assert
		assert.Equal(t, 5, result)
	})

	/*
		Input: s = "   fly me   to   the moon  "
		Output: 4
		Explanation: The last word is "moon" with length 4.
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		s := "   fly me   to   the moon  "
		//act
		result := lengthOfLastWord(s)
		//assert
		assert.Equal(t, 4, result)
	})

	/*
		Input: s = "luffy is still joyboy"
		Output: 6
		Explanation: The last word is "joyboy" with length 6.
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		s := "luffy is still joyboy"
		//act
		result := lengthOfLastWord(s)
		//assert
		assert.Equal(t, 6, result)
	})

}
