package binary_tree_inorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {

	/*
		Input: root = [1,null,2,3]
		Output: [1,3,2]
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		tree := TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		//act
		result := inorderTraversal(&tree)
		//assert
		assert.Equal(t, []int{1, 3, 2}, result)
	})

	/*
		Input: root = []
		Output: []
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		var tree *TreeNode
		//act
		result := inorderTraversal(tree)
		//assert
		assert.Equal(t, []int{}, result)
	})

	/*
		Input: root = [1]
		Output: [1]
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		tree := TreeNode{Val: 1}
		//act
		result := inorderTraversal(&tree)
		//assert
		assert.Equal(t, []int{1}, result)
	})

}
