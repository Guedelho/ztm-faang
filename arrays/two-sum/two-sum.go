package twosum

func contains(num int, nums map[int]int) bool {
	for k := range nums {
		if k == num {
			return true
		}
	}
	return false
}

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if contains(num, numsMap) {
			return []int{numsMap[num], i}
		}
		diff := target - num
		numsMap[diff] = i
	}
	return nil
}
