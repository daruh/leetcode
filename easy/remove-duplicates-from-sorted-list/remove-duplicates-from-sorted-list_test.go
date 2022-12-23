package remove_duplicates_from_sorted_list

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

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/
func TestRemoveDuplicates(t *testing.T) {

	/*Input: head = [1,1,2]
	Output: [1,2]*/
	t.Run("Example 1", func(t *testing.T) {

		//arrange
		list := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		}
		//act
		resultList := deleteDuplicates(&list)
		//assert
		result := traverse(resultList)
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})

	/*
		Input: head = [1,1,2,3,3]
		Output: [1,2,3]
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		list := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		}
		//act
		resultList := deleteDuplicates(&list)
		//assert
		result := traverse(resultList)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("Example 3", func(t *testing.T) {
		//arrange
		list := ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		}
		//act
		resultList := deleteDuplicates(&list)
		//assert
		result := traverse(resultList)
		expected := []int{1}
		assert.Equal(t, expected, result)
	})

	t.Run("Example 4", func(t *testing.T) {
		//arrange
		list := ListNode{
			Val: 1,
		}
		//act
		resultList := deleteDuplicates(&list)
		//assert
		result := traverse(resultList)
		expected := []int{1}
		assert.Equal(t, expected, result)
	})

	t.Run("Example 5", func(t *testing.T) {
		//arrange
		var list *ListNode = nil
		//act
		resultList := deleteDuplicates(list)
		//assert
		result := traverse(resultList)
		var expected []int
		assert.Equal(t, expected, result)
	})
}
