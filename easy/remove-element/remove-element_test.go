package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	/*
		Input: nums = [3,2,2,3], val = 3
		Output: 2, nums = [2,2,_,_]
		Explanation: Your function should return k = 2, with the first two elements of nums being 2.
		It does not matter what you leave beyond the returned k (hence they are underscores).
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		nums := []int{3, 2, 2, 3}
		val := 3
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 2, result)
	})

	/*Input: nums = [0,1,2,2,3,0,4,2], val = 2
	Output: 5, nums = [0,1,4,0,3,_,_,_]
	Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
	Note that the five elements can be returned in any order.
	It does not matter what you leave beyond the returned k (hence they are underscores)
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		val := 2
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 5, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		//arrange
		nums := []int{2}
		val := 2
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 0, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		//arrange
		nums := []int{2, 2, 2}
		val := 2
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 0, result)
	})

	t.Run("Example 5", func(t *testing.T) {
		//arrange
		nums := []int{2, 2, 2, 3, 4, 5}
		val := 2
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 3, result)
	})
	t.Run("Example 6", func(t *testing.T) {
		//arrange
		nums := []int{2, 1, 1, 3, 4, 2}
		val := 2
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 4, result)
	})

	t.Run("Example 7", func(t *testing.T) {
		//arrange
		nums := []int{4, 5}
		val := 5
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 1, result)
	})

	t.Run("Example 8", func(t *testing.T) {
		//arrange
		nums := []int{2, 2, 3}
		val := 2
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 1, result)
	})

	t.Run("Example 9", func(t *testing.T) {
		//arrange
		nums := []int{3, 3}
		val := 5
		//act
		result := removeElement(nums, val)
		//assert
		assert.Equal(t, 2, result)
	})

}
