package EBitEn

import "image/color"

type Color color.RGBA

func (c Color) Color() color.Color {
	return color.RGBA(c)
}

var Black = Color{0, 0, 0, 255}
var White = Color{255, 255, 255, 255}
var Red = Color{255, 0, 0, 255}
var Green = Color{0, 255, 0, 255}
var Blue = Color{0, 0, 255, 255}
var Yellow = Color{255, 255, 0, 255}
var Cyan = Color{0, 255, 255, 255}
var Magenta = Color{255, 0, 255, 255}
