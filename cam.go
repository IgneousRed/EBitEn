package EduTen

var camPos v2
var camRot rad
var camScl = 1.

func Cam() (pos v2, rot rad, scl f64) {
	return windowHalf.Sub(camPos), -camRot, 1 / camScl
}
func CamSet(pos v2, rot rad, scl f64) {
	camPos = windowHalf.Sub(pos)
	camPos[0] = windowHalf[0] - pos[0]
	camPos[1] = windowHalf[1] - pos[1]
	camRot, camScl = -rot, 1/scl
}
func CamTrans(pos v2, rot rad, scl f64) {
	sclRcp := 1 / scl
	camPos[0] = (camPos[0] - windowHalf[0]) * sclRcp
	camPos[1] = (camPos[1] - windowHalf[1]) * sclRcp
	camPos = camPos.Rot(-rot)
	camPos[0] += windowHalf[0] - pos[0]
	camPos[1] += windowHalf[1] - pos[1]
	camRot, camScl = camRot-rot, camScl*sclRcp
}
func CamVerts(v Verts) Verts {
	return v.Transform1(camPos, camRot, camScl)
}
func CamVec2(v v2) v2 {
	return v.Mul1(camScl).Rot(camRot).Add(camPos)
}
func CamDrawTriangles(scr *Image, trigs Trigs, clr Color) {
	trigs.Verts = CamVerts(trigs.Verts)
	DrawTriangles(scr, trigs, clr)
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
