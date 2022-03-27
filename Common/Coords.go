package common

import . "github.com/fpiwowarczyk/ants-sim/Types"

func (c *Coords) Equals(o Coords) bool {
	return c.X == o.X && c.Y == o.Y
}
