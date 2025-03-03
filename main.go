package main

import "fmt"
import "lk/hot100"

func main() {

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
