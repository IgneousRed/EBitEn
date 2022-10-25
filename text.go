package ebitengine

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	m "github.com/IgneousRed/gomisc"
	eb "github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text"
)

type Font struct {
	face font.Face
	size float32
}

func NewFont(path string, size float32) (Font, error) {
	var result Font
	if fcm, err := os.ReadFile(path); err != nil {
		return result, err
	} else if ttf, err := opentype.Parse(fcm); err != nil {
		return result, err
	} else if face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	}); err != nil {
		return result, err
	} else {
		return Font{face, size}, nil
	}
}
func DrawTextI(scr *eb.Image, font Font, txt string, pos m.Vec[int], clr Color) {
	text.Draw(scr, txt, font.face, pos[0], WindowSizeY-pos[1], clr.Color())
}
func DrawTextF(scr *eb.Image, font Font, txt string, pos m.Vec[float32], clr Color) {
	text.Draw(scr, txt, font.face, int(pos[0]), WindowSizeY-int(pos[1]), clr.Color())
}
