package _1780

import "testing"

func TestCheckPowersOfThree(t *testing.T) {
	t.Helper()
	if ans := checkPowersOfThree(12); ans != true {
		t.Errorf("expected be true, but %t got", ans)
	}
}
