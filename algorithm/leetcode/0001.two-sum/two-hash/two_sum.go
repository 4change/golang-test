package two_hash

func twoSum(nums []int, target int) []int {
	// map key: 数组值, value: 数组下标
	index := make(map[int]int, len(nums))

	// 数组元素入map
	for i := 0; i < len(nums); i++ {
		index[nums[i]] = i
	}

	// 查找map中是否包含对应的两个元素
	for i := 0; i < len(nums); i++ {
		if j, ok := index[target - nums[i]]; ok && j != i {
			return []int{i, j}
		}
	}

	return nil
}
