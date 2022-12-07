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

var windowSize, windowHalf Vec2

type gameInternal struct {
	update func()
	draw   func(scr *Image)
}

func WindowSize() Vec2 { return windowSize }
func WindowHalf() Vec2 { return windowHalf }

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
	cursor = Vec2{float64(x), windowSize[1] - float64(y)}

	// update wheel
	xf, yf := eb.Wheel()
	wheel = Vec2{xf, yf}

	// run user code
	g.update()
	return nil
}

func (g *gameInternal) Draw(scr *Image) {
	g.draw(scr)
}

func (g *gameInternal) Layout(outsideX, outsideY int) (screenX, screenY int) {
	return int(windowSize[0]), int(windowSize[1])
}

func InitGame(name string, size Vec2, game Game) {
	windowSize, windowHalf = size, size.Div1(2)
	eb.SetWindowTitle(name)
	eb.SetWindowSize(int(size[0]), int(size[1]))
	m.FatalErr(eb.RunGame(&gameInternal{game.Update, game.Draw}), "")
}
