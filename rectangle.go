package geometry

type Rectangle struct {
	Width float64 `json:"width"`
	Height float64 `json:"height"`
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * r.Width + 2 *  r.Height
}
