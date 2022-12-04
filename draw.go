package EduTen

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

func DrawTriangles(vertices []m.Vec2F, indices []uint16, clr Color) {
	colorR := float32(clr.R) / 255
	colorG := float32(clr.G) / 255
	colorB := float32(clr.B) / 255
	colorA := float32(clr.A) / 255
	verts := make([]eb.Vertex, len(vertices))
	for i, v := range vertices {
		verts[i].DstX = float32(v[0])
		verts[i].DstY = windowSizeY32 - float32(v[1])
		verts[i].ColorR = colorR
		verts[i].ColorG = colorG
		verts[i].ColorB = colorB
		verts[i].ColorA = colorA
	}
	Screen.DrawTriangles(verts, indices, emptyImg, &eb.DrawTrianglesOptions{})
}
func DrawLine(a, b m.Vec2F, thickness float64, clr Color) {
	normal := b.Sub(a).Rot90().Norm().Mul1(thickness * .5)
	DrawTriangles([]m.Vec2F{
		a.Sub(normal),
		a.Add(normal),
		b.Sub(normal),
		b.Add(normal),
	}, []uint16{0, 1, 2, 1, 2, 3}, clr)
}
func DrawRectangle(pos, size m.Vec2F, clr Color) {
	DrawTriangles([]m.Vec2F{
		pos,
		pos.Add(m.Vec2F{size[0], 0}),
		pos.Add(m.Vec2F{0, size[1]}),
		pos.Add(m.Vec2F{size[0], size[1]}),
	}, []uint16{0, 1, 2, 1, 2, 3}, clr)
}
func DrawCircle(pos m.Vec2F, size float64, points int, clr Color) {
	verts := make([]m.Vec2F, points)
	for i := range verts {
		ang := float64(i) / float64(points) * math.Pi * 2
		verts[i] = m.Vec2F{m.Cos(ang), m.Sin(ang)}.Mul1(size).Add(pos)
	}
	inds := make([]uint16, 0, (points-2)*3)
	for i := 2; i < points; i++ {
		inds = append(inds, 0, uint16(i-1), uint16(i))
	}
	DrawTriangles(verts, inds, clr)
}
