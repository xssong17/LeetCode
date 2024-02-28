package leetcode

/*
@Project ：LeetCode
@File    ：lengthOfLongestSubstring.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/2/28 23:17
*/

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	ans := 0
	window := [128]bool{}
	left := 0
	for right, c := range s {
		for window[c] {
			window[s[left]] = false
			left++
		}
		window[c] = true
		ans = max(ans, right-left+1)

	}
	return ans

}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
