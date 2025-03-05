package hot100

/*
2:
移动0：只能对原数组进行操作
*/
func MoveZeros(nums []int) []int {

	//官方解法：使用双指针
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++

		}
		right++
	}
	return nums

	////个人解法：超出时间限制
	//for i, x := range nums {
	//	if x == 0 && i < len(nums) {
	//		nums = append(nums[:i], nums[i+1:]...)
	//		nums = append(nums, 0)
	//	}
	//}
	//return nums
}
