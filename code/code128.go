package code

import (
	"image"
	"github.com/boombuler/barcode/code128"
)

type Code128Generator struct {
	Data string
}

func (g Code128Generator) GetImage() (image.Image, error) {
	img, err := code128.Encode(g.Data)
	if err != nil {
		return nil, err
	}

	return img, err
}
