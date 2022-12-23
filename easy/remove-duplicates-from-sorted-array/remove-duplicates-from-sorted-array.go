package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	var ptr, fwd int
	for fwd < len(nums)-1 {
		fwd++
		if nums[ptr] == nums[fwd] {
			continue
		} else {
			if fwd == ptr+1 {
				ptr++
				continue
			} else {
				ptr++
				nums[ptr] = nums[fwd]
			}
		}
	}
	return ptr + 1
}

//case1
//234
//pf
//.pf

//case 2
//2234
//p.f

//case 3
//22222222222223X
//p............f

//2322222222222333334567
// p...........f----f'
