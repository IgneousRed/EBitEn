package EduTen

import m "github.com/IgneousRed/gomisc"

type Rad float64
type Deg float64
type Vec2 [2]float64

// Convert to degrees.
func (r Rad) Deg() Deg {
	return Deg(r * m.Rad2Deg)
}

// Convert to radians.
func (r Deg) Rad() Rad {
	return Rad(r * m.Deg2Rad)
}

// Radian cosine.
func (r Rad) Cos() float64 {
	return m.Cos(float64(r))
}

// Degree cosine.
func (r Deg) Cos() float64 {
	return m.Cos(float64(r.Rad()))
}

// Radian sine.
func (r Rad) Sin() float64 {
	return m.Sin(float64(r))
}

// Degree sine.
func (r Deg) Sin() float64 {
	return m.Sin(float64(r.Rad()))
}

// Are `v` and `other` identical.
func (v Vec2) Eq(other Vec2) bool {
	return v[0] == other[0] && v[1] == other[1]
}

// Changes sign of each `v` element.
func (v Vec2) Neg() Vec2 {
	return Vec2{-v[0], -v[1]}
}

// Reciprocates each `v` element.
func (v Vec2) Rcp() Vec2 {
	return Vec2{1 / v[0], 1 / v[1]}
}

// `v` and `other` pairwise add.
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v[0] + other[0], v[1] + other[1]}
}

// Add `other` to each `v` element.
func (v Vec2) Add1(other float64) Vec2 {
	return Vec2{v[0] + other, v[1] + other}
}

// `v` and `other` pairwise subtract.
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v[0] - other[0], v[1] - other[1]}
}

// Subtract `other` from each `v` element.
func (v Vec2) Sub1(other float64) Vec2 {
	return Vec2{v[0] - other, v[1] - other}
}

// `v` and `other` pairwise multiply.
func (v Vec2) Mul(other Vec2) Vec2 {
	return Vec2{v[0] * other[0], v[1] * other[1]}
}

// Multiply `other` with each `v` element.
func (v Vec2) Mul1(other float64) Vec2 {
	return Vec2{v[0] * other, v[1] * other}
}

// `v` and `other` pairwise divide.
func (v Vec2) Div(other Vec2) Vec2 {
	return Vec2{v[0] / other[0], v[1] / other[1]}
}

// Divide `other` from each `v` element.
func (v Vec2) Div1(other float64) Vec2 {
	return Vec2{v[0] / other, v[1] / other}
}

// `v` and `other` pairwise wrap.
func (v Vec2) Wrap(lens Vec2) Vec2 {
	return Vec2{m.Wrap(v[0], lens[0]), m.Wrap(v[1], lens[1])}
}

// Wrap `len` to each `v` element.
func (v Vec2) Wrap1(len float64) Vec2 {
	return Vec2{m.Wrap(v[0], len), m.Wrap(v[1], len)}
}

// Make `v` elements absolute.
func (v Vec2) Abs() Vec2 {
	return Vec2{m.Abs(v[0]), m.Abs(v[1])}
}

// Lowest `v` element.
func (v Vec2) Min() float64 {
	if v[1] < v[0] {
		return v[1]
	}
	return v[0]
}

// Highest `v` element.
func (v Vec2) Max() float64 {
	if v[1] > v[0] {
		return v[1]
	}
	return v[0]
}

// `v` element Sum.
func (v Vec2) Sum() float64 {
	return v[0] + v[1]
}

// Angle to direction.
func (a Rad) Vec2() Vec2 {
	return Vec2{a.Cos(), a.Sin()}
}

// Angle to direction.
func (a Deg) Vec2() Vec2 {
	return Vec2{a.Cos(), a.Sin()}
}

// Direction to angle.
func (v Vec2) Rad() Rad {
	return Rad(m.Atan2(v[1], v[0]))
}

// Direction to angle.
func (v Vec2) Deg() Deg {
	return Rad(m.Atan2(v[1], v[0])).Deg()
}

