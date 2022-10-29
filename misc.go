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

var windowSizeIY int
var windowSizeFY float32

func InitGame(name string, windowSize m.Vec[int], game eb.Game) {
	windowSizeIY = windowSize[1]
	windowSizeFY = float32(windowSizeIY)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(windowSize[0], windowSize[1])
	m.FatalErr("", eb.RunGame(game))
}

var keysOld []eb.Key
var keysDown map[eb.Key]struct{}
var keysPressed map[eb.Key]struct{}
var keysUp map[eb.Key]struct{}

// Run once per Update
func KeysUpdate() {
	keysNew := inpututil.AppendPressedKeys(nil)
	keysDown = map[eb.Key]struct{}{}
	for _, k := range keysNew {
		if _, ok := keysPressed[k]; !ok {
			keysDown[k] = struct{}{}
		}
	}
	keysPressed = map[eb.Key]struct{}{}
	for _, k := range keysNew {
		keysPressed[k] = struct{}{}
	}
	keysUp = map[eb.Key]struct{}{}
	for _, k := range keysOld {
		if _, ok := keysPressed[k]; !ok {
			keysUp[k] = struct{}{}
		}
	}
	keysOld = keysNew
}

// Returns true if any key was just pressed
func KeysDown(keys ...eb.Key) bool {
	for _, k := range keys {
		if _, ok := keysDown[k]; ok {
			return true
		}
	}
	return false
}

// Returns true if any key is pressed
func KeysPressed(keys ...eb.Key) bool {
	for _, k := range keys {
		if _, ok := keysPressed[k]; ok {
			return true
		}
	}
	return false
}

// Returns true if any key was just released
func KeysUp(keys ...eb.Key) bool {
	for _, k := range keys {
		if _, ok := keysPressed[k]; ok {
			return true
		}
	}
	return false
}
