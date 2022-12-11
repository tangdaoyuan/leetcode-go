package _1827

import "testing"

func TestMinOperations(t *testing.T) {
	t.Helper()
	if ans := minOperations([]int{1, 1, 1}); ans != 3 {
		t.Errorf("expected be 3, but %d got", ans)
	}
	if ans := minOperations([]int{1, 5, 2, 4, 1}); ans != 14 {
		t.Errorf("expected be 14, but %d got", ans)
	}
	if ans := minOperations([]int{8}); ans != 0 {
		t.Errorf("expected be 0, but %d got", ans)
	}
}
