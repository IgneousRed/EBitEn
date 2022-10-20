package ebitengine

import (
	"image/color"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	eb "github.com/hajimehoshi/ebiten/v2"
	// text "github.com/hajimehoshi/ebiten/v2/text"
)

var emptyImg = eb.NewImage(1, 1)

func init() {
	emptyImg.Fill(color.White)
}
func Font(path string, size float64) (font.Face, error) {
	var result font.Face
	if fcm, err := os.ReadFile(path); err != nil {
		return result, err
	} else if ttf, err := opentype.Parse(fcm); err != nil {
		return result, err
	} else if result, err = opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}); err != nil {
		return result, err
	}
	return result, nil
}
