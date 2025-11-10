package _2025

//func splitText(text string, maxLen int) []string {
//	var strs []string
//	for i := 0; i < len(text); i++ {
//		if text[i] != ' ' {
//			continue
//		}
//	}
//	return strs
//}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, num := range nums[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	// 不使用递归
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
