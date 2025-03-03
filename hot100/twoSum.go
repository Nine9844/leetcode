package hot100

func TwoSum(nums []int, target int) []int {
	//官方中哈希表解法
	hashMap := map[int]int{}
	for i, x := range nums {
		//使用差值来进行比较
		if p, ok := hashMap[target-x]; ok {
			return []int{p, i}
		}
		hashMap[x] = i
	}
	return nil
}

//// 个人解题思路：暴力解法
//l := len(nums)
//var ans []int
//for i := 0; i < l-1; i++ {
//	for j := i + 1; j < l; j++ { // 从 i+1 开始，并且检查到数组的末尾
//		if i == j {
//			continue
//		}
//		if nums[i]+nums[j] == target {
//			ans = append(ans, i)
//			ans = append(ans, j)
//			return ans // 找到结果后直接返回，提高效率
//		}
//	}
//}
//return ans // 如果没有找到结果，返回空数组
//}
