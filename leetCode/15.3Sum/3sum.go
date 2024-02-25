package leetcode

import (
	"sort"
)

/*
@Project ：LeetCode
@File    ：3sum.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/2/25 23:35
*/

//	https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0015.3Sum/

func threeSum(nums []int) [][]int {
	//  TODO 存在问题 未考虑重复的情况 测试用例未全部通过
	result := make([][]int, 0)
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++ //  去重 计数 后续用于判断第三个数字是否存在
	}
	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums) //  避免后续遍历的有序性，因此进行排序

	for i := 0; i < len(uniqNums); i++ {
		for j := i + 1; j < len(uniqNums); j++ {
			pairNum := 0 - uniqNums[i] - uniqNums[j]

			if pairNum > uniqNums[j] && counter[pairNum] > 0 {
				result = append(result, []int{uniqNums[i], uniqNums[j], pairNum})
			}
		}
	}

	return result
}
