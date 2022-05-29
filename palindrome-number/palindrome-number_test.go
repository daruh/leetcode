package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {

	/*Input: x = 121
	Output: true
	Explanation: 121 reads as 121 from left to right and from right to left.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		var x = 121
		//act
		var result = isPalindrome(x)
		//assert
		assert.Equal(t, true, result)
	})

	/*Input: x = -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		var x = 121
		//act
		var result = isPalindrome(x)
		//assert
		assert.Equal(t, false, result)
	})

	/*Input: x = 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		var x = 121
		//act
		var result = isPalindrome(x)
		//assert
		assert.Equal(t, false, result)
	})
}
