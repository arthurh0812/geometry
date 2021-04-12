package geometry

type Cuboid struct {
	height float64
	width float64
	depth float64
}

func NewCuboid(height, width, depth float64) Cuboid {
	return Cuboid{
		height: height,
		width: width,
		depth: depth,
	}
}

func (c *Cuboid) Height() float64 {
	return c.height
}

func (c *Cuboid) SetHeight(h float64) {
	c.height = h
}

func (c *Cuboid) Width() float64 {
	return c.width
}

func (c *Cuboid) SetWidth(w float64) {
	c.width = w
}

func (c *Cuboid) Depth() float64 {
	return c.depth
}

func (c *Cuboid) SetDepth(d float64) {
	c.depth = d
}

func (c *Cuboid) Volume() float64 {
	return c.height * c.width * c.depth
}
