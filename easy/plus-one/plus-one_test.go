package plus_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {

	/*
		Input: digits = [1,2,3]
		Output: [1,2,4]
		Explanation: The array represents the integer 123.
		Incrementing by one gives 123 + 1 = 124.
		Thus, the result should be [1,2,4].
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		digits := []int{1, 2, 3}
		//act
		result := plusOne(digits)
		//assert
		assert.Equal(t, []int{1, 2, 4}, result)
	})

	/*
		Input: digits = [4,3,2,1]
		Output: [4,3,2,2]
		Explanation: The array represents the integer 4321.
		Incrementing by one gives 4321 + 1 = 4322.
		Thus, the result should be [4,3,2,2].
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		digits := []int{4, 3, 2, 1}
		//act
		result := plusOne(digits)
		//assert
		assert.Equal(t, []int{4, 3, 2, 2}, result)
	})

	/*
		Input: digits = [9]
		Output: [1,0]
		Explanation: The array represents the integer 9.
		Incrementing by one gives 9 + 1 = 10.
		Thus, the result should be [1,0].
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		digits := []int{9}
		//act
		result := plusOne(digits)
		//assert
		assert.Equal(t, []int{1, 0}, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		//arrange
		digits := []int{9, 9, 9}
		//act
		result := plusOne(digits)
		//assert
		assert.Equal(t, []int{1, 0, 0, 0}, result)
	})
}
