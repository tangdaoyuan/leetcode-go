package _1781

func beautySum(s string) int {
	ans := 0
	for i := range s {
		cnt := [26]int{}
		maxFreq := 0
		for _, c := range s[i:] {
			cnt[c-'a']++
			maxFreq = max(maxFreq, cnt[c-'a'])
			minFreq := len(s)
			for _, c := range cnt {
				if c > 0 {
					minFreq = min(minFreq, c)
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
