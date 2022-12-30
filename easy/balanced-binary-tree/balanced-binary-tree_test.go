package binary_tree_inorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	/*
		Input: root = [3,9,20,null,null,15,7]
		Output: true
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
		result := isBalanced(&tree)
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
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					}, Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		}
		//act
		result := isBalanced(&tree)
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
		//act
		result := isBalanced(&tree)
		//assert
		assert.Equal(t, true, result)
	})

}
