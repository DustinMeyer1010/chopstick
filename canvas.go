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

// Get the value at the given position
func (c canvas) getValue(p Position) string {
	return string(c[p.Y][p.Y].value)
}

// Get the metadata at the given position
func (c canvas) getMetaData(p Position) any {
	return c[p.Y][p.X].MetaData
}

// Get the full element at the current position
func (c canvas) getElement(p Position) *Element {
	return &c[p.Y][p.X]
}

// Set the element at a given location
func (c *canvas) setElement(p Position, element Element) {
	(*c)[p.Y][p.X] = element
}
