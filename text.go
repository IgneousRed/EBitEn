package EBitEn

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	m "github.com/IgneousRed/gomisc"
	text "github.com/hajimehoshi/ebiten/v2/text"
)

var fontId = 0

type Font struct {
	id   int
	data []byte
}

func FontNew(path string) (Font, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Font{}, err
	}
	result := Font{fontId, data}
	fontId++
	return result, nil
}

var fonts = map[string]font.Face{}

func DrawTextI(f Font, size int, pos m.Vec[int], txt string, clr Color) {
	str := fmt.Sprint(f.id, size)
	face, ok := fonts[str]
	if !ok {
		ttf, _ := opentype.Parse(f.data)
		face, _ = opentype.NewFace(ttf, &opentype.FaceOptions{
			Size:    float64(size),
			DPI:     72,
			Hinting: font.HintingFull,
		})
		fonts[str] = face
	}
	text.Draw(Screen, txt, face, pos[0], windowSizeIY-pos[1], clr.Color())
}
func DrawTextF(font Font, size float64, pos m.Vec[float64], txt string, clr Color) {
	DrawTextI(font, int(size), pos.Int(), txt, clr)
}
