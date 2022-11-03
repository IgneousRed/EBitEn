package EBitEn

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
