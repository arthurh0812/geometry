package geometry

import "math"

type Circle struct {
	Radius float64 `json:"radius"`
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
