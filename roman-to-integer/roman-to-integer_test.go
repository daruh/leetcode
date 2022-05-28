package roman_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInteger(t *testing.T) {

	/*Input: s = "III"
	Output: 3
	Explanation: III = 3.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		s := "III"
		//act
		result := romanToInt(s)
		//assert
		assert.Equal(t, 3, result)
	})

	/*Input: s = "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		s := "LVIII"
		//act
		result := romanToInt(s)
		//assert
		assert.Equal(t, 58, result)
	})

	/*Input: s = "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		s := "MCMXCIV"
		//act
		result := romanToInt(s)
		//assert
		assert.Equal(t, 1994, result)
	})
}
