package arrays

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums := append(nums1, nums2...)
	sort.Ints(nums)
	middle := len(nums) / 2
	var median float64 = 0
	if isRemainderZero := len(nums) % 2; isRemainderZero != 0 {
		median = float64(nums[middle])
	} else {
		median = (float64(nums[middle-1]) + float64(nums[middle])) / 2
	}
	return median
}
