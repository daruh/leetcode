package merge_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {

	/*
		Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
		Output: [1,2,2,3,5,6]
		Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
		The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		m := 3
		n := 3
		//act
		merge(nums1, m, nums2, n)
		//assert
		assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
	})

	/*
		Input: nums1 = [1], m = 1, nums2 = [], n = 0
		Output: [1]
		Explanation: The arrays we are merging are [1] and [].
		The result of the merge is [1].
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		nums1 := []int{1}
		nums2 := []int{0}
		m := 1
		n := 0
		//act
		merge(nums1, m, nums2, n)
		//assert
		assert.Equal(t, []int{1}, nums1)
	})

	/*
		Input: nums1 = [0], m = 0, nums2 = [1], n = 1
		Output: [1]
		Explanation: The arrays we are merging are [] and [1].
		The result of the merge is [1].
		Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
	*/
	t.Run("Example 3", func(t *testing.T) {
		nums1 := []int{0}
		nums2 := []int{1}
		m := 0
		n := 1
		//act
		merge(nums1, m, nums2, n)
		//assert
		assert.Equal(t, []int{1}, nums1)
	})

	t.Run("Example 4", func(t *testing.T) {
		nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
		nums2 := []int{1, 2, 2}
		m := 6
		n := 3
		//act
		merge(nums1, m, nums2, n)
		//assert
		assert.Equal(t, []int{-1, 0, 0, 1, 2, 2, 3, 3, 3}, nums1)
	})
}
