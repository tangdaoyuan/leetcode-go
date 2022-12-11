package _1827

func minOperations(nums []int) int {
	var maxVal = 0
	var count = 0
	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		} else {
			count += maxVal - val + 1
			maxVal += 1
		}
	}
	return count
}
