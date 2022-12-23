package search_insert_position

func searchInsert(nums []int, target int) int {

	return treeSearch(&nums, 0, len(nums)-1, target)
}

func treeSearch(nums *[]int, left, right, target int) int {

	if left == right {
		if (*nums)[left] == target {
			return left
		}
		if target < (*nums)[left] {
			return max(0, left-1)
		}
		if target > (*nums)[left] {
			return left + 1
		}
	}
	if left+1 == right {
		if (*nums)[left] == target {
			return left
		}
		if (*nums)[right] == target {
			return right
		}
		if target < (*nums)[left] {
			return max(0, left-1)
		}
		if target > (*nums)[left] && target < (*nums)[right] {
			return left + 1
		}
		if target >= (*nums)[right] {
			return right + 1
		}
	}

	half := (right - left + 1) / 2
	val := (*nums)[left+half]
	if val == target {
		return left + half
	}
	if target < val {
		return treeSearch(nums, left, left+half-1, target)
	}
	if target > val {
		return treeSearch(nums, left+half, right, target)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
