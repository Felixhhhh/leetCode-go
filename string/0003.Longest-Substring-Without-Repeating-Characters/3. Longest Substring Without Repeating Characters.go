package leetcode

//滑动窗口-哈希桶
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	res := 0 // 记录结果
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		window[c]++
		// 判断左侧窗口是否要收缩
		for window[c] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			window[d]--
		}
		// 在这里更新答案
		res = max(res, right-left)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
