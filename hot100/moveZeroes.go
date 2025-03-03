package hot100

/*
移动0：只能对原数组进行操作
*/
func MoveZeros(nums []int) []int {

	//个人解法：超出时间限制
	for i, x := range nums {
		if x == 0 && i < len(nums) {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
	return nums
}
