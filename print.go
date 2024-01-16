package main

import (
	"fmt"
	"github.com/datumbrain/label-printer/tag"
	"image"
	"image/png"
	"os"
	"time"
)

func PrintTag(text, qrText string, codeType tag.CodeType) error {
	tg := tag.NewGenerator(96, 220, codeType)

	img, err := tg.GenerateImage(text, qrText)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("images/%d.png", time.Now().UnixMicro())

	err = saveImageToPng(filename, img)
	if err != nil {
		return err
	}

	err = runNiimprint("-c", "usb", "-a", "/dev/ttyACM0", "-d", "1", "-i", filename)
	if err != nil {
		return err
	}

	return nil // os.Remove(filename)
}

func saveImageToPng(filename string, img image.Image) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
