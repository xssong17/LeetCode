package leetCode

/*
@Project ：LeetCode
@File    ：twoSum.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/2/21 22:25
*/

/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:


	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1]
*/

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for index, num := range nums {
		pairNum := target - num
		pairNumIndex, ok := numsMap[pairNum]
		if !ok {
			numsMap[num] = index
			continue
		}
		return []int{pairNumIndex, index}
	}
	return []int{-1, -1}

}
