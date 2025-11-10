/*
*
@author:baiyuechu
@desc:
@data:2025/11/7
*
*/
package _2025

import "sort"

// 算法题：寻找数组中第k大的数
func findKthLargest(nums []int, k int) int {
	// 按照选择排序的方式，找到第k大的数
	// 从大到小排序
	for i := 0; i < k; i++ {
		maxIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
	}
	return nums[k-1]
}

// 数组的全排列
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return res
}

// 全排列的优化版本
func permuteOptimize(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return res
}

// 三数之和
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// 三角形的边
func triangleNumber(nums []int) int {
	var res int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left] > nums[right] {
				res += right - left // 当nums[i]+nums[left] > nums[right]时，nums[i]作为三角形的最短边，nums[left]到nums[right-1]的数都可以作为三角形的第二条边
				left++              // 这里left++是为了找到下一个可能的三角形，因为nums[left]到nums[right-1]的数都可以作为三角形的第二条边
			} else {
				right--
			}
		}
	}
	return res
}
