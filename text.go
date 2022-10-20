package ebitengine

import (
	"image/color"
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

var emptyImg = eb.NewImage(1, 1)

func init() {
	emptyImg.Fill(color.White)
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
func DrawText(scr *eb.Image, font Font, txt string, pos m.Vec[float32], clr color.Color) {
	text.Draw(scr, txt, font.face, int(pos[0]), int(pos[1]+font.size), clr)
}
