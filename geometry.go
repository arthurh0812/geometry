package geometry

// A Shape is a 2D mathematical object
type Shape interface {
	Area() float64
}

// Figure is a 3D mathematical object
type Figure interface {
	Volume() float64
}