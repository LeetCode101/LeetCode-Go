package algorithms

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	tail := len(nums1) - 1

	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[tail] = nums1[i]
			i -= 1
		} else {
			nums1[tail] = nums2[j]
			j -= 1
		}

		tail -= 1
	}
}
