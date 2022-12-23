package two_sum

import "math"

func twoSum(nums []int, target int) []int {

	setint := map[int]int{}
	output := [2]int{}

	for i := 0; i < len(nums); i++ {
		candidate := nums[i]
		complementaryIx := -1
		toBeComplementary := target - candidate
		found := false

		if val, ok := setint[toBeComplementary]; ok {
			complementaryIx = val
			output[0] = int(math.Min(float64(i), float64(complementaryIx)))
			output[1] = int(math.Max(float64(i), float64(complementaryIx)))
			found = true
		} else {
			setint[candidate] = i
		}
		if found {
			break
		}
	}

	return output[:]
}
