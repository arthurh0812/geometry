package geometry

import "math"

type Triangle struct {
	A float64 		`json:"a"`
	B float64 		`json:"b"`
	C float64 		`json:"c"`
	Alpha float64 	`json:"alpha"` 	// must be in degrees
	Beta float64 	`json:"beta"` 	// must be in degrees
	Gamma float64 	`json:"gamma"` 	// must be in degrees
	Height float64 	`json:"height"`
}

func NewTriangle(base float64) *Triangle {
	return &Triangle{C: base}
}

func (t *Triangle) SetA(a float64) {
	t.A = a
}

func (t *Triangle) SetB(b float64) {
	t.B = b
}

func (t *Triangle) SetAlpha(alpha float64) {
	t.Alpha = alpha
}

func (t *Triangle) SetBeta(beta float64) {
	t.Beta = beta
}

func (t *Triangle) SetGamma(gamma float64) {
	t.Gamma = gamma
}

func (t *Triangle) Area() float64 {
	h := t.Height
	if h == 0 {
		h = t.CalcHeight() // try to calculate the height
	}
	return (t.C*h) / 2
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t *Triangle) SetHeight(h float64) {
	t.Height = h
}

func (t *Triangle) CalcHeight() float64 {
	var res float64
	if t.A != 0 && t.B != 0 && t.C != 0 { // use heron's formular
		res = t.heronsFormula()
	}
	if t.A != 0 && t.Gamma != 0 {
		res = t.A * math.Sin(t.Gamma)
	}
	return res
}

func (t *Triangle) IsEquilateral() bool {
	return t.Alpha == t.Beta
}

func (t *Triangle) heronsFormula() float64 {
	per := t.Perimeter()
	s := per/2
	area := math.Sqrt(s*(s-t.A)*(s-t.B)*(s-t.C)) // area = (c * height) / 2 => height = (2* area) /c
	return (2 * area) / t.C
}