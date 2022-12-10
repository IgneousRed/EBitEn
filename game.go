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

func WindowSize() Vec2 { return screenSize }
func WindowHalf() Vec2 { return screenHalf }

var screenSize, screenHalf Vec2
var screenCB []func(size, half Vec2)

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
	cursor = Vec2{float64(x), screenSize[1] - float64(y)}

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
	size := Vec2{float64(outsideX), float64(outsideY)}
	if size != screenSize {
		half := size.Div1(2)
		for _, cb := range screenCB {
			cb(size, half)
		}
		screenSize, screenHalf = size, half
	}
	return outsideX, outsideY
}

func WindowSizeSet(size Vec2) {
	eb.SetWindowSize(int(size[0]), int(size[1]))
}

func WindowTitleSet(title string) {
	eb.SetWindowTitle(title)
}

func Run(game Game) {
	m.FatalErr(eb.RunGame(&gameInternal{game.Update, game.Draw}), "")
}
