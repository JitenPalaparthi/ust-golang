package square

type Square float32

func (r *Square) What() string {
	return "Square"
}
func (s *Square) Area() float32 {
	return float32(*s * *s)
}

func (s *Square) Perimeter() float32 {
	return float32(4 * *s)
}
