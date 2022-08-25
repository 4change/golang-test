package quick_sort

func SortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums) - 1)
}

func QuickSort(nums []int, left int, right int) []int {
	i, j := left, right

	pivot := nums[(left + right) / 2]

	temp := 0
	for i <= j {
		for nums[i] < pivot {
			i++
		}

		for nums[j] > pivot {
			j--
		}

		if i <= j {
			temp = nums[i]
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
