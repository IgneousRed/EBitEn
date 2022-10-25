package ebitengine

import (
	"image/color"

	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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

var WindowSizeX, WindowSizeY int
var WindowSizeFX, WindowSizeFY float32

func InitGame(name string, windowSizeX, windowSizeY int, game eb.Game) {
	WindowSizeX, WindowSizeY = windowSizeX, windowSizeY
	WindowSizeFX, WindowSizeFY = float32(windowSizeX), float32(windowSizeY)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(windowSizeX, windowSizeY)
	m.FatalErr("", eb.RunGame(game))
}
func Vec2I(x, y int) m.Vec[int] {
	return m.Vec[int]{x, y}
}
func Vec2F(x, y float32) m.Vec[float32] {
	return m.Vec[float32]{x, y}
}

type Keys map[eb.Key]struct{}

func GetKeys() Keys {
	result := Keys{}
	for _, k := range inpututil.AppendPressedKeys(nil) {
		result[k] = struct{}{}
	}
	return result
}
func (k Keys) Pressed(key eb.Key) bool {
	_, b := k[key]
	return b
}
