package shape

type Shaper interface {
	Area() float32
	Perimeter() float32
}

type Typer interface {
	What()
}
