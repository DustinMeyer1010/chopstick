package chopstick

type Element struct {
	value    rune
	MetaData any
}

type canvas [][]Element

func makeCanvas(height, width int) canvas {
	canvas := make([][]Element, height) // make the outer slice with 'height' rows
	for y := range canvas {
		row := make([]Element, width)
		for x := range row {
			row[x].value = ' '
		}
		canvas[y] = row // make each row with 'width' columns
	}
	if Debug != nil {
		Debug.Println(len(canvas[0]))
	}

	return canvas
}

// Set the metadata for the current position
func (c *canvas) setMetaData(p Position, metadata any) {
	(*c)[p.Y][p.X].MetaData = metadata
}

// Set the value for the current position
func (c *canvas) setValue(p Position, value rune) {
	(*c)[p.Y][p.X].value = value
}

// Set the element at a given location
func (c *canvas) setElement(p Position, element Element) {
	(*c)[p.Y][p.X] = element
}

// Returns Value at given position
func (c canvas) getValue(p Position) string {
	return string(c[p.Y][p.Y].value)
}

// Returns metadata at given position
func (c canvas) getMetaData(p Position) any {
	return c[p.Y][p.X].MetaData
}

// Returns element at given position
func (c canvas) getElement(p Position) *Element {
	return &c[p.Y][p.X]
}
