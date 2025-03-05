package hot100

/*
盛水最多的容器
2025/03/03 最题思路错误   2025/03/05 参考答案默写
问题：
1、理解不了题目意思：容水量是有两个指针只想数字中的最小值*指针之间的距离 ---在横轴减小的情况下要找一个最大，那么纵轴就要从低到高，那么就是低的那个移动去寻找高的那个
2、双指针的使用：如双指针在什么情况下自增自减
*/

func MaxArea(hight []int) int {
	n := len(hight)
	maxArea := 0
	left, right := 0, n-1

	for left < right {
		hh := min(hight[left], hight[right])
		area := hh * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if hight[left] < hight[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
