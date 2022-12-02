package EduTen

import eb "github.com/hajimehoshi/ebiten/v2"

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
