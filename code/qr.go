package code

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
)

type QrGenerator struct {
	Data string
	Size int
}

func (g QrGenerator) GetImage() (image.Image, error) {
	img, err := qrcode.Encode(g.Data, qrcode.Medium, g.Size)
	if err != nil {
		return nil, err
	}

	return png.Decode(bytes.NewReader(img))
}
