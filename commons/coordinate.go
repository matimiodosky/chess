package commons

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Row() int {
	return c.y
}

func (c Coordinate) Column() int {
	return c.x
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}
