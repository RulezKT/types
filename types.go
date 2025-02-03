package types

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
