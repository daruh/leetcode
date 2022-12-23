package add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAddTwoNumbers https://leetcode.com/problems/add-two-numbers/

func TestAddTwoNumbers(t *testing.T) {
	/*
		Input: l1 = [2,4,3], l2 = [5,6,4]
		Output: [7,0,8]
		Explanation: 342 + 465 = 807.
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		l1 := ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		}

		l2 := ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
				},
			},
		}
		//act
		result := addTwoNumbers(&l1, &l2)

		//assert
		assert.Equal(t, 7, result.Val)
		result = result.Next
		assert.Equal(t, 0, result.Val)
		result = result.Next
		assert.Equal(t, 8, result.Val)
	})
	/*
		Input: l1 = [0], l2 = [0]
		Output: [0]
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		l1 := ListNode{
			Val: 0,
		}
		l2 := ListNode{
			Val: 0,
		}
		//act
		result := addTwoNumbers(&l1, &l2)
		//assert
		assert.Equal(t, 0, result.Val)
	})

	/*
		Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
		Output: [8,9,9,9,0,0,0,1]
	*/

	t.Run("Example 3", func(t *testing.T) {
		//arrange
		l1 := ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
		}

		l2 := ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		}
		//act
		result := addTwoNumbers(&l1, &l2)

		//assert
		assert.Equal(t, 8, result.Val)
		result = result.Next
		assert.Equal(t, 9, result.Val)
		result = result.Next
		assert.Equal(t, 9, result.Val)
		result = result.Next
		assert.Equal(t, 9, result.Val)
		result = result.Next
		assert.Equal(t, 0, result.Val)
		result = result.Next
		assert.Equal(t, 0, result.Val)
		result = result.Next
		assert.Equal(t, 0, result.Val)
		result = result.Next
		assert.Equal(t, 1, result.Val)
	})
}
