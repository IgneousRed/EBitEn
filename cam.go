package EduTen

type Cam struct {
	pos Vec2
	rot Rad
	scl float64
}

func CamNew() Cam {
	result := Cam{pos: screenHalf, scl: 1}
	screenCB = append(screenCB, func(size, half Vec2) {
		result.pos.Sub(screenHalf).Add(half)
	})
	return result
}
func (c Cam) Pos() Vec2    { return screenHalf.Sub(c.pos) }
func (c Cam) Rot() Rad     { return -c.rot }
func (c Cam) Scl() float64 { return 1 / c.scl }
func (c *Cam) TransformAbsolute(pos Vec2, rot Rad, scl float64) {
	c.pos = screenHalf.Sub(pos)
	c.pos[0] = screenHalf[0] - pos[0]
	c.pos[1] = screenHalf[1] - pos[1]
	c.rot, c.scl = -rot, 1/scl
}
func (c *Cam) TransformRelative(pos Vec2, rot Rad, scl float64) {
	sclRcp := 1 / scl
	c.pos[0] = (c.pos[0] - screenHalf[0]) * sclRcp
	c.pos[1] = (c.pos[1] - screenHalf[1]) * sclRcp
	c.pos = c.pos.Rot(-rot)
	c.pos[0] += screenHalf[0] - pos[0]
	c.pos[1] += screenHalf[1] - pos[1]
	c.rot, c.scl = c.rot-rot, c.scl*sclRcp
}
func (c Cam) Verts(v Verts) Verts {
	return v.Transform1(c.pos, c.rot, c.scl)
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
