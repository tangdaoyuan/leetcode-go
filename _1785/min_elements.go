package _1785

import "math"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	sum = goal - sum

	if sum < 0 {
		sum = -sum
	}
	ans := math.Ceil(float64(sum) / float64(limit))
	return int(ans)
}
