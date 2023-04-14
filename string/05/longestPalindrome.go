package main

func main() {

	s := palindrome("abacada", 3, 3)
	println(s)
}

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		// res = longest(res, s1, s2)
		if len(res) > len(s1) {
			res = res
		} else {
			res = s1
		}
		if len(res) > len(s2) {
			res = res
		} else {
			res = s2
		}
	}
	return res
}

//从中心向两端扩散的双指针技巧
func palindrome(s string, l int, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 双指针，向两边展开
		l--
		r++
	}
	// 返回以 s[l] 和 s[r] 为中心的最长回文串
	return s[l+1 : r]
}
