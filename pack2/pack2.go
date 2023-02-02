package pack2

type PackA struct {
	X, Y int
}

func (point *PackA) CreatePoints() {
	point.X, point.Y = 0, 0
}

func (point *PackA) SumPoint() int {
	return point.X + point.Y
}

func (point *PackA) Square() int {
	sq := point.Y * point.X
	point.Y = 1
	return sq
}

func (point PackA) Square2() int {
	sq := point.Y * point.X
	point.Y = 3
	return sq
}
