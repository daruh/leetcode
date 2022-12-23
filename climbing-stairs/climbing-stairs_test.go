package climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {

	/*
		Input: n = 2
		Output: 2
		Explanation: There are two ways to climb to the top.
		1. 1 step + 1 step
		2. 2 steps
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		n := 2
		//act
		result := climbStairs(n)
		//assert
		assert.Equal(t, 2, result)
	})

	/*
		Input: n = 3
		Output: 3
		Explanation: There are three ways to climb to the top.
		1. 1 step + 1 step + 1 step
		2. 1 step + 2 steps
		3. 2 steps + 1 step
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		n := 3
		//act
		result := climbStairs(n)
		//assert
		assert.Equal(t, 3, result)
	})

	//t.Run("Example 3", func(t *testing.T) {
	//	//arrange
	//	n := 4
	//	//act
	//	result := climbStairs(n)
	//	//assert
	//	assert.Equal(t, 3, result)
	//})
}
