package misc

func TwoSum(nums []int, target int) []int {

	/*
		Input: nums = [2,7,11,15], target = 9
		Output: [0,1]
		Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	*/

	n := len(nums)
	sum := make([]int, 2)

	for i := 0; i < n; i++ {
		for j := n - 1; j > 0; j-- {
			if i < j {
				sumValue := nums[i] + nums[j]
				if sumValue == target {
					sum[i] = j
					return []int{i, sum[i]}
				}
			}
		}
	}

	return []int{-1}

}

func TwoSumOptimized(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}

	return []int{}

}
