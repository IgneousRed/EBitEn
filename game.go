package EduTen

import (
	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game interface {
	Update()
	Draw()
}

var windowSizeX, windowSizeY int
var windowSizeY32 float32
var windowSizeY64 float64
var Screen *eb.Image

type gameInternal struct {
	update func()
	draw   func()
}

func (g *gameInternal) Update() error {
	// update keys
	keysNew := m.MapF(inpututil.AppendPressedKeys(nil),
		func(k eb.Key) Key { return Key(k) },
	)
	keysDown = map[Key]struct{}{}
	for _, k := range keysNew {
		if _, ok := keysPressed[k]; !ok {
			keysDown[k] = struct{}{}
		}
	}
	keysPressed = map[Key]struct{}{}
	for _, k := range keysNew {
		keysPressed[k] = struct{}{}
	}
	keysUp = map[Key]struct{}{}
	for _, k := range keysOld {
		if _, ok := keysPressed[k]; !ok {
			keysUp[k] = struct{}{}
		}
	}
	keysOld = keysNew

	// updade cursor
	x, y := eb.CursorPosition()
	cursor = m.Vec2F{X: float64(x), Y: float64(windowSizeY - y)}

	// run user code
	g.update()
	return nil
}

func (g *gameInternal) Draw(scr *eb.Image) {
	Screen = scr
	g.draw()
}

func (g *gameInternal) Layout(outsideX, outsideY int) (screenX, screenY int) {
	return windowSizeX, windowSizeY
}

func InitGame(name string, windowSize m.Vec2F, game Game) {
	windowSizeX, windowSizeY = int(windowSize.X), int(windowSize.Y)
	windowSizeY32, windowSizeY64 = float32(windowSizeY), float64(windowSizeY)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(windowSizeX, windowSizeY)
	a := gameInternal{}
	a.update, a.draw = game.Update, game.Draw
	m.FatalErr("", eb.RunGame(&a))
}
