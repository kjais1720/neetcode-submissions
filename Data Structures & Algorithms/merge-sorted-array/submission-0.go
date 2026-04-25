func merge(nums1 []int, m int, nums2 []int, n int) {
	for i:=0; i<n; i++ {
		j := m+i-1
		for ; j>=0; j-- {
			if nums1[j] > nums2[i] {
				nums1[j+1] = nums1[j]
				nums1[j] = nums2[i]
			} else {
				break
			}
		}
		if j == m+i-1 {
			nums1[m+i] = nums2[i]
		}
	}
}
