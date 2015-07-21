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
	log.Printf("Found %d blobs\n", len(blobs))
	for _, b := range blobs {
		log.Printf("blob: %v\n", b)
		drawRect(g, b.Bounds)
	}
	log.Println("Saving boxes")
	savePNG("sblobtest.out.png", g)
}

func drawRect(img *image.Gray, r image.Rectangle) {
	c := color.Gray{0x99}
	log.Printf("Drawing top and bottom of: %v\n", r)
	// Draw top and bottom lines
	for x := r.Min.X - 1; x <= r.Max.X; x++ {
		img.SetGray(x, r.Min.Y-1, c)
		img.SetGray(x, r.Max.Y+1, c)
	}

	log.Println("Drawing sides")
	// Draw side lines
	for y := r.Min.Y - 1; y <= r.Max.Y; y++ {
		img.SetGray(r.Min.X-1, y, c)
		img.SetGray(r.Max.X+1, y, c)
	}
}
