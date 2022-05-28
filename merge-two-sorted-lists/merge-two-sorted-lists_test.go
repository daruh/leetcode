package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func traverse(node *ListNode) []int {
	var result []int
	current := node
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestMergeTwoSortedLists(t *testing.T) {

	/*Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		l1 := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
		}

		l2 := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		}
		//act
		mergedList := mergeTwoLists(&l1, &l2)
		//assert
		result := traverse(mergedList)
		expected := []int{1, 1, 2, 3, 4, 4}
		assert.Equal(t, expected, result)
	})

	/*Input: list1 = [], list2 = []
	Output: []
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		var (
			l1 ListNode
			l2 ListNode
		)
		//act
		mergedList := mergeTwoLists(&l1, &l2)
		//assert
		result := traverse(mergedList)
		var expected []int
		assert.Equal(t, expected, result)
	})

	/*Input: list1 = [], list2 = [0]
	Output: [0]
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		var l1 ListNode

		l2 := ListNode{
			Val: 0,
		}
		//act
		mergedList := mergeTwoLists(&l1, &l2)
		//assert
		result := traverse(mergedList)
		expected := []int{0}
		assert.Equal(t, expected, result)
	})
}
