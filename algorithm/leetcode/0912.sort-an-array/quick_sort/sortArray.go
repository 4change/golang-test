package quick_sort

func sortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums) - 1)
}

func quickSort(nums []int, left int, right int) []int {
	i := left
	j := right

	pivot := nums[(left + right) / 2]

	for i <= j {
		for nums[i] < pivot {
			i++
		}

		for nums[j] > pivot {
			j--
		}

		if i <= j {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp

			i++
			j--
		}
	}

	if i < right {
		quickSort(nums, i, right)
	}

	if left < j {
		quickSort(nums, left, j)
	}

	return nums
}
