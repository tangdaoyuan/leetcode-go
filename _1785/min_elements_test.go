package _1785

import "testing"

func TestMinElements(t *testing.T) {
	t.Helper()
	if ans := minElements([]int{1, -1, 1}, 3, -4); ans != 2 {
		t.Errorf("expected is 2, but %d got", ans)
	}
	if ans := minElements([]int{1, -10, 9, 1}, 100, 0); ans != 1 {
		t.Errorf("expected is 1, but %d got", ans)
	}
}
