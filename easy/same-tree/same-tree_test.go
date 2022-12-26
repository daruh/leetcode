package same_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	/*
		    Input: p = [1,2,3], q = [1,2,3]
			Output: true
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		tree2 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}
		//act
		result := isSameTree(&tree1, &tree2)
		//assert
		assert.Equal(t, true, result)
	})

	/*
		    Input: p = [1,2], q = [1,null,2]
			Output: true
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}
		tree2 := TreeNode{
			Val: 1,

			Right: &TreeNode{
				Val: 2,
			},
		}
		//act
		result := isSameTree(&tree1, &tree2)
		//assert
		assert.Equal(t, false, result)
	})

	/*
		Input: p = [1,2,1], q = [1,1,2]
		Output: false
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		tree1 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 1,
			},
		}
		tree2 := TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}
		//act
		result := isSameTree(&tree1, &tree2)
		//assert
		assert.Equal(t, false, result)
	})

}
