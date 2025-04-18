package geometry

type Rectangle struct {
	Breadth float64
	Length  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float64 {
	return (r.Length + r.Breadth) * 2
}