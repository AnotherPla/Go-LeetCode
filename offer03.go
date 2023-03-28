package main

func FindRepeatNumber() int {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	// 原地置换
	/*
		for i := 0; i < len(nums); {
			if i == nums[i] {
				i++
				continue
			}
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
		return -1 */

	// hash **
	hash := make([]int, len(nums))
	for _, val := range nums {
		if hash[val] > 0 {
			return val
		} else {
			hash[val] = 1
		}
	}

	return -1
}
