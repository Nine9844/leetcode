package hot100

import "sort"

// 字母异位词分组
/*
问题：2025/03/01 该题未能做出   2025/03/03 参考答案后写出
1、sort的使用：sort.Ints(); sort.Float64s(); sort.Strings(); sort.Slice()
2、键值对相关的定义和使用
3、复合数据结构的含义及使用
4、数据之间的相互转换，像一个字符串，要转为字节型进行比较后再转为字符串 []byte()  string()
*/

func GroupAnagrams(strs []string) [][]string {
	//创建一个map映射来存储同位词
	kk := make(map[string][]string)

	for _, x := range strs {
		xx := []byte(x)
		sort.Slice(xx, func(i, j int) bool {
			return xx[i] < xx[j]
		})
		xxx := string(xx)
		kk[xxx] = append(kk[xxx], x)
	}

	var result [][]string
	for _, y := range kk {
		result = append(result, y)
	}
	return result
}
