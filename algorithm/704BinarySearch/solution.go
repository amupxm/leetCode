package binarysearch

func Search(nums []int, target int) int {
	l := 0
	h := len(nums) - 1

	for l <= h {
		mid := int((h + l) / 2)
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
			continue
		}
		if target < nums[mid] {
			h = mid - 1
			continue
		}
	}
	return -1
}
