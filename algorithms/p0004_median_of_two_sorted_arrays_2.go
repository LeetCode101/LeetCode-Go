package algorithms

import "math"

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2

	return float64(findKth(nums1, 0, nums2, 0, left)+findKth(nums1, 0, nums2, 0, right)) / 2.0
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}

	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if k == 1 {
		return int(math.Min(float64(nums1[i]), float64(nums2[j])))
	}

	middleValue1 := math.MaxInt64
	middleValue2 := math.MaxInt64

	if i+k/2-1 < len(nums1) {
		middleValue1 = nums1[i+k/2-1]
	}

	if j+k/2-1 < len(nums2) {
		middleValue2 = nums2[j+k/2-1]
	}

	if middleValue1 < middleValue2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j+k/2, k-k/2)
	}
}
