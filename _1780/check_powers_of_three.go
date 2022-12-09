package _1780

func checkPowersOfThree(n int) bool {
	current := n
	for {
		if current == 1 {
			return true
		}
		if current%3 == 0 {
			current = current / 3
			continue
		}
		if (current-1)%3 == 0 {
			current = (current - 1) / 3
			continue
		}
		break
	}
	return false
}
