package EBitEn

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

var keysOld []eb.Key
var keysDown map[eb.Key]struct{}
var keysPressed map[eb.Key]struct{}
var keysUp map[eb.Key]struct{}

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

type Game interface {
	Update()
	Draw()
}

var windowSizeIX, windowSizeIY int
var windowSizeFY float32
var screen *eb.Image

type gameInternal struct {
	update func()
	draw   func()
}

func (g *gameInternal) Update() error {
	// update keys
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

	// run user code
	g.update()
	return nil
}

func (g *gameInternal) Draw(scr *eb.Image) {
	screen = scr
	g.draw()
}

func (g *gameInternal) Layout(outsideX, outsideY int) (screenX, screenY int) {
	return windowSizeIX, windowSizeIY
}

func InitGame(name string, windowSize m.Vec[int], game Game) {
	windowSizeIX, windowSizeIY = windowSize[0], windowSize[1]
	windowSizeFY = float32(windowSizeIY)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(windowSize[0], windowSize[1])
	a := gameInternal{}
	a.update = game.Update
	a.draw = game.Draw
	m.FatalErr("", eb.RunGame(&a))
}
