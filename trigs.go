package EduTen

type Verts []Vec2
type Inds []uint16
type Trigs struct {
	Verts Verts
	Inds  Inds
}

// Add `amount` to each point.
func Translate(verts []Vec2, amount Vec2) []Vec2 {
	result := make([]Vec2, len(verts))
	for i, v := range verts {
		result[i] = v.Add(amount)
	}
	return result
}

// Rotate every point around origin by `amount`.
func Rotate(verts []Vec2, amount Rad) []Vec2 {
	newX := amount.Vec2()
	newY := newX.Rot90()
	result := make([]Vec2, len(verts))
	for i, v := range verts {
		result[i] = newX.Mul1(v[0]).Add(newY.Mul1(v[1]))
	}
	return result
}

// Multiply `amount` with each point.
func Scale(verts []Vec2, amount Vec2) []Vec2 {
	result := make([]Vec2, len(verts))
	for i, v := range verts {
		result[i] = v.Mul(amount)
	}
	return result
}

// Multiply `amount` with each point.
func Scale1(verts []Vec2, amount float64) []Vec2 {
	result := make([]Vec2, len(verts))
	for i, v := range verts {
		result[i] = v.Mul1(amount)
	}
	return result
}
