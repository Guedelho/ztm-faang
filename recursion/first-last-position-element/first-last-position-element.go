package firstlastpositionelement

func SearchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	left, right := FindLeft(nums, lo, hi, target, -1), FindRight(nums, lo, hi, target, -1)
	return []int{left, right}
}

func FindLeft(nums []int, lo, hi, target, i int) int {
	if lo > hi {
		return i
	}
	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		return FindLeft(nums, lo, mid-1, target, mid)
	} else if nums[mid] < target {
		return FindLeft(nums, mid+1, hi, target, i)
	} else {
		return FindLeft(nums, lo, mid-1, target, i)
	}
}

func FindRight(nums []int, lo, hi, target, i int) int {
	if lo > hi {
		return i
	}
	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		return FindRight(nums, mid+1, hi, target, mid)
	} else if nums[mid] < target {
		return FindRight(nums, mid+1, hi, target, i)
	} else {
		return FindRight(nums, lo, mid-1, target, i)
	}
}
