package leetcode

/*
@Project ：LeetCode
@File    ：minSubArrayLen.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/2/27 22:35
*/

// 	209. 长度最小的子数组 https://leetcode.cn/problems/minimum-size-subarray-sum/description/

func minSubArrayLen(target int, nums []int) int {
	numsLen := len(nums)
	ans := numsLen + 1
	left := 0
	s := 0

	for right := 0; right < len(nums); right++ {
		s += nums[right]

		for s >= target {
			tmpAns := right - left + 1
			if tmpAns < ans {
				ans = tmpAns
			}
			s -= nums[left]
			left++

		}
	}
	if ans <= numsLen {
		return ans
	}

	return 0
}
