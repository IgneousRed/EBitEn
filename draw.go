package ebitengine

import (
	"image/color"
	"math"

	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
)

func DrawTriangles(screen *eb.Image, vertices []m.Vec[float32], indices []uint16, col color.Color) {
	r, g, b, a := col.RGBA()
	colorR := float32(r) / 65535
	colorG := float32(g) / 65535
	colorB := float32(b) / 65535
	colorA := float32(a) / 65535
	verts := make([]eb.Vertex, len(vertices))
	for i, e := range vertices {
		verts[i].DstX = e[0]
		verts[i].DstY = e[1]
		verts[i].ColorR = colorR
		verts[i].ColorG = colorG
		verts[i].ColorB = colorB
		verts[i].ColorA = colorA
	}
	screen.DrawTriangles(verts, indices, emptyImg, &eb.DrawTrianglesOptions{})
}
func DrawLine(screen *eb.Image, a, b m.Vec[float32], thickness float32, col color.Color) {
	normal := b.Sub(a).Normalize().Rotate90().Mul1(thickness * .5)
	DrawTriangles(screen, []m.Vec[float32]{
		a.Sub(normal),
		a.Add(normal),
		b.Sub(normal),
		b.Add(normal),
	}, []uint16{0, 1, 2, 1, 2, 3}, col)
}
func DrawRectangle(screen *eb.Image, position, size m.Vec[float32], col color.Color) {
	DrawTriangles(screen, []m.Vec[float32]{
		position,
		position.Add(m.Vec[float32]{size[0], 0}),
		position.Add(m.Vec[float32]{0, size[1]}),
		position.Add(m.Vec[float32]{size[0], size[1]}),
	}, []uint16{0, 1, 2, 1, 2, 3}, col)
}
func DrawCircle(screen *eb.Image, position m.Vec[float32], size float32, points int, col color.Color) {
	verts := make([]m.Vec[float32], points)
	for i := range verts {
		ang := float32(i) / float32(points) * math.Pi * 2
		verts[i] = m.Vec[float32]{m.Cos(ang), m.Sin(ang)}.Mul1(size).Add(position)
	}
	inds := make([]uint16, 0, (points-2)*3)
	for i := 2; i < points; i++ {
		inds = append(inds, 0, uint16(i-1), uint16(i))
	}
	DrawTriangles(screen, verts, inds, col)
}
