package boxer

import (
	"image"
	"image/color"
	"log"
	"testing"
)

func TestBlober(t *testing.T) {
	img := loadPNG("test4.png")
	bin := Binarize(img, 32)
	blobs := Blobify(bin)

	log.Printf("Found %d blobs\n", len(blobs.blobs))

	c := color.Gray{0x99}
	lc := color.Gray{0xcc}
	for _, b := range blobs.blobs {
		if b.Bridges != nil {
			drawLine(bin, b.Bridges[0].Position, b.Bounds.Min.Y, b.Bounds.Dy(), lc)
		}
		drawRect(bin, b.Bounds, c)
	}
	log.Println("Saving boxes")
	savePNG("sblobtest.out.png", bin)
}

func drawLine(img *image.Gray, x, y, h int, c color.Gray) {
	for cy := y; cy <= y+h; cy++ {
		img.SetGray(x, cy, c)
	}
}

func drawRect(img *image.Gray, r image.Rectangle, c color.Gray) {
	// Draw top and bottom lines
	for x := r.Min.X - 1; x <= r.Max.X; x++ {
		img.SetGray(x, r.Min.Y-1, c)
		img.SetGray(x, r.Max.Y+1, c)
	}

	// Draw side lines
	for y := r.Min.Y - 1; y <= r.Max.Y; y++ {
		img.SetGray(r.Min.X-1, y, c)
		img.SetGray(r.Max.X+1, y, c)
	}
}
