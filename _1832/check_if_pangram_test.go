package _1832

import "testing"

func TestCheckIfPangram(t *testing.T) {
	t.Helper()

	if ans := checkIfPangram("thequickbrownfoxjumpsoverthelazydog"); ans != true {
		t.Errorf("expected is true, but %t got", ans)
	}
	if ans := checkIfPangram("leetcode"); ans != false {
		t.Errorf("expected is false, but %t got", ans)
	}
}