// Floor `v` elements.
func (v Vec2) Floor() Vec2 {
	return Vec2{m.Floor(v[0]), m.Floor(v[1])}
}

// Round `v` elements.
func (v Vec2) Round() Vec2 {
	return Vec2{m.Round(v[0]), m.Round(v[1])}
}

// `v` and `other` linear interpolation.
func (v Vec2) Lerp(other Vec2, t float64) Vec2 {
	return other.Sub(v).Mul1(t).Add(v)
}

// Magnitude.
func (v Vec2) Mag() float64 {
	return m.Sqrt(v[0]*v[0] + v[1]*v[1])
}

// `v` direction with `value` magnitude.
func (v Vec2) MagSet(value float64) Vec2 {
	if mag := v.Mag(); mag != 0 {
		return v.Mul1(value / mag)
	}
	return v
}

// `v` direction with 1 magnitude.
func (v Vec2) Norm() Vec2 {
	return v.MagSet(1)
}

// `v` and `other` dot product.
func (v Vec2) Dot(other Vec2) float64 {
	return v.Mul(other).Sum()
}

// Angle from `v` to `other`.
func (v Vec2) AngTo(other Vec2) Rad {
	return other.Sub(v).Rad()
}

// Clamps `v` magnitude.
func (v Vec2) ClampMag(max float64) Vec2 {
	if v.Mag() > max {
		return v.MagSet(max)
	}
	return v
}

// Distance between `v` and `other`.
func (v Vec2) Dst(other Vec2) float64 {
	return other.Sub(v).Mag()
}

// Move `v` towards `other` by `dlt`.
// `dlt` must be >= 0.
func (v Vec2) MoveTowards(other Vec2, dlt float64) Vec2 {
	offset := other.Sub(v)
	dst := offset.Mag()
	if dst <= dlt {
		return other
	}
	return offset.Mul1(dlt / dst).Add(v)
}

// Project `other` onto `v`, changing magnitude of `v`.
func (v Vec2) Project(other Vec2) Vec2 {
	return v.MagSet(v.Dot(other))
}

// Rotate `v` 90 degrees.
func (v Vec2) Rot90() Vec2 {
	return Vec2{-v[1], v[0]}
}

// Rotate `v` with angle `amount`.
func (v Vec2) RotRad(amount Rad) Vec2 {
	newX := amount.Vec2()
	return newX.Rot90().Mul1(v[1]).Add(newX.Mul1(v[0]))
}

// Rotate `v` with angle `amount`.
func (v Vec2) RotDeg(amount Deg) Vec2 {
	return v.RotRad(amount.Rad())
}

// Reflect `v` on `norm`.
// `norm` magnitude determines bounce magnitude.
func (v Vec2) Reflect(norm Vec2) Vec2 {
	return norm.Rot90().Norm().Project(v).Sub(norm.Project(v))
}

// Add `amount` to each point.
func Translate(points []Vec2, amount Vec2) []Vec2 {
	result := make([]Vec2, len(points))
	for i, p := range points {
		result[i] = p.Add(amount)
	}
	return result
}

// Rotate every point around origin by `amount`.
func RotateRad(points []Vec2, amount Rad) []Vec2 {
	newX := amount.Vec2()
	newY := newX.Rot90()
	result := make([]Vec2, len(points))
	for i, p := range points {
		result[i] = newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
	}
	return result
}

// Rotate every point around origin by `amount`.
func RotateDeg(points []Vec2, amount Deg) []Vec2 {
	newX := amount.Vec2()
	newY := newX.Rot90()
	result := make([]Vec2, len(points))
	for i, p := range points {
		result[i] = newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
	}
	return result
}

// Multiply `amount` with each point.
func Scale(points []Vec2, amount Vec2) []Vec2 {
	result := make([]Vec2, len(points))
	for i, p := range points {
		result[i] = p.Mul(amount)
	}
	return result
}

// Multiply `amount` with each point.
func Scale1(points []Vec2, amount float64) []Vec2 {
	result := make([]Vec2, len(points))
	for i, p := range points {
		result[i] = p.Mul1(amount)
	}
	return result
}
