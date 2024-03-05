package leetcode

/*
@Project ：LeetCode
@File    ：longestSubarray.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/3/5 23:21
*/

//	1493. 删掉一个元素以后全为 1 的最长子数组 https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/

func longestSubarray(nums []int) int {
	//  找一个包含一个0的最长子串 思路类似 https://leetcode.cn/problems/max-consecutive-ones-iii/
	ans := 0
	left := 0
	count0 := 0
	for right, num := range nums {
		count0 += 1 - num
		for count0 > 1 {
			count0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left)
	}
	return ans

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
