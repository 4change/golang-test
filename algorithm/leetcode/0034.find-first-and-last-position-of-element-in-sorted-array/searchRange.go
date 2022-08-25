package _034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{leftBoundBinarySearch(nums, target), rightBoundBinarySearch(nums, target)}
}

func leftBoundBinarySearch(nums []int, target int) int {
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

func rightBoundBinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}