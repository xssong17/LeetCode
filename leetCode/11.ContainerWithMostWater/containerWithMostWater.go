package leetCode

/*
@Project ：LeetCode
@File    ：containerWithMostWater.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2024/2/22 23:18
*/

//	https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0011.Container-With-Most-Water/

func maxArea(nums []int) (output int) {
	var (
		right = len(nums) - 1
		left  = 0
	)
	for left < right {

		x := right - left
		y := 0
		/*
			利用性质移动指针寻找最大值
			ai和aj构成的容器（记为A容器）大于所有min(ai, aj) 与中间的墙构成的容器（记为B类容器）
			因为所有B容器的宽度都小于A容器，而B容器的高度不大于B容器的一堵墙的高度，也就是min(ai, aj)。min(ai, aj)又是A容器的高度
		*/
		if nums[left] < nums[right] {
			y = nums[left]
			left++
		} else {
			y = nums[right]
			right--
		}
		s := x * y
		if s >= output {
			output = s
		}

	}
	return
}
