package EduTen

import (
	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game interface {
	Update()
	Draw(scr *Image)
}

var windowSizeX, windowSizeY int
var windowSizeY32 float32

type gameInternal struct {
	update func()
	draw   func(scr *Image)
}

func (g *gameInternal) Update() error {
	// update keys
	keysNew := inpututil.AppendPressedKeys(nil)
	keysDown = [keyCount]bool{}
	for _, k := range keysNew {
		keysDown[k] = !keysPressed[k]
	}
	keysPressed = [keyCount]bool{}
	for _, k := range keysNew {
		keysPressed[k] = true
	}
	keysUp = [keyCount]bool{}
	for _, k := range keysOld {
		keysUp[k] = !keysPressed[k]
	}
	keysOld = keysNew

	// update buttons
	for b := Button(0); b < buttonCount; b++ {
		new := eb.IsMouseButtonPressed(b)
		buttonsDown[b] = new && !buttonsPressed[b]
		buttonsUp[b] = !new && buttonsPressed[b]
		buttonsPressed[b] = new
	}

	// update cursor
	x, y := eb.CursorPosition()
	cursor = Vec2{float64(x), float64(windowSizeY - y)}

	// update wheel
	xf, yf := eb.Wheel()
	cursor = Vec2{xf, yf}

	// run user code
	g.update()
	return nil
}

func (g *gameInternal) Draw(scr *Image) {
	g.draw(scr)
}

func (g *gameInternal) Layout(outsideX, outsideY int) (screenX, screenY int) {
	return windowSizeX, windowSizeY
}

func InitGame(name string, windowSize Vec2, game Game) {
	windowSizeX, windowSizeY = int(windowSize[0]), int(windowSize[1])
	windowSizeY32 = float32(windowSizeY)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(windowSizeX, windowSizeY)
	m.FatalErr(eb.RunGame(&gameInternal{game.Update, game.Draw}), "")
}
