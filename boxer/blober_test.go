package boxer

import (
	"image"
	"image/color"
	"log"
	"testing"
)

func TestBlober(t *testing.T) {
	img := loadPNG("final.png")
	g := toGray(img)
	blobs := Blobify(g)

	// Compute the average width
	avgWidth := 0.0
	for _, b := range blobs {
		avgWidth += float64(b.Bounds.Dx()) / float64(len(blobs))
	}
	multiW := avgWidth * 1.75
	log.Printf("Found %d blobs\n", len(blobs))
	log.Printf("Average blob width: %f, multW: %f\n", avgWidth, multiW)

	c := color.Gray{0x99}
	bc := color.Gray{0x44}
	lc := color.Gray{0xcc}
	for _, b := range blobs {
		if float64(b.Bounds.Dx()) >= multiW {
			bridges := b.bridges(avgWidth)
			log.Printf("For big blob at %d, %d: 2 best brides: %d @ %d and %d @ %d\n",
				b.Bounds.Min.X,
				b.Bounds.Min.Y,
				bridges[0].TotalThickness,
				bridges[0].Position,
				bridges[1].TotalThickness,
				bridges[1].Position)

			drawLine(g, bridges[0].Position, b.Bounds.Min.Y, b.Bounds.Dy(), lc)
			drawRect(g, b.Bounds, bc)
		} else {
			drawRect(g, b.Bounds, c)
		}
	}
	log.Println("Saving boxes")
	savePNG("sblobtest.out.png", g)
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
