package remove_duplicates_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	/*
		Input: nums = [1,1,2]
		Output: 2, nums = [1,2,_]
		Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
		It does not matter what you leave beyond the returned k (hence they are underscores).
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		nums := []int{1, 1, 2}
		//act
		result := removeDuplicates(nums)
		//assert
		assert.Equal(t, 2, result)
	})

	/*Input: nums = [0,0,1,1,1,2,2,3,3,4]
	Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
	Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		//act
		result := removeDuplicates(nums)
		//assert
		assert.Equal(t, 5, result)
	})

}
