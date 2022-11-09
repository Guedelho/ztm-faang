package twosum

func TwoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if val, ok := numsMap[num]; ok {
			return []int{val, i}
		}
		numsMap[target-num] = i
	}
	return []int{}
}
