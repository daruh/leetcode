package maximum_depth_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepth(t *testing.T) {

	/*
		Input: root = [3,9,20,null,null,15,7]
		Output: 3
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
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
		result := maxDepth(&tree1)
		//assert
		assert.Equal(t, 3, result)
	})

	/*
		Input: root = [1,null,2]
		Output: 2
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		}

		//act
		result := maxDepth(&tree1)
		//assert
		assert.Equal(t, 2, result)
	})
}
