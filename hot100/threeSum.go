package hot100

import "sort"

/*
2025/03/03 暴力超出时间限制  2025/03/04 尝试复现官方解法逻辑失败  2025/03/05抄答案理解逻辑
问题：
1、对数据可以根据题目内容做一些前置处理，简化运算
2、对于取数的逻辑
3、双指针的使用：是前后双指针还是左右双指针，指针在什么时候进行变动
4、复合类型的字符串切片怎么拼接
*/

func ThreeSum(nums []int) [][]int {

	//官方解法，在查询第二个和第三个时使用双指针,且对数据做一些前置处理。如和前面的数据一致就不再查询
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		//去掉重复计算
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//序号
		third := n - 1
		//值
		target := -nums[first]
		//对第二个进行枚举
		for second := first + 1; second < n; second++ {
			//去掉重复计算
			if nums[second] > first+1 && nums[second] != nums[second-1] {
				continue
			}
			//然后写入右指针的移动条件。要确认第二个小于第三个，不要重复统计
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			//写退出循环的条件
			if second == third {
				break
			}
			//最后记录符合条件的数据
			if nums[first]+nums[second]+nums[third] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}

	}
	return ans

	////个人解法，超出时间限制
	//// 先对数组排序
	//sort.Ints(nums)
	//n := len(nums)
	//result := [][]int{}
	//
	//// 使用 map 存储两数之和及其对应的索引
	//all := make(map[int][][]int)
	//
	//// 遍历数组，存储两数之和及其索引
	//for i := 0; i < n; i++ {
	//	for j := i + 1; j < n; j++ {
	//		sum := nums[i] + nums[j]
	//		all[sum] = append(all[sum], []int{i, j})
	//	}
	//}
	//
	//// 遍历数组，查找满足条件的三元组
	//for k := 0; k < n; k++ {
	//	target := -nums[k]
	//	if pairs, ok := all[target]; ok {
	//		for _, pair := range pairs {
	//			i, j := pair[0], pair[1]
	//			// 确保三个索引不重复
	//			if i != k && j != k {
	//				// 将三元组排序，避免重复
	//				triplet := []int{nums[i], nums[j], nums[k]}
	//				sort.Ints(triplet)
	//				// 检查是否已经存在
	//				if !contains(result, triplet) {
	//					result = append(result, triplet)
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//return result
}

// 辅助函数：检查切片中是否包含某个三元组
func contains(result [][]int, triplet []int) bool {
	for _, r := range result {
		if r[0] == triplet[0] && r[1] == triplet[1] && r[2] == triplet[2] {
			return true
		}
	}
	return false
}
