package common

type Coords struct {
	X int32
	Y int32
}

func (c *Coords) Equals(o Coords) bool {
	return c.X == o.X && c.Y == o.Y
}
