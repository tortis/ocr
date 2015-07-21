package boxer

import (
	"image"
	"log"
)

type Blob struct {
	Bounds image.Rectangle
}

func Blobify(img *image.Gray) []*Blob {
	var blobs []*Blob
	visited := make([]bool, img.Bounds().Max.X*img.Bounds().Max.Y)

	// Scan for first/next blob.
	x, y := 0, 18

	for {
		x, y = nextPixel(x, y, img)
		log.Printf("Found a black pixel at (%d, %d)\n", x, y)
		if x < 0 || y < 0 {
			return blobs
		}

		// Extract blob.
		b := blobAt(x, y, img, visited)
		x = b.Bounds.Max.X + 1
		y = (b.Bounds.Max.Y + b.Bounds.Max.Y) / 2
		blobs = append(blobs, b)
	}
}

////////////////////////////////////////////////////////////////////////////////
// Does a top to bottom, left to right scan of img starting at x,y. When a    //
// black pixel is encountered, the scan stops and the position is returned    //
////////////////////////////////////////////////////////////////////////////////
func nextPixel(x, y int, img *image.Gray) (int, int) {
	for cy := y; cy < img.Bounds().Max.Y; cy++ {
		for cx := x; cx < img.Bounds().Max.X; cx++ {
			if img.GrayAt(cx, cy) == black {
				return cx, cy
			}
		}
	}
	return -1, -1
}

func blobAt(x, y int, img *image.Gray, vis []bool) *Blob {
	if img.GrayAt(x, y) == white {
		return nil
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	r := image.Rect(x, y, x, y)
	q := newQueue()
	q.q(point{x, y})

	for q.length != 0 {
		p := q.dq()
		if vis[p.x+width*p.y] {
			continue
		}
		vis[p.x+width*p.y] = true
		if p.x > r.Max.X {
			r.Max.X = p.x
		}
		if p.x < r.Min.X {
			r.Min.X = p.x
		}
		if p.y > r.Max.Y {
			r.Max.Y = p.y
		}
		if p.y < r.Min.Y {
			r.Min.Y = p.y
		}

		// Up
		if p.y-1 >= 0 {
			if !vis[p.x+(p.y-1)*width] && img.GrayAt(p.x, p.y-1) == black {
				q.q(point{p.x, p.y - 1})
			}
		}
		// Right
		if p.x+1 < width {
			if !vis[p.x+1+p.y*width] && img.GrayAt(p.x+1, p.y) == black {
				q.q(point{p.x + 1, p.y})
			}
		}
		// Down
		if p.y+1 < height {
			if !vis[p.x+(p.y+1)*width] && img.GrayAt(p.x, p.y+1) == black {
				q.q(point{p.x, p.y + 1})
			}
		}
		// Left
		if p.x-1 >= 0 {
			if !vis[p.x-1+p.y*width] && img.GrayAt(p.x-1, p.y) == black {
				q.q(point{p.x - 1, p.y})
			}
		}
	}

	return &Blob{Bounds: r}
}
