package _1832

func checkIfPangram(sentence string) bool {
	cnt := [26]int{}
	count := 0
	for _, ch := range sentence {
		if cnt[ch-'a'] == 0 {
			count++
		}
		cnt[ch-'a']++
		if count >= 26 {
			break
		}
	}
	return count >= 26
}
