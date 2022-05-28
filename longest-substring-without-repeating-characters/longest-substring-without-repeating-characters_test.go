package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {

	/*
		Input: s = "abcabcbb"
		Output: 3
		Explanation: The answer is "abc", with the length of 3.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		s := "abcabcbb"
		//act
		result := lengthOfLongestSubstring(s)
		//assert
		assert.Equal(t, 3, result)
	})

	/*Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.
	*/

	t.Run("Example 2", func(t *testing.T) {
		//arrange
		s := "bbbbb"
		//act
		result := lengthOfLongestSubstring(s)
		//assert
		assert.Equal(t, 1, result)
	})

	/*Input: s = "pwwkew"
	  Output: 3
	  Explanation: The answer is "wke", with the length of 3.
	  Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		s := "pwwkew"
		//act
		result := lengthOfLongestSubstring(s)
		//assert
		assert.Equal(t, 3, result)
	})
}
