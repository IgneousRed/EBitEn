package EduTen

type Cam struct {
	pos   Vec2
	rot   Rad
	scl   float64
	trans Transform
}

func (c Cam) Pos() Vec2    { return c.pos }
func (c Cam) Rot() Rad     { return c.rot }
func (c Cam) Scl() float64 { return c.scl }
func (c *Cam) TransformSet(pos Vec2, rot Rad, scl float64) {
	c.pos, c.rot, c.scl = pos, rot, scl
	c.trans.Rot, c.trans.Scl = -rot, 1/scl
	c.trans.Pos = pos.Sub(windowHalf).Rot(rot).Mul1(scl).Add(windowHalf)
}
func (c *Cam) TransformChange(pos Vec2, rot Rad, scl float64) {
	c.pos, c.rot, c.scl = c.pos.Add(pos), c.rot+rot, c.scl*scl
	c.trans.Rot, c.trans.Scl = -c.rot, 1/c.scl
	c.trans.Pos = c.pos.Sub(windowHalf).Rot(c.rot).Mul1(c.scl).Add(windowHalf)
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
