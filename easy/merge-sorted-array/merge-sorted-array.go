package merge_sorted_array

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int) {

	ptr1 := 0
	l := m

	for i := 0; i < n; i++ {
		val2 := nums2[i]
		var j int
		for j = ptr1; nums1[j] <= val2; j++ {
			if j == l {
				break
			}
		}
		ptr1 = j

		insertAndShfit(j, l, val2, &nums1)
		l++
	}

}

func insertAndShfit(index int, l int, value int, nums *[]int) {

	for i := l; i > index; i-- {
		(*nums)[i] = (*nums)[i-1]
	}
	(*nums)[index] = value
}

//     3
// 1 2 x 4 5 6 _
