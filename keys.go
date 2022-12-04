package EduTen

// A Key represents a keyboard key.
// These keys represent pysical keys of US keyboard.
// For example, KeyQ represents Q key on US keyboards and ' (quote) key on Dvorak keyboards.
type Key int

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
	KeyAltLeft
	KeyAltRight
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyArrowUp
	KeyBackquote
	KeyBackslash
	KeyBackspace
	KeyBracketLeft
	KeyBracketRight
	KeyCapsLock
	KeyComma
	KeyContextMenu
	KeyControlLeft
	KeyControlRight
	KeyDelete
	KeyDigit0
	KeyDigit1
	KeyDigit2
	KeyDigit3
	KeyDigit4
	KeyDigit5
	KeyDigit6
	KeyDigit7
	KeyDigit8
	KeyDigit9
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
	KeyMetaLeft
	KeyMetaRight
	KeyMinus
	KeyNumLock
	KeyNumpad0
	KeyNumpad1
	KeyNumpad2
	KeyNumpad3
	KeyNumpad4
	KeyNumpad5
	KeyNumpad6
	KeyNumpad7
	KeyNumpad8
	KeyNumpad9
	KeyNumpadAdd
	KeyNumpadDecimal
	KeyNumpadDivide
	KeyNumpadEnter
	KeyNumpadEqual
	KeyNumpadMultiply
	KeyNumpadSubtract
	KeyPageDown
	KeyPageUp
	KeyPause
	KeyPeriod
	KeyPrintScreen
	KeyQuote
	KeyScrollLock
	KeySemicolon
	KeyShiftLeft
	KeyShiftRight
	KeySlash
	KeySpace
	KeyTab
	KeyAlt
	KeyControl
	KeyShift
	KeyMeta
	KeyMax Key = KeyMeta
)

var keysOld []Key
var keysDown map[Key]struct{}
var keysPressed map[Key]struct{}
var keysUp map[Key]struct{}

// Returns true if any key was just pressed
func KeysDown(keys ...Key) bool {
	for _, k := range keys {
		if _, ok := keysDown[k]; ok {
			return true
		}
	}
	return false
}

// Returns true if any key is pressed
func KeysPressed(keys ...Key) bool {
	for _, k := range keys {
		if _, ok := keysPressed[k]; ok {
			return true
		}
	}
	return false
}

// Returns true if any key was just released
func KeysUp(keys ...Key) bool {
	for _, k := range keys {
		if _, ok := keysPressed[k]; ok {
			return true
		}
	}
	return false
}
