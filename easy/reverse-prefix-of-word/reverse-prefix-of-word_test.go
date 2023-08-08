package reverse_prefix_of_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePrefixWord(t *testing.T) {

	//Input: word = "abcdefd", ch = "d"
	//Output: "dcbaefd"
	//Explanation: The first occurrence of "d" is at index 3.
	//Reverse the part of word from 0 to 3 (inclusive), the resulting string is
	//"dcbaefd".
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		word := "abcdefd"
		ch := byte('d')
		//act
		result := reversePrefix(word, ch)
		//assert
		assert.Equal(t, "dcbaefd", result)
	})

	//Input: word = "xyxzxe", ch = "z"
	//Output: "zxyxxe"
	//Explanation: The first and only occurrence of "z" is at index 3.
	//Reverse the part of word from 0 to 3 (inclusive), the resulting string is
	//"zxyxxe"
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		word := "xyxzxe"
		ch := byte('z')
		//act
		result := reversePrefix(word, ch)
		//assert
		assert.Equal(t, "zxyxxe", result)
	})

	// Example 3:
	//
	//
	//Input: word = "abcd", ch = "z"
	//Output: "abcd"
	//Explanation: "z" does not exist in word.
	//You should not do any reverse operation, the resulting string is "abcd".
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		word := "abcd"
		ch := byte('z')
		//act
		result := reversePrefix(word, ch)
		//assert
		assert.Equal(t, "abcd", result)
	})
}
