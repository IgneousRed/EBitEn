package EduTen

var camPos v2
var camRot rad
var camScl = 1.

func Cam() (pos v2, rot rad, scl f64) {
	return windowHalf.Sub(camPos), -camRot, 1 / camScl
}
func CamSet(pos v2, rot rad, scl f64) {
	camPos = windowHalf.Sub(pos)
	camRot, camScl = -rot, 1/scl
}
func CamTrans(pos v2, rot rad, scl f64) {
	camPos = camPos.Sub(windowHalf).Div1(scl).
		Rot(-rot).Sub(pos).Add(windowHalf)
	camRot, camScl = camRot-rot, camScl/scl
}
func CamVerts(v Verts) Verts {
	return v.Transform1(camPos, camRot, camScl)
}
func CamTrigs(t Trigs) Trigs {
	return Trigs{CamVerts(t.Verts), t.Inds}
}
func CamVec2(v v2) v2 {
	return v.Mul1(camScl).Rot(camRot).Add(camPos)
}
func CamDrawTriangles(scr *Image, trigs Trigs, clr Color) {
	DrawTriangles(scr, CamTrigs(trigs), clr)
}
func CamDrawLine(scr *Image, a, b v2, thickness f64, clr Color) {
	DrawLine(scr, CamVec2(a), CamVec2(b), thickness*camScl, clr)
}
func CamDrawRectangle(scr *Image, pos, size v2, clr Color) {
	DrawRectangle(scr, CamVec2(pos), size.Mul1(camScl), clr)
}
func CamDrawCircle(scr *Image, pos v2, size f64, points int, clr Color) {
	DrawCircle(scr, CamVec2(pos), size*camScl, points, clr)
}
