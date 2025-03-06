package hot100

/*
无重复字符的最长子串
2025/03/05 个人解法，但是过了一会忘了是怎么写的了 2025/03/06
问题：
滑动窗口
*/

func LengthOfLongestSubstring(s string) int {

	//官方解法
	//定义一个映射和字符串的长度
	m := map[byte]int{} //要循环遍历字符串，是对比字节
	n := len(s)

	//定义一个右指针
	rk := -1
	//定义结果
	ans := 0
	//左指针开始遍历
	for i := 0; i < n; i++ {
		// 此步骤为移动窗口，则需要去掉重复的那个字符
		if i != 0 {
			delete(m, s[i-1])
		}

		//从i开始计算不重复字符
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		ans = max(ans, rk-i+1)
	}
	return ans

	//个人解法：。。。。解完自己也忘了为啥这么求了
	//long := 0
	//kk := 0
	//s1 := []byte(s)
	//ss := make(map[byte]bool)
	//start := 0 // 当前子串的起始位置
	//
	//for _, x := range s1 {
	//	// 如果字符已经出现过，并且其位置在当前子串的起始位置之后
	//	if ss[x] {
	//		// 更新最长子串的长度
	//		if kk > long {
	//			long = kk
	//		}
	//		// 移动起始位置，直到重复字符被移除
	//		for s1[start] != x {
	//			delete(ss, s1[start])
	//			start++
	//			kk--
	//		}
	//		start++
	//	} else {
	//		// 如果字符未重复，增加当前子串的长度
	//		ss[x] = true
	//		kk++
	//	}
	//}
	//// 最后检查一次，确保最长子串的长度被更新
	//if kk > long {
	//	long = kk
	//}
	//return long
}
