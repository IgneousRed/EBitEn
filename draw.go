package EBitEn

import (
	"image/color"
	"math"

	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
)

var emptyImg = eb.NewImage(1, 1)

func init() {
	emptyImg.Fill(color.White)
}

func DrawTrianglesF(vertices []m.Vec[float32], indices []uint16, clr Color) {
	colorR := float32(clr.R) / 255
	colorG := float32(clr.G) / 255
	colorB := float32(clr.B) / 255
	colorA := float32(clr.A) / 255
	verts := make([]eb.Vertex, len(vertices))
	for i, v := range vertices {
		verts[i].DstX = v[0]
		verts[i].DstY = windowSizeFY - v[1]
		verts[i].ColorR = colorR
		verts[i].ColorG = colorG
		verts[i].ColorB = colorB
		verts[i].ColorA = colorA
	}
	Screen.DrawTriangles(verts, indices, emptyImg, &eb.DrawTrianglesOptions{})
}
func DrawTrianglesI(vertices []m.Vec[int], indices []uint16, clr Color) {
	verts := m.MapF(vertices, func(v m.Vec[int]) m.Vec[float32] { return v.Float32() })
	DrawTrianglesF(verts, indices, clr)
}
func DrawLineF(a, b m.Vec[float32], thickness float32, clr Color) {
	normal := b.Sub(a).Normalize().Rotate90().Mul1(thickness * .5)
	DrawTrianglesF([]m.Vec[float32]{
		a.Sub(normal),
		a.Add(normal),
		b.Sub(normal),
		b.Add(normal),
	}, []uint16{0, 1, 2, 1, 2, 3}, clr)
}
func DrawLineI(a, b m.Vec[int], thickness int, clr Color) {
	DrawLineF(a.Float32(), b.Float32(), float32(thickness), clr)
}
func DrawRectangleF(pos, size m.Vec[float32], clr Color) {
	DrawTrianglesF([]m.Vec[float32]{
		pos,
		pos.Add(m.Vec[float32]{size[0], 0}),
		pos.Add(m.Vec[float32]{0, size[1]}),
		pos.Add(m.Vec[float32]{size[0], size[1]}),
	}, []uint16{0, 1, 2, 1, 2, 3}, clr)
}
func DrawRectangleI(pos, size m.Vec[int], clr Color) {
	DrawRectangleF(pos.Float32(), size.Float32(), clr)
}
func DrawCircleF(pos m.Vec[float32], size float32, points int, clr Color) {
	verts := make([]m.Vec[float32], points)
	for i := range verts {
		ang := float32(i) / float32(points) * math.Pi * 2
		verts[i] = m.Vec[float32]{m.Cos(ang), m.Sin(ang)}.Mul1(size).Add(pos)
	}
	inds := make([]uint16, 0, (points-2)*3)
	for i := 2; i < points; i++ {
		inds = append(inds, 0, uint16(i-1), uint16(i))
	}
	DrawTrianglesF(verts, inds, clr)
}
func DrawCircleI(pos m.Vec[int], size int, points int, clr Color) {
	DrawCircleF(pos.Float32(), float32(size), points, clr)
}
