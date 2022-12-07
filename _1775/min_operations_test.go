package _1775

import "testing"

func TestBasic(t *testing.T) {
	t.Helper()
	if ans := minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}); ans != 3 {
		t.Errorf("expected be 3, but %d got", ans)
	}
}
