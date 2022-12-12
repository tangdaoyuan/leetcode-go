package _1781

import "testing"

func TestBeautySum(t *testing.T) {
	t.Helper()
	if ans := beautySum("aabcb"); ans != 5 {
		t.Errorf("expected be 5, but %d got", ans)
	}
	if ans := beautySum("aabcbaa"); ans != 17 {
		t.Errorf("expected be 5, but %d got", ans)
	}
}
