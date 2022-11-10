package containerwithmostwater

func MaxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		var area int
		length := right - left
		if height[left] < height[right] {
			area = height[left] * length
			left += 1
		} else {
			area = height[right] * length
			right -= 1
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
