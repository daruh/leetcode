package search_insert_position

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	/*
		Input: nums = [1,3,5,6], target = 5
		Output: 2
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		nums := []int{1, 3, 5, 6}
		target := 5
		//act
		result := searchInsert(nums, target)
		//assert
		assert.Equal(t, 2, result)
	})

	/*Input: nums = [1,3,5,6], target = 2
	Output: 1
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		nums := []int{1, 3, 5, 6}
		target := 5
		//act
		result := searchInsert(nums, target)
		//assert
		assert.Equal(t, 2, result)
	})
	/*
		Input: nums = [1,3,5,6], target = 7
		Output: 4
	*/

	t.Run("Example 3", func(t *testing.T) {
		//arrange
		nums := []int{1, 3, 5, 6}
		target := 7
		//act
		result := searchInsert(nums, target)
		//assert
		assert.Equal(t, 4, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		//arrange
		nums := []int{2, 7, 8, 9, 10}
		target := 9
		//act
		result := searchInsert(nums, target)
		//assert
		assert.Equal(t, 3, result)
	})

}
