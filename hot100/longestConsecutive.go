package hot100

func LongestConsecutive(nums []int) int {

	//官方解法  使用哈希表
	numSet := map[int]bool{}
	//对数组去重简化统计
	for _, x := range nums {
		numSet[x] = true
	}
	//最长
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak

	////个人解法 思路先排序 因为是要求取数组中最长序列
	//nn := nums
	//sort.Ints(nn)
	//count := 0
	//max := 0
	//if len(nums) == 0 {
	//	return 0
	//}
	//for i, v := range nn {
	//	if i < len(nums)-1 {
	//		if nn[i+1] == v+1 {
	//			count++
	//		} else if nn[i+1] == v {
	//			continue
	//		} else {
	//			if count >= max {
	//				max = count + 1
	//			}
	//			count = 0
	//		}
	//	}
	//}
	//if count >= max {
	//	max = count + 1
	//}
	//return max
	//
}
