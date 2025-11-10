package _2025

// 两数之和
func twoSum(target int, nums []int) []int {
	numMap := make(map[int]int)
	for index, value := range nums {
		otherValue := target - value
		if otherIndex, found := numMap[otherValue]; found {
			return []int{index, otherIndex}
		}
		numMap[value] = index
	}
	return []int{}
}
