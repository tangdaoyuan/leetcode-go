package _1945

import "testing"

func TestGetLucky(t *testing.T) {
	t.Helper()
	if ans := getLucky("iiii", 1); ans != 36 {
		t.Errorf("expected is 36, but %d got", ans)
	}
	if ans := getLucky("leetcode", 2); ans != 6 {
		t.Errorf("expected is 6, but %d got", ans)
	}
}
