package mm

import "testing"

func TestMinI(t *testing.T) {
	if r := Min(0, 1); r != 0 {
		t.Fatal("Min should return 0 but returned ", r)
	}
}

func TestMaxI(t *testing.T) {
	if r := Max(0, 1); r != 1 {
		t.Fatal("Min should return 1 but returned ", r)
	}
}

func TestMinF(t *testing.T) {
	if r := Min(0.1, 1.0); r != 0.1 {
		t.Fatal("Min should return 0.1 but returned ", r)
	}
}

func TestMaxF(t *testing.T) {
	if r := Max(0.1, 1.0); r != 1.0 {
		t.Fatal("Min should return 1.0 but returned ", r)
	}
}
