package convert_sorted_array_to_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil || q == nil {
		return p == nil && q == nil
	}

	if p.Val != q.Val {
		return false
	}

	vLeft := isSameTree(p.Left, q.Left)
	if vLeft != true {
		return false
	}

	vRight := isSameTree(p.Right, q.Right)
	if vRight != true {
		return false
	}
	return true

}

func TestSortedArrayToBST(t *testing.T) {

	/*
		Input: nums = [-10,-3,0,5,9]
		Output: [0,-3,9,-10,null,5]
		Explanation: [0,-10,5,null,-3,null,9] is also accepted:
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		nums := []int{-10, -3, 0, 5, 9}
		outputTree1 := &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -10,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}

		//act
		result := sortedArrayToBST(nums)
		//assert
		assert.True(t, isSameTree(outputTree1, result))
	})

	/*
		Input: nums = [1,3]
		Output: [3,1]
		Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		nums := []int{1, 3}
		outputTree1 := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
		}

		//act
		result := sortedArrayToBST(nums)
		//assert
		assert.True(t, isSameTree(outputTree1, result))
	})

}
