package kthlargestelement

func Partition(nums []int, left, right int) int {
	i := left
	pivot := nums[right]
	for j := left; j <= right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func QuickSelect(nums []int, left, right, k int) int {
	pivot := Partition(nums, left, right)
	if pivot == k {
		return nums[pivot]
	} else if pivot > k {
		return QuickSelect(nums, left, pivot-1, k)
	} else {
		return QuickSelect(nums, pivot+1, right, k)
	}
}

func FindKthLargest(nums []int, k int) int {
	k = len(nums) - k
	QuickSelect(nums, 0, len(nums)-1, k)
	return nums[k]
}
