package main

import "fmt"
import "lk/hot100"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := hot100.TwoSum(nums, target)
	if result != nil {
		fmt.Println("Indices:", result)
	} else {
		fmt.Println("No solution found.")
	}
}
