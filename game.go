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

var windowSizeIX, windowSizeIY int
var windowSizeY32 float32
var windowSizeY64 float64
var Screen *eb.Image

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
	Screen = scr
	g.draw()
}

func (g *gameInternal) Layout(outsideX, outsideY int) (screenX, screenY int) {
	return windowSizeIX, windowSizeIY
}

func InitGame(name string, windowSize m.Vec2F, game Game) {
	windowSizeIX, windowSizeIY = int(windowSize[0]), int(windowSize[1])
	windowSizeY32, windowSizeY64 = float32(windowSizeIY), float64(windowSizeIY)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(windowSizeIX, windowSizeIY)
	a := gameInternal{}
	a.update, a.draw = game.Update, game.Draw
	m.FatalErr("", eb.RunGame(&a))
}
