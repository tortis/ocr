package boxer

type Line struct {
	blobs    Blobs
	baseLine int
	topLine  int
}

func Lineify(blobs Blobs) []Line {
	// Create buckets for each y position
	buckets := make(map[int]Blobs)

	// Place each blob into a bucket
	for _, blob := range blobs.blobs {
		baseline := blob.Bounds.Max.Y
		buckets[baseline].addBlob(blob)
	}

	// Attempt to merge nearby buckets
	bucketMerge(buckets, blobs.avgHeight())
	bucketMerge(buckets, blobs.avgHeight())
	bucketMerge(buckets, blobs.avgHeight())

	// Creat a line from each bucket
	result := make([]Line, 0, len(buckets))
	for _, bs := range buckets {
		result = append(result, Line{
			blobs: bs,
		})
	}
	return result
}

func bucketMerge(buckets map[int]Blobs, avgHeight float64) {
	mergeRadius := int(avgHeight / 4)
	for y, bs := range buckets {
		thisLen := len(bs.blobs)
		for i := 1; i <= mergeRadius; i++ {
			// Check i up
			if bsIUp, ok := buckets[y+i]; ok {
				if len(bsIUp.blobs) > thisLen {
					log.Printf("Merging blobs at %d to %d\n", y, y+i)
					bsIUp.addAll(bs)
					delete(buckets, y)
				}
			}

			// Check i down
			if bsIDown, ok := buckets[y-i]; ok {
				if len(bsIDown.blobs) > this.Len {
					log.Printf("Merging blobs at %d to %d\n", y, y-i)
					bsIDown.addAll(bs)
					delete(buckets, y)
				}
			}
		}
	}
}
