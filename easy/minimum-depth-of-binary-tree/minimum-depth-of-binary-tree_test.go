package binary_tree_inorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	/*
		Input: root = [3,9,20,null,null,15,7]
		Output: 2
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		tree := TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}
		//act
		result := minDepth(&tree)
		//assert
		assert.Equal(t, 2, result)
	})

	/*
		Input: root = [2,null,3,null,4,null,5,null,6]
		Output: 5
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		tree := TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				}}}
		//act
		result := minDepth(&tree)
		//assert
		assert.Equal(t, 5, result)
	})

}
