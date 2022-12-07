package EduTen

type Verts []Vec2
type Inds []uint16
type Trigs struct {
	Verts Verts
	Inds  Inds
}
type Transform struct {
	Pos Vec2
	Rot Rad
	Scl float64
}

func TransformNew() Transform {
	return Transform{Scl: 1}
}

// Add `amount` to each Vert.
func (v Verts) Translate(amount Vec2) Verts {
	result := make(Verts, len(v))
	for i, p := range v {
		result[i] = p.Add(amount)
	}
	return result
}

// Rotate every Vert around origin by `amount`.
func (v Verts) Rotate(amount Rad) Verts {
	newX := amount.Vec2()
	newY := newX.Rot90()
	result := make(Verts, len(v))
	for i, p := range v {
		result[i] = newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
	}
	return result
}

// Multiply `amount` with each Vert.
func (v Verts) Scale(amount Vec2) Verts {
	result := make(Verts, len(v))
	for i, p := range v {
		result[i] = p.Mul(amount)
	}
	return result
}

// Multiply `amount` with each Vert.
func (v Verts) Scale1(amount float64) Verts {
	result := make(Verts, len(v))
	for i, p := range v {
		result[i] = p.Mul1(amount)
	}
	return result
}

// Multiply `amount` with each Vert.
func (v Verts) Transform(trans Transform) Verts {
	return v.Scale1(trans.Scl).Rotate(trans.Rot).Translate(trans.Pos)
}
