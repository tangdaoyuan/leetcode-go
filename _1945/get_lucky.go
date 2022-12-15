package _1945

import (
	"strconv"
)

func getLucky(s string, k int) int {
	str := ""
	for _, ch := range s {
		str += strconv.Itoa(int(ch - 'a' + 1))
	}

	for i := k; i > 0; i-- {
		sum := 0
		for _, ch := range str {
			sum += int(ch - '0')
		}
		str = strconv.Itoa(sum)
		if len(str) <= 1 {
			break
		}
	}
	ans, _ := strconv.Atoi(str)
	return ans
}
