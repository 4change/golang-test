package left_bound_lt_eq

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}

	return left
}
