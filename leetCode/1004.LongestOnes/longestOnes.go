package _004_LongestOnes

/*
@Project ：LeetCode
@File    ：longestOnes.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/2/29 23:26
*/

//	https://leetcode.cn/problems/max-consecutive-ones-iii/

func longestOnes(nums []int, k int) int {
	ans := 0
	count0 := 0
	left := 0
	for right, num := range nums {
		count0 += 1 - num //  关键步骤
		for count0 > k {
			count0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
