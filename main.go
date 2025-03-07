package main

import "fmt"
import "lk/hot100"

func main() {

	//无重复字符最长字串
	lengthOfLongestSubstring := "abcabcbb"
	lengthOfLongestSubstringResult := hot100.LengthOfLongestSubstring(lengthOfLongestSubstring)
	fmt.Println("lengthOfLongestSubstring:", lengthOfLongestSubstringResult)

	//盛最多水的容器
	maxArea := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxAreaResult := hot100.MaxArea(maxArea)
	fmt.Println("maxAreaResult:", maxAreaResult)

	//三数之和
	threeSum := []int{-1, 0, 1, 2, -1, -4}
	resultThreeSum := hot100.ThreeSum(threeSum)
	fmt.Println("threeSum:", resultThreeSum)

	//移动0
	zeroInts := []int{0, 1, 0, 3, 12}
	resultZero := hot100.MoveZeros(zeroInts)
	fmt.Println("moveZeros:", resultZero)

	//字母异位分组
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result2 := hot100.GroupAnagrams(strs)
	fmt.Println("groupAnagrams:", result2)

	//最长连续序列
	nums1 := []int{100, 4, 200, 1, 3, 2}
	result1 := hot100.LongestConsecutive(nums1)
	fmt.Println("long:", result1)

	//两数之和等于目标数
	nums := []int{2, 7, 11, 15}
	target := 9
	result := hot100.TwoSum(nums, target)
	if result != nil {
		fmt.Println("Indices:", result)
	} else {
		fmt.Println("No solution found.")
	}
}
