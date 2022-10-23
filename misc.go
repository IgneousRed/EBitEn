package ebitengine

import (
	"image/color"

	m "github.com/IgneousRed/gomisc"
)

var Black = color.RGBA{0, 0, 0, 255}
var White = color.RGBA{255, 255, 255, 255}
var Red = color.RGBA{255, 0, 0, 255}
var Green = color.RGBA{0, 255, 0, 255}
var Blue = color.RGBA{0, 0, 255, 255}
var Yellow = color.RGBA{255, 255, 0, 255}
var Cyan = color.RGBA{0, 255, 255, 255}
var Magenta = color.RGBA{255, 0, 255, 255}

func Vec2i(x, y int) m.Vec[int] {
	return m.Vec[int]{x, y}
}
func Vec2f(x, y float32) m.Vec[float32] {
	return m.Vec[float32]{x, y}
}
