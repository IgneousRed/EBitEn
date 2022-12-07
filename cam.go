package EduTen

type Cam struct {
	pos Vec2
	rot Rad
	scl float64
}

func (c Cam) Pos() Vec2    { return c.pos.Neg() }
func (c Cam) Rot() Rad     { return -c.rot }
func (c Cam) Scl() float64 { return 1 / c.scl }
func (c *Cam) Transform(pos Vec2, rot Rad, scl float64) {
	c.rot, c.scl = -rot, 1/scl
	c.pos = pos.Sub(windowHalf).Rot(rot).Mul1(scl).Add(windowHalf)

	// c.scaleLinear += zoom * c.zoomMult
	// scaleNew := m.Pow(2, c.scaleLinear)
	// c.pos = c.pos.Sub(pos.Mul(windowSize).Mul1(c.moveMult)).
	// 	Sub(windowHalf).Mul1(scaleNew / c.scale).Add(windowHalf)
	// c.scale = scaleNew
}
func (c Cam) Verts(v Verts) Verts {
	return Translate(Rotate(Scale1(v, c.scl), c.rot), c.pos)
}
func (c Cam) Vec2(v Vec2) Vec2 {
	return v.Mul1(c.scl).Rot(c.rot).Add(c.pos)
}
func (c Cam) DrawTriangles(scr *Image, trigs Trigs, clr Color) {
	trigs.Verts = c.Verts(trigs.Verts)
	DrawTriangles(scr, trigs, clr)
}
func (c Cam) DrawLine(scr *Image, a, b Vec2, thickness float64, clr Color) {
	DrawLine(scr, c.Vec2(a), c.Vec2(b), thickness*c.scl, clr)
}
func (c Cam) DrawRectangle(scr *Image, pos, size Vec2, clr Color) {
	DrawRectangle(scr, c.Vec2(pos), size.Mul1(c.scl), clr)
}
func (c Cam) DrawCircle(scr *Image, pos Vec2, size float64, points int, clr Color) {
	DrawCircle(scr, c.Vec2(pos), size*c.scl, points, clr)
}
