package boxer

type Bridge struct {
	Position       int
	TotalThickness int
	Thickest       int
}

type Bridges []Bridge

func (this Bridges) Len() int {
	return len(this)
}

func (this Bridges) Less(i, j int) bool {
	return this[i].TotalThickness < this[j].TotalThickness
}

func (this Bridges) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
