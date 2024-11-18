package cuboid

type Cuboid struct {
	L, B, H float32
}

func (r *Cuboid) What() string {
	return "Cuboid"
}

func (r *Cuboid) Area() float32 {
	return r.L * r.B * r.H
}
func (r *Cuboid) Perimeter() float32 {
	return 4 * (r.L + r.B + r.H)
}
