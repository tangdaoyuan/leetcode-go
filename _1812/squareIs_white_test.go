package _1812

import "testing"

func TestSquareIsWhite(t *testing.T) {
	t.Helper()
	if ans := squareIsWhite("a1"); ans != false {
		t.Errorf("expected be false, but %t got", ans)
	}
	if ans := squareIsWhite("h3"); ans != true {
		t.Errorf("expected be true, but %t got", ans)
	}
	if ans := squareIsWhite("c7"); ans != false {
		t.Errorf("expected be false, but %t got", ans)
	}
}
