package EduTen

var camPos Vec2
var camRot Rad
var camScl = 1.

func Cam() (pos Vec2, rot Rad, scl float64) {
	return windowHalf.Sub(camPos), -camRot, 1 / camScl
}
func CamSet(pos Vec2, rot Rad, scl float64) {
	camPos = windowHalf.Sub(pos)
	camPos[0] = windowHalf[0] - pos[0]
	camPos[1] = windowHalf[1] - pos[1]
	camRot, camScl = -rot, 1/scl
}
func CamTrans(pos Vec2, rot Rad, scl float64) {
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
func CamVec2(v Vec2) Vec2 {
	return v.Mul1(camScl).Rot(camRot).Add(camPos)
}
func CamDrawTriangles(scr *Image, trigs Trigs, clr Color) {
	trigs.Verts = CamVerts(trigs.Verts)
	DrawTriangles(scr, trigs, clr)
}
func CamDrawLine(scr *Image, a, b Vec2, thickness float64, clr Color) {
	DrawLine(scr, CamVec2(a), CamVec2(b), thickness*camScl, clr)
}
func CamDrawRectangle(scr *Image, pos, size Vec2, clr Color) {
	DrawRectangle(scr, CamVec2(pos), size.Mul1(camScl), clr)
}
func CamDrawCircle(scr *Image, pos Vec2, size float64, points int, clr Color) {
	DrawCircle(scr, CamVec2(pos), size*camScl, points, clr)
}
