package chopstick

type Element struct {
	Value    rune
	MetaData any
}

type canvas [][]Element

func makeCanvas(height, width int) canvas {
	canvas := make([][]Element, height) // make the outer slice with 'height' rows
	for y := range canvas {
		row := make([]Element, width)
		for x := range row {
			row[x].Value = ' '
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
	(*c)[p.Y][p.X].Value = value
}

// Set the element at a given location
func (c *canvas) setElement(p Position, element Element) {
	(*c)[p.Y][p.X] = element
}

// Returns Value at given position
func (c canvas) getValue(p Position) string {
	return string(c[p.Y][p.Y].Value)
}

// Returns metadata at given position
func (c canvas) getMetaData(p Position) any {
	return c[p.Y][p.X].MetaData
}

// Returns element at given position
func (c canvas) getElement(p Position) *Element {
	return &c[p.Y][p.X]
}

func (c *canvas) ClearCanvas() {
	for y := range *c {
		for x := range (*c)[y] {
			(*c)[y][x] = Element{Value: ' '}
		}
	}
}

// Clear the canvas from position given to end of canvas
func (c *canvas) ClearToEndOfCanvas(p Position) {
	for y := range *c {
		if y >= p.Y {
			for x := range (*c)[y] {
				if x >= p.X {
					(*c)[y][x] = Element{Value: ' '}
				}
			}
		}
	}
}

// Clears the canvas from position give to start of canvas
func (c *canvas) ClearToStartOfCanvas(p Position) {
	for y := range *c {
		if y <= p.Y {
			for x := range (*c)[y] {
				if x <= p.X {
					(*c)[y][x] = Element{Value: ' '}
				}
			}
		}
	}
}

// Clears the canvas for position give to start of row of canvas
func (c *canvas) ClearToStartOfLineCanvas(p Position) {
	for x := range (*c)[p.Y] {
		if x <= p.X {
			(*c)[p.Y][x] = Element{Value: ' '}
		}
	}
}

// Clears the canvas for position give to end of row of canvas
func (c *canvas) ClearToEndOfLineCanvas(p Position) {
	for x := range (*c)[p.Y] {
		if x >= p.X {
			(*c)[p.Y][x] = Element{Value: ' '}
		}
	}
}

// Clears the canvas for entire row of the position it is at
func (c *canvas) ClearLineCanvas(p Position) {
	for x := range (*c)[p.Y] {
		(*c)[p.Y][x] = Element{Value: ' '}
	}
}
