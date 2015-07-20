package boxer

import "image"

type Boxer struct {
	x, y int
	img  *image.Gray
}

func New(img *image.Gray) *Boxer {
	return &Boxer{
		img: img,
	}
}

func Next() *image.Image {
	return nil
}
