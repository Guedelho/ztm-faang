package trapwater

func TrapWater(height []int) int {
	totalWater := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				totalWater += maxLeft - height[left]
			}
			left += 1
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				totalWater += maxRight - height[right]
			}
			right -= 1
		}
	}
	return totalWater
}
