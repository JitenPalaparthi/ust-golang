package shape

type Shaper interface {
	Area() float32
	Perimeter() float32
	What() string
}

type IEmpty interface {
}
