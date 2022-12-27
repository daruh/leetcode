package same_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	/*
		Input: root = [1,2,2,3,4,4,3]
		Output: true
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				}, Right: &TreeNode{
					Val: 3,
				},
			},
		}
		//act
		result := isSymmetric(&tree1)
		//assert
		assert.Equal(t, true, result)
	})

	/*
		Input: root = [1,2,2,null,3,null,3]
		Output: false
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		}
		//act
		result := isSymmetric(&tree1)
		//assert
		assert.Equal(t, false, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
			},
		}
		//act
		result := isSymmetric(&tree1)
		//assert
		assert.Equal(t, false, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}
		//act
		result := isSymmetric(&tree1)
		//assert
		assert.Equal(t, true, result)
	})
}
