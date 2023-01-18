package path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathSum(t *testing.T) {

	/*
		Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
		Output: true
		Explanation: The root-to-leaf path with the target sum is shown.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		tree := TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 11,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		}
		targetSum := 22
		//act
		result := hasPathSum(&tree, targetSum)
		//assert
		assert.Equal(t, true, result)
	})

	/*
		Input: root = [1,2,2,3,3,null,null,4,4]
		Output: false
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		tree := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		targetSum := 5
		//act
		result := hasPathSum(&tree, targetSum)
		//assert
		assert.Equal(t, false, result)
	})

	/*
		Input: root = []
		Output: true
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		tree := TreeNode{}
		targetSum := 0
		//act
		result := hasPathSum(&tree, targetSum)
		//assert
		assert.Equal(t, false, result)
	})
}
