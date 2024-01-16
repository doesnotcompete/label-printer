package code

import "image"

type CodeGenerator interface {
	GetImage() (image.Image, error)
}
