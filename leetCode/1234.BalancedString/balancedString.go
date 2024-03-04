package leetcode

/*
@Project ：LeetCode
@File    ：balancedString.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/3/4 23:36
*/

// https://leetcode.cn/problems/replace-the-substring-for-balanced-string/description/

func balancedString(s string) int {
	cnt := ['X']int{}
	m := len(s) / 4 //  预期出现次数
	for _, c := range s {
		cnt[c]++ //  计数
	}
	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m { //  符合条件
		return 0
	}

	ans := len(s)
	left := 0
	for right, c := range s {
		cnt[c]--
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-left+1)
			cnt[s[left]]++
			left++ // 缩小子串
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
