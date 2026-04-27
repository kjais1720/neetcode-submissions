func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 || n > 0 {
		if n == 0 || (m > 0 && nums1[m-1] > nums2[n-1]) {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else if m == 0 || (n > 0 && nums2[n-1] > nums1[m-1]) {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			nums1[m+n-2] = nums2[n-1]
			n--
			m--
		}
	}
}
