package geometry

import "testing"

func TestTriangle(t *testing.T) {
	myTriangle := NewTriangle(10)
	myTriangle.SetA(3)
	myTriangle.SetGamma(40)

	h := myTriangle.CalcHeight()
	expected := 1.928
	if h != expected {
		t.Errorf("incorrect height for triangle 'myTriangle': expected=%v, got=%v", expected, h)
	}
}
