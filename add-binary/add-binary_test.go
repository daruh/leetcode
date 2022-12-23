package plus_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {

	/*
		Input: a = "11", b = "1"
		Output: "100"
	*/
	t.Run("Example 1", func(t *testing.T) {
		//arrange
		a := "11"
		b := "1"

		//act
		result := addBinary(a, b)
		//assert
		assert.Equal(t, "100", result)
	})

	/*
		Input: a = "1010", b = "1011"
		Output: "10101"
	*/
	t.Run("Example 2", func(t *testing.T) {
		//arrange
		a := "1010"
		b := "1011"
		//act
		result := addBinary(a, b)
		//assert
		assert.Equal(t, "10101", result)
	})

	/*
		Input: a = "1001", b = "1011"
		Output: "10100"
	*/
	t.Run("Example 3", func(t *testing.T) {
		//arrange
		a := "1001"
		b := "1011"
		//act
		result := addBinary(a, b)
		//assert
		assert.Equal(t, "10100", result)
	})
}
