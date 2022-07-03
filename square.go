package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver Square) End() Point {
	return Point{x: receiver.start.x + receiver.start.x, y: receiver.start.y + receiver.start.y}
}

func (receiver Square) Area() uint {
	return receiver.a * receiver.a
}

func (receiver Square) Perimeter() uint {
	return receiver.a * 4
}
