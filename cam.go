package EduTen

type Cam struct {
	trans Transform
}

func CamNew() Cam {
	return Cam{trans: Transform{Pos: windowHalf, Scl: 1}}
}
func (c Cam) Pos() Vec2    { return c.trans.Pos.Sub(windowHalf).Neg() }
func (c Cam) Rot() Rad     { return -c.trans.Rot }
func (c Cam) Scl() float64 { return 1 / c.trans.Scl }
func (c *Cam) TransformAbsolute(pos Vec2, rot Rad, scl float64) {
	c.trans.Pos = windowHalf.Sub(pos)
	c.trans.Rot, c.trans.Scl = -rot, 1/scl
}
func (c *Cam) TransformRelative(pos Vec2, rot Rad, scl float64) {
	sclRcp := 1 / scl
	for i := range c.trans.Pos {
		c.trans.Pos[i] -= windowHalf[i] + pos[i]
	}
	c.trans.Pos = c.trans.Pos.Rot(-rot)
	for i, p := range c.trans.Pos {
		c.trans.Pos[i] = p*sclRcp + windowHalf[i]
	}
	c.trans.Rot, c.trans.Scl = c.trans.Rot-rot, c.trans.Scl*sclRcp
}
func (c Cam) Verts(v Verts) Verts {
	return v.Transform(c.trans)
}
func (c Cam) Vec2(v Vec2) Vec2 {
	return v.Mul1(c.trans.Scl).Rot(c.trans.Rot).Add(c.trans.Pos)
}
func (c Cam) DrawTriangles(scr *Image, trigs Trigs, clr Color) {
	trigs.Verts = c.Verts(trigs.Verts)
	DrawTriangles(scr, trigs, clr)
}
func (c Cam) DrawLine(scr *Image, a, b Vec2, thickness float64, clr Color) {
	DrawLine(scr, c.Vec2(a), c.Vec2(b), thickness*c.trans.Scl, clr)
}
func (c Cam) DrawRectangle(scr *Image, pos, size Vec2, clr Color) {
	DrawRectangle(scr, c.Vec2(pos), size.Mul1(c.trans.Scl), clr)
}
func (c Cam) DrawCircle(scr *Image, pos Vec2, size float64, points int, clr Color) {
	DrawCircle(scr, c.Vec2(pos), size*c.trans.Scl, points, clr)
}
