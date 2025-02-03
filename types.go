package types

const (

	EARTH   = 0
	MOON    = 1
	SUN     = 2
	MERCURY = 3
	VENUS   = 4
	MARS    = 5
	CERES   = 6
	JUPYTER = 7
	SATURN  = 8
	CHIRON  = 9
	URANUS  = 10
	NEPTUNE = 11
	PLUTO   = 12
	NNODE   = 13
	SNODE   = 14
	
)

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p *Position) Minus(p2 *Position) Position {

	return Position{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
		Z: p.Z - p2.Z,
	}
}

// Plus
func (p *Position) Plus(p2 *Position) Position {
	return Position{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}
