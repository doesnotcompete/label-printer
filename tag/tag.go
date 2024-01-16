package tag

import (
	"github.com/datumbrain/label-printer/code"
	"github.com/datumbrain/label-printer/text"
	"image"
	"bytes"
	"encoding/json"
)

type CodeType int
const (
	Qr CodeType = iota
	Code128
)

func (s CodeType) String() string {
	return toString[s]
}

var toString = map[CodeType]string {
	Qr: "qr",
	Code128: "code128",
}

var toID = map[string]CodeType {
	"qr": Qr,
	"code128": Code128,
}

func (s CodeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *CodeType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toID[j]
	return nil
}

type Generator struct {
	height int
	width  int
	code CodeType
}

func NewGenerator(height, width int, code CodeType) *Generator {
	return &Generator{height: height, width: width, code: code}
}

func (g Generator) GenerateImage(tag, codeString string) (image.Image, error) {
	var codegen code.CodeGenerator
	switch g.code {
	case Qr:
		codegen = &code.QrGenerator{Data: codeString, Size: g.width}
	case Code128:
		codegen = &code.Code128Generator{Data: codeString}
	}
	qrCode, err := codegen.GetImage()
	if err != nil {
		return nil, err
	}

	// getting text image
	txt, err := text.GetImage(text.Config{
		Height:       25,
		Width:        220,
		DPI:          240.0,
		Padding:      10,
		FontFile:     "fonts/Arial.ttf",
		FontSize:     6.0,
		Hinting:      text.Full,
		Spacing:      1.0,
		WhiteOnBlack: false,
	}, tag)
	if err != nil {
		return nil, err
	}

	//joining and rotating the image
	finalImage := joinImages(g.height, g.width, qrCode, txt)

	return rotateImageCounterClockwise(finalImage), nil
}
