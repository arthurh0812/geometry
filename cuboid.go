package geometry

type Cuboid struct {
	height float64
	width float64
	depth float64
}

func (c *Cuboid) Height() float64 {
	return c.height
}

func (c *Cuboid) Width() float64 {
	return c.width
}

func (c *Cuboid) Depth() float64 {
	return c.depth
}

func (c *Cuboid) Volume() float64 {
	return c.height * c.width * c.depth
}
