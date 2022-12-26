package EduTen

import (
	eb "github.com/hajimehoshi/ebiten/v2"
)

// A Key represents a keyboard key.
// These keys represent pysical keys of US keyboard.
// For example, KeyQ represents Q key on US keyboards and ' (quote) key on Dvorak keyboards.
type Key = eb.Key

// Keys.
const (
	KeyA Key = iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyAltL
	KeyAltR
	KeyDown
	KeyLeft
	KeyRight
	KeyUp
	KeyBackquote
	KeyBackslash
	KeyBackspace
	KeyBracketL
	KeyBracketR
	KeyCapsLock
	KeyComma
	KeyContextMenu
	KeyControlL
	KeyControlR
	KeyDelete
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyEnd
	KeyEnter
	KeyEqual
	KeyEscape
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyHome
	KeyInsert
	KeyMetaL
	KeyMetaR
	KeyMinus
	KeyNumLock
	KeyNP0
	KeyNP1
	KeyNP2
	KeyNP3
	KeyNP4
	KeyNP5
	KeyNP6
	KeyNP7
	KeyNP8
	KeyNP9
	KeyNPAdd
	KeyNPDecimal
	KeyNPDivide
	KeyNPEnter
	KeyNPEqual
	KeyNPMultiply
	KeyNPSubtract
	KeyPageDown
	KeyPageUp
	KeyPause
	KeyPeriod
	KeyPrintScreen
	KeyQuote
	KeyScrollLock
	KeySemicolon
	KeyShiftL
	KeyShiftR
	KeySlash
	KeySpace
	KeyTab
	KeyAlt
	KeyControl
	KeyShift
	KeyMeta
	KeyMax   = KeyMeta
	keyCount = 109
)

// A Button represents a mouse button.
type Button = eb.MouseButton

// MouseButtons
const (
	ButtonL Button = iota
	ButtonR
	ButtonM
	buttonCount = iota
)

var keysOld []Key
var keysDown, keysPressed, keysUp [keyCount]bool
var buttonsDown, buttonsPressed, buttonsUp [buttonCount]bool
var cursor Vec2
var wheel float64

// Returns true if any key was just pressed
func KeysDown(keys ...Key) bool {
	for _, k := range keys {
		if keysDown[k] {
			return true
		}
	}
	return false
}

// Returns true if any key is pressed
func KeysPressed(keys ...Key) bool {
	for _, k := range keys {
		if keysPressed[k] {
			return true
		}
	}
	return false
}

// Returns true if any key was just released
func KeysUp(keys ...Key) bool {
	for _, k := range keys {
		if keysUp[k] {
			return true
		}
	}
	return false
}

// Returns true if button was just pressed
func ButtonDown(button Button) bool {
	return buttonsDown[button]
}

// Returns true if button is pressed
func ButtonPressed(button Button) bool {
	return buttonsPressed[button]
}

// Returns true if button was just released
func ButtonUp(button Button) bool {
	return buttonsUp[button]
}

// Returns current cursor position
func Cursor() Vec2 {
	return cursor
}

// Returns wheel movement since last Update
func Wheel() float64 {
	return wheel
}
