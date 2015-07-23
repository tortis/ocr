package boxer

import (
	"image"
	"image/color"
	"sort"
)

type Blob struct {
	Bounds  image.Rectangle
	Img     image.Image
	Bridges []Bridge
}

type Blobs struct {
	blobs       []*Blob
	totalWidth  float64
	totalHeight float64
}

func (this *Blobs) addBlob(blob *Blob) {
	this.blobs = append(this.blobs, blob)
	this.totalWidth += float64(blob.Bounds.Dx())
	this.totalHeight += float64(blob.Bounds.Dy())
}

func (this *Blobs) addAll(blobs Blobs) {
	for _, b := range blobs.blobs {
		this.addBlob(b)
	}
}

func (this *Blobs) avgWidth() float64 {
	return this.totalWidth / float64(len(this.blobs))
}

func (this *Blobs) avgHeight() float64 {
	return this.totalHeight / float64(len(this.blobs))
}

func Blobify(img *image.Gray) Blobs {
	var blobs Blobs
	visited := make([]bool, img.Bounds().Max.X*img.Bounds().Max.Y)

	// Scan for first/next blob.
	x, y := 0, 0

	for {
		x, y = nextPixel(x, y, img, visited)
		if x < 0 || y < 0 {
			for _, b := range blobs.blobs {
				// Only compute bridges for wide blobs
				if float64(b.Bounds.Dx()) >= blobs.avgWidth()*1.75 {
					b.computeBridges(blobs.avgWidth())
				}
			}

			return blobs
		}

		// Extract blob.
		b := blobAt(x, y, img, visited)
		x = b.Bounds.Max.X + 1

		blobs.addBlob(b)
	}
}

////////////////////////////////////////////////////////////////////////////////
// Does a top to bottom, left to right scan of img starting at x,y. When a    //
// black pixel is encountered, the scan stops and the position is returned    //
////////////////////////////////////////////////////////////////////////////////
func nextPixel(x, y int, img *image.Gray, vis []bool) (int, int) {
	offset := img.Bounds().Max.X
	for cy := y; cy < img.Bounds().Max.Y; cy++ {
		start := 0
		if cy == y {
			start = x
		}
		for cx := start; cx < img.Bounds().Max.X; cx++ {
			if !vis[cx+cy*offset] && img.GrayAt(cx, cy) == black {
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

	r.Max.Y += 1
	r.Max.X += 1

	return &Blob{Bounds: r, Img: img.SubImage(r)}
}

func (this *Blob) computeBridges(avgW float64) {
	var bridges Bridges
	edgeWidth := int(avgW * 0.7)

	// Run vertical scans over the blob
	for x := this.Img.Bounds().Min.X + edgeWidth; x <= this.Img.Bounds().Max.X-edgeWidth; x++ {
		//thickest := 0
		totalThickness := 0
		for y := this.Img.Bounds().Min.Y; y <= this.Img.Bounds().Max.Y; y++ {
			if this.Img.At(x, y).(color.Gray).Y == 0x0 {
				totalThickness += 1
			}
		}
		bridges = append(bridges, Bridge{
			Position:       x,
			TotalThickness: totalThickness,
		})
	}

	sort.Sort(bridges)
	this.Bridges = bridges
}
