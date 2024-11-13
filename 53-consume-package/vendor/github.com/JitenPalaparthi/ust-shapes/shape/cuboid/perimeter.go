package cuboid

func (c *Cuboid) Perimeter() float32 {
	return 4 * (c.L + c.B + c.H)
}
