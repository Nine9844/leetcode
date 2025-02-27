package hot100

func TwoSum(nums []int, target int) []int {
	l := len(nums)
	var ans []int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ { // 从 i+1 开始，并且检查到数组的末尾
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				ans = append(ans, i)
				ans = append(ans, j)
				return ans // 找到结果后直接返回，提高效率
			}
		}
	}
	return ans // 如果没有找到结果，返回空数组
}
