package algorithms

import "sort"

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	mergedNumbers := append(nums1, nums2...)
	middle := len(mergedNumbers) / 2

	sort.Ints(mergedNumbers)

	if len(mergedNumbers)%2 == 1 {
		return float64(mergedNumbers[middle])
	} else {
		return float64(mergedNumbers[middle-1]+mergedNumbers[middle]) / 2
	}
}
