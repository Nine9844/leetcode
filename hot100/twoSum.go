package hot100

/*
问题总结：
1、for循环使用不熟练：for i,x :=range nums {} 用逗号隔开   for i:=0;i<len(nums);i++ {} 用分号隔开
2、个人思路只有暴力解法，使用双重for循环。在快速查找、去重、统计频率、缓存、映射关系、集合操作、唯一标识、解决冲突都可以考虑到使用哈希表，在golang中哈希表通过map展现
3、不能很好的判别复杂度
*/

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
