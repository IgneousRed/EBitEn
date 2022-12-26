package EduTen

import m "github.com/IgneousRed/gomisc"

type Rad = m.Rad
type Vec2 = m.Vector2

var V2 = m.Vec2

type Verts []Vec2
type Inds []uint16
type Trigs struct {
	Verts Verts
	Inds  Inds
}

// // Add `amount` to each Vert.
// func (v Verts) Translate(amount Vec2) Verts {
// 	result := make(Verts, len(v))
// 	for i, p := range v {
// 		result[i] = p.Add(amount)
// 	}
// 	return result
// }

// // Rotate every Vert around origin by `amount`.
// func (v Verts) Rotate(amount Rad) Verts {
// 	newX := amount.Vec2()
// 	newY := newX.Rot90()
// 	result := make(Verts, len(v))
// 	for i, p := range v {
// 		result[i] = newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
// 	}
// 	return result
// }

// // Multiply `amount` with each Vert.
// func (v Verts) Scale(amount Vec2) Verts {
// 	result := make(Verts, len(v))
// 	for i, p := range v {
// 		result[i] = p.Mul(amount)
// 	}
// 	return result
// }

// // Multiply `amount` with each Vert.
// func (v Verts) Scale1(amount float64) Verts {
// 	result := make(Verts, len(v))
// 	for i, p := range v {
// 		result[i] = p.Mul1(amount)
// 	}
// 	return result
// }

// Multiply `amount` with each Vert.
func (v Verts) Transform(pos Vec2, rot Rad, scl Vec2) Verts {
	// return m.MapF(v, func(p Vec2) Vec2 {
	// 	return p.Mul(scl).Rot(rot).Add(pos)
	// })

	result := make(Verts, len(v))
	for i, p := range v {
		// x, y := p[0]*scl[0], p[1]*scl[1]
		result[i] = p.Mul(scl).Rot(rot).Add(pos)
	}
	return result
}

// Multiply `amount` with each Vert.
func (v Verts) Transform1(pos Vec2, rot Rad, scl float64) Verts {
	// return m.MapF(v, func(p Vec2) Vec2 {
	// 	return p.Mul1(scl).Rot(rot).Add(pos)
	// })

	result := make(Verts, len(v))
	for i, p := range v {
		result[i] = p.Mul1(scl).Rot(rot).Add(pos)
	}
	return result
}

// Multiply `amount` with each Vert.
func (t Trigs) Transform(pos Vec2, rot Rad, scl Vec2) Trigs {
	// t.Verts = m.MapF(t.Verts, func(p Vec2) Vec2 {
	// 	return p.Mul(scl).Rot(rot).Add(pos)
	// })
	// return t

	result := Trigs{make(Verts, len(t.Verts)), t.Inds}
	for i, p := range t.Verts {
		result.Verts[i] = p.Mul(scl).Rot(rot).Add(pos)
	}
	return result
}

// Multiply `amount` with each Vert.
func (t Trigs) Transform1(pos Vec2, rot Rad, scl float64) Trigs {
	// t.Verts = m.MapF(t.Verts, func(p Vec2) Vec2 {
	// 	return p.Mul1(scl).Rot(rot).Add(pos)
	// })
	// return t

	result := Trigs{make(Verts, len(t.Verts)), t.Inds}
	for i, p := range t.Verts {
		result.Verts[i] = p.Mul1(scl).Rot(rot).Add(pos)
	}
	return result
}
