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

type gameInternal struct {
	update func()
	draw   func(scr *Image)
}

// Mode in which a user resizes the window.
//
// Regardless of the resizing mode, an Ebiten application can still change the window size or make
// the window fullscreen by calling Ebiten functions.
type WindowResizing = eb.WindowResizingModeType

// WindowResizingModeTypes
const (
	// Disallow resizing the window by a user.
	WRDisabled WindowResizing = WindowResizing(eb.WindowResizingModeDisabled)

	// Disallow resizing the window,
	// but allow to make the window fullscreen by a user.
	// This works only on macOS so far.
	// On the other platforms, this is the same as WindowResizingModeDisabled.
	WROnlyFullscreen WindowResizing = WindowResizing(eb.WindowResizingModeOnlyFullscreenEnabled)

	// Allow resizing the window by a user.
	WREnabled WindowResizing = WindowResizing(eb.WindowResizingModeEnabled)
)

var windowSizeOld, windowHalfOld Vec2
var windowSize, windowHalf Vec2

func WindowSizeOld() Vec2 { return windowSizeOld }
func WindowHalfOld() Vec2 { return windowHalfOld }
func WindowSize() Vec2    { return windowSize }
func WindowHalf() Vec2    { return windowHalf }

func WindowSizeSet(size Vec2) {
	eb.SetWindowSize(int(size.X), int(size.Y))
}

func WindowMinimize() {
	eb.MinimizeWindow()
}

func WindowMaximize() {
	eb.MaximizeWindow()
}

func WindowResizingSet(mode WindowResizing) {
	eb.SetWindowResizingMode(mode)
}

func WindowTitleSet(title string) {
	eb.SetWindowTitle(title)
}

func (g *gameInternal) Layout(outsideX, outsideY int) (screenX, screenY int) {
	// window size
	windowSizeOld, windowHalfOld = windowSize, windowHalf
	windowSize = V2(float64(outsideX), float64(outsideY))
	windowHalf = windowSize.Div1(2)

	// camera
	camPos = camPos.Sub(windowHalfOld).Add(windowHalf)

	return outsideX, outsideY
}

func (g *gameInternal) Update() error {
	// keys
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

	// buttons
	for b := Button(0); b < buttonCount; b++ {
		new := eb.IsMouseButtonPressed(b)
		buttonsDown[b] = new && !buttonsPressed[b]
		buttonsUp[b] = !new && buttonsPressed[b]
		buttonsPressed[b] = new
	}

	// cursor
	x, y := eb.CursorPosition()
	cursor = V2(float64(x), windowSize.Y-float64(y))

	// wheel
	_, yf := eb.Wheel()
	wheel = yf

	// user code
	g.update()

	return nil
}

func (g *gameInternal) Draw(scr *Image) {
	g.draw(scr)
}

func Run(game Game) {
	m.FatalErr("", eb.RunGame(&gameInternal{game.Update, game.Draw}))
}
