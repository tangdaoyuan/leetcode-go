package _1775

func minOperations(nums1, nums2 []int) (ans int) {
	if 6*len(nums1) < len(nums2) || 6*len(nums2) < len(nums1) {
		return -1
	}
	d := 0
	for _, x := range nums2 {
		d += x
	}
	for _, x := range nums1 {
		d -= x
	}
	if d < 0 {
		d = -d
		nums1, nums2 = nums2, nums1
	}
	cnt := [6]int{}
	for _, x := range nums1 {
		cnt[6-x]++
	}
	for _, x := range nums2 {
		cnt[x-1]++
	}
	for i := 5; ; i-- {
		if i*cnt[i] >= d {
			return ans + (d+i-1)/i
		}
		ans += cnt[i]
		d -= i * cnt[i]
	}
}
