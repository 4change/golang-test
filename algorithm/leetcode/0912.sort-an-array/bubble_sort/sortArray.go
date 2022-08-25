package bubble_sort

func SortArray(nums []int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := 0; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j + 1] {
				temp := nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = temp
			}
		}
	}

	return nums
}
