package boxer

import "image"

type Blob struct {
	Bounds image.Rectangle
}

func Blobify(img *image.Gray) []Blob {
	// Scan for first/next blob.

	// Extract blob.

	// Scan left from center of blob to find the next blob
}

func nextPixel(x, y int, *image.Gray) (int, int) {

}

func blobAt(x, y int, img *image.Gray) *Blob {

}
