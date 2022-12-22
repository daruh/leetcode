package remove_element

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	var ptr int
	tail := len(nums) - 1
	l := len(nums)
	for tail >= ptr && tail > 0 && ptr < l {

		if nums[ptr] != val && ptr < l {
			ptr++
		}
		if nums[tail] == val && tail > 0 {
			tail--
		}
		if ptr >= l {
			break
		}
		if tail < 0 {
			break
		}

		if nums[ptr] == val && nums[tail] != val {
			nums[ptr] = nums[tail]
			if tail <= ptr {
				break
			}
			ptr++
			tail--
		}

	}

	if ptr == 0 && nums[ptr] == val {
		return 0
	}
	if ptr+1 > tail+1 {
		return tail + 1
	} else {
		return ptr + 1
	}
}

//nums = [0,1,2,2,3,0,4,2], val = 2
//1 [0,1,2,2,3,0,4,2], val = 2
//   p p'x        t' t
